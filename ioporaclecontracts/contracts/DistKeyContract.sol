// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./RegistryContract.sol";

contract DistKeyContract {
    // ============================================================
    // 说明：
    // - 本合约默认“锁死”DKG，使用硬编码 master pubkey（G2）
    // - 仅当 admin 显式开启 dkgEnabled 后，才允许 generate / setPublicKey / dispute
    // - publicKey 的存储顺序固定为“合约侧 pairing 输入顺序（你的系统约定）”：
    //      [X_IM, X_RE, Y_IM, Y_RE]
    //
    // ✅ 关键修正（消除 IM/RE 认知反转）：
    // 你当前 Kyber MarshalBinary 的 G2 前 64 bytes 是：
    //   rb[0:32] = 0x198e...   rb[32:64] = 0x1800...
    // 因此在本系统中，将 0x198e 视为 X_IM，将 0x1800 视为 X_RE；
    // 同理，将 0x0906 视为 Y_IM，将 0x12c8 视为 Y_RE。
    // 这样链上“命名”和 Go 侧“实际编码”一致，避免靠 swap 兜底。
    // ============================================================

    // 禁用自动 DKG，保护硬编码公钥
    uint256 public constant KEY_GEN_INTERVAL = 999;
    uint256 public constant KEY_FINAL_TIME = 0;

    // BN254 / alt_bn128 base field prime
    uint256 internal constant BN254_P =
        21888242871839275222246405745257275088696311157297823662689037894645226208583;

    event DistKeyGenerationLog(uint256 threshold);
    event DKGEnabledChanged(bool enabled);
    event RegistryContractSet(address registry);
    event PublicKeyUpdated(uint256[4] pubkey, uint256 numberOfValidators);

    address public immutable admin;

    uint256 private signatureThreshold;
    uint256 private validatorThreshold;
    bool private generating;

    // G2 public key in contract order [X_IM, X_RE, Y_IM, Y_RE]
    uint256[4] private publicKey;
    uint256 private numberOfValidators;

    address[] private disputes;
    uint256 private publicKeyFinal;

    RegistryContract private registryContract;

    // 默认锁死，防止 dispute / setPublicKey 改掉
    bool public dkgEnabled;

    constructor() {
        admin = msg.sender;

        // ============================================================
        // ✅ 修正后的硬编码 G2 Generator (Priv=1)，按你系统约定存储：
        // publicKey = [X_IM, X_RE, Y_IM, Y_RE]
        //
        // 其中：
        // X_IM = 0x198e...  X_RE = 0x1800...
        // Y_IM = 0x0906...  Y_RE = 0x12c8...
        // ============================================================
        publicKey[0] = 0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2; // X_IM
        publicKey[1] = 0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed; // X_RE
        publicKey[2] = 0x090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b; // Y_IM
        publicKey[3] = 0x12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa; // Y_RE

        _validateG2Limbs(publicKey);

        // 立即 final（因为 KEY_FINAL_TIME=0）
        publicKeyFinal = 0;
        numberOfValidators = 1;

        // 默认禁用 DKG（锁死）
        dkgEnabled = false;
        generating = false;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin, "not admin");
        _;
    }

    function setDKGEnabled(bool enabled) external onlyAdmin {
        // 强制要求 registry 已设置，避免误开启后被 0 地址卡死
        require(address(registryContract) != address(0), "registry not set");
        dkgEnabled = enabled;

        // 关闭 DKG 时，清理 generating 状态与争议
        if (!enabled) {
            generating = false;
            delete disputes;
        }

        emit DKGEnabledChanged(enabled);
    }

    function setRegistryContract(address _registryContract) external onlyAdmin {
        require(_registryContract != address(0), "registry=0");
        require(address(registryContract) == address(0), "registry already set");
        registryContract = RegistryContract(_registryContract);
        emit RegistryContractSet(_registryContract);
    }

    // Registry 合约触发 key generation
    function generate() public {
        require(msg.sender == address(registryContract), "invalid sender");
        require(dkgEnabled, "DKG disabled");
        _generateKey();
    }

    function _generateKey() private {
        uint256 noNodes = registryContract.countOracleNodes();
        require(noNodes > 0, "no nodes");

        signatureThreshold = (noNodes + 1) / 2;
        validatorThreshold = signatureThreshold + (noNodes - signatureThreshold) / 2;

        generating = true;
        emit DistKeyGenerationLog(signatureThreshold);
    }

    function getNumberOfValidators() external view returns (uint256) {
        return numberOfValidators;
    }

    function getPublicKey() external view returns (uint256[4] memory) {
        require(block.number >= publicKeyFinal, "pubKey not final");
        return publicKey;
    }

    function setPublicKey(uint256[4] calldata _publicKey, uint256 _numberOfValidators) external {
        require(dkgEnabled, "DKG disabled");
        require(generating, "not generating");
        require(address(registryContract) != address(0), "registry not set");
        require(registryContract.oracleNodeIsRegistered(msg.sender), "not registered");
        require(_numberOfValidators >= validatorThreshold, "too few validators");

        // 关键：防止传入 >= p 的 limb 导致 pairing precompile 直接失败
        _validateG2Limbs(_publicKey);

        publicKey = _publicKey;
        numberOfValidators = _numberOfValidators;

        // KEY_FINAL_TIME=0 => 立即 final
        publicKeyFinal = block.number + KEY_FINAL_TIME;

        generating = false;
        delete disputes;

        emit PublicKeyUpdated(_publicKey, _numberOfValidators);
    }

    function disputePublicKey() external {
        require(dkgEnabled, "DKG disabled");
        require(KEY_FINAL_TIME > 0, "dispute disabled"); // KEY_FINAL_TIME=0 时 dispute 无意义
        require(address(registryContract) != address(0), "registry not set");
        require(registryContract.oracleNodeIsRegistered(msg.sender), "not registered");
        require(block.number < publicKeyFinal, "key already final");

        for (uint256 i = 0; i < disputes.length; i++) {
            require(disputes[i] != msg.sender, "dispute already issued");
        }
        disputes.push(msg.sender);

        if (disputes.length >= signatureThreshold) {
            delete publicKeyFinal;
            delete disputes;
            _generateKey();
        }
    }

    // ============================================================
    // 内部校验：G2 limbs 必须 < p，且不能全 0
    // （不做完整 on-curve/subgroup 检查；这里只排除最常见的 precompile 失败输入）
    // ============================================================
    function _validateG2Limbs(uint256[4] memory pk) internal pure {
        require(pk[0] < BN254_P, "pk[0] >= p");
        require(pk[1] < BN254_P, "pk[1] >= p");
        require(pk[2] < BN254_P, "pk[2] >= p");
        require(pk[3] < BN254_P, "pk[3] >= p");
        require(!(pk[0] == 0 && pk[1] == 0 && pk[2] == 0 && pk[3] == 0), "pk is zero");
    }
}

