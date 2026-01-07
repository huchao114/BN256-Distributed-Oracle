
const OracleContract = artifacts.require("OracleContract");
module.exports = async function(callback) {
  try {
    const fileHash = process.env.TARGET_HASH;
    let sStr = process.env.TARGET_S || ""; 
    let rStr = process.env.TARGET_R || ""; 

    if (!fileHash || !sStr || !rStr) {
        console.error("❌ 缺少签名数据 (Hash/S/R)，请先执行步骤9！");
        return callback();
    }

    const oracle = await OracleContract.deployed();
    const accounts = await web3.eth.getAccounts();
    const funder = accounts[0];

    // === 1. 自动为合约充值 (解决 Transfer Failed) ===
    // 检查合约余额，如果不足 1 ETH，则充值 10 ETH
    const balance = await web3.eth.getBalance(oracle.address);
    if (web3.utils.toBN(balance).lt(web3.utils.toBN(web3.utils.toWei("1", "ether")))) {
        console.log("💰 正在为合约充值 10 ETH 以便支付奖励...");
        await web3.eth.sendTransaction({
            from: funder,
            to: oracle.address,
            value: web3.utils.toWei("10", "ether")
        });
        console.log("✅ 合约充值成功。");
    }

    // === 2. 数据清洗与切分 ===
    const sClean = sStr.replace(/0x/g, "").replace(/[^0-9a-fA-F]/g, "");
    const rClean = rStr.replace(/0x/g, "").replace(/[^0-9a-fA-F]/g, "");
    
    const sigStruct = {
        S: [
            "0x" + (sClean.substring(0, 64) || "0"),
            "0x" + (sClean.substring(64, 128) || "0")
        ],
        R: [
            "0x" + (rClean.substring(0, 64) || "0"),
            "0x" + (rClean.substring(64, 128) || "0"),
            "0x" + (rClean.substring(128, 192) || "0"),
            "0x" + (rClean.substring(192, 256) || "0")
        ]
    };

    console.log("============================================");
    console.log("🚀 正在上链存证 (submitBlockValidationIBOS)");
    console.log("--------------------------------------------");
    console.log("📄 文件哈希:", fileHash);

    const tx = await oracle.submitBlockValidationIBOS(
        fileHash,
        true, 
        [web3.utils.utf8ToHex("Node1")], 
        [sigStruct], 
        Date.now() 
    );

    console.log("✅ 上链成功! 交易哈希:", tx.tx);
    console.log("FINAL_SUCCESS_FLAG"); 

  } catch (error) {
    console.error("❌ 错误:", error);
  }
  callback();
};
