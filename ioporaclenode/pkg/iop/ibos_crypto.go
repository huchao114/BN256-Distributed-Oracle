package iop

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/random"
	
	// 确保引用项目内部兼容 EVM 的 bn256
	"ioporaclenode/internal/pkg/kyber/pairing/bn256"
)

// ==========================================
// 核心数学常量 (与 Solidity BN256G1 一致)
// ==========================================
var (
	// BN254/alt_bn128 base field prime
	ibosPP, _      = new(big.Int).SetString("30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47", 16)
	ibosSqrtExp, _ = new(big.Int).SetString("0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52", 16)
	ibosB          = big.NewInt(3)
)

// IBOSSystem 封装 Kyber 的配对环境
type IBOSSystem struct {
	Suite pairing.Suite
	Pub   kyber.Point  // 主公钥 (G2)
	
	// [修复] 将字段重命名为 G1Gen，避免与下方的 G1() 方法冲突
	G1Gen kyber.Point  // G1 生成元
}

// IBOSSignature 对应论文中的 σ_i = (S_i, R_i)
type IBOSSignature struct {
	S kyber.Point // G1
	R kyber.Point // G2
}

// 1. Setup: 初始化系统
// 这里的 masterSecret 通常是 DKG 生成的，或者测试时硬编码为 1
func NewIBOSSystem(masterSecret kyber.Scalar) *IBOSSystem {
	suite := bn256.NewSuiteG2()
	
	// 主公钥 Pub = x * G2 (Kyber 默认 G2 为基点)
	pub := suite.G2().Point().Mul(masterSecret, nil)

	return &IBOSSystem{
		Suite: suite,
		Pub:   pub,
		// [修复] 对应上面的重命名
		G1Gen: suite.G1().Point().Base(),
	}
}

// 2. Extract: 提取私钥 d_ID = x * H(ID)
func (sys *IBOSSystem) Extract(masterSecret kyber.Scalar, id []byte) (kyber.Point, error) {
	// H(ID) -> G1 (必须使用与 Solidity 兼容的 HashToPoint)
	qID, err := sys.SolidityHashToG1(id)
	if err != nil {
		return nil, err
	}
	
	// d_ID = x * Q_ID
	// 使用 sys.G1() 方法获取群接口来创建新点
	dID := sys.G1().Point().Mul(masterSecret, qID)
	return dID, nil
}

// 3. OrderSign: 有序签名算法
// prevSig: 上一个签名的结果。如果是第一个，传 nil。
// message: 原始文档 m
// mySk: 当前签名者私钥 d_ID (G1 Point)
func (sys *IBOSSystem) OrderSign(prevSig *IBOSSignature, message []byte, mySk kyber.Point) (*IBOSSignature, error) {
	// 构造当前消息 M_i
	var currentMsg []byte
	if prevSig == nil {
		// Signer 1: M_1 = m
		currentMsg = message
	} else {
		// Signer i: M_i = m || Keccak256(S_{i-1})
		prevSBytes, err := prevSig.S.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal prev S: %v", err)
		}
		h := crypto.Keccak256(prevSBytes)
		currentMsg = append(message, h...)
	}

	// 生成随机数 r
	r := sys.Suite.G1().Scalar().Pick(random.New())

	// R_i = r * G2 (注意 R 在 G2 群)
	R := sys.Suite.G2().Point().Mul(r, nil)

	// 计算 H(M_i) -> G1
	hm, err := sys.SolidityHashToG1(currentMsg)
	if err != nil {
		return nil, fmt.Errorf("hash to point failed: %v", err)
	}

	// S_i = -(d_ID + r * H(M_i))
	rHM := sys.Suite.G1().Point().Mul(r, hm)
	sum := sys.Suite.G1().Point().Add(mySk, rHM)
	S := sys.Suite.G1().Point().Neg(sum)

	return &IBOSSignature{S: S, R: R}, nil
}

// ==========================================
// 核心工具函数：Solidity 兼容的 HashToPoint
// ==========================================

func (sys *IBOSSystem) SolidityHashToG1(data []byte) (kyber.Point, error) {
	// 1. Hash with SHA256
	h := sha256.Sum256(data)
	x := new(big.Int).SetBytes(h[:])
	x.Mod(x, ibosPP)

	// 2. Try-and-Increment
	for i := 0; i < 1000; i++ {
		// y^2 = x^3 + 3 mod p
		y2 := new(big.Int).Mul(x, x)
		y2.Mod(y2, ibosPP)
		y2.Mul(y2, x)
		y2.Mod(y2, ibosPP)
		y2.Add(y2, ibosB)
		y2.Mod(y2, ibosPP)

		y, hasRoot := sys.bn254Sqrt(y2)
		if hasRoot {
			return sys.pointFromXY(x, y)
		}

		x.Add(x, big.NewInt(1))
		x.Mod(x, ibosPP)
	}
	return nil, fmt.Errorf("point not found")
}

// 辅助：求模平方根
func (sys *IBOSSystem) bn254Sqrt(xx *big.Int) (*big.Int, bool) {
	y := new(big.Int).Exp(xx, ibosSqrtExp, ibosPP)
	// check
	check := new(big.Int).Mul(y, y)
	check.Mod(check, ibosPP)
	return y, check.Cmp(xx) == 0
}

// 辅助：构造 Kyber Point
func (sys *IBOSSystem) pointFromXY(x, y *big.Int) (kyber.Point, error) {
	xBytes := common.LeftPadBytes(x.Bytes(), 32)
	yBytes := common.LeftPadBytes(y.Bytes(), 32)
	// Kyber BN256 unmarshal expects 64 bytes for G1 (X||Y)
	ptBytes := append(xBytes, yBytes...)
	
	p := sys.Suite.G1().Point()
	if err := p.UnmarshalBinary(ptBytes); err != nil {
		return nil, err
	}
	return p, nil
}

// 辅助：快捷获取 G1/G2 接口
// [注] 这是导致冲突的方法，现在 struct 字段已改名，此方法保留用于获取 Group 接口
func (sys *IBOSSystem) G1() kyber.Group { return sys.Suite.G1() }
func (sys *IBOSSystem) G2() kyber.Group { return sys.Suite.G2() }
