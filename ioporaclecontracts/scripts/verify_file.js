// 文件路径: ioporaclecontracts/scripts/verify_file.js
const fs = require('fs');
const path = require('path');
const OracleContract = artifacts.require("OracleContract");

module.exports = async function(callback) {
  try {
    // 1. 设置要验证的文件名
    const fileName = "my_document.txt";
    const filePath = path.join(__dirname, fileName);

    console.log("============================================");
    console.log("📂 正在读取文件:", fileName);

    // 2. 读取文件内容
    if (!fs.existsSync(filePath)) {
      throw new Error(`找不到文件: ${filePath}，请先创建它！`);
    }
    const fileContent = fs.readFileSync(filePath, 'utf8');
    console.log("📄 文件内容摘要:", fileContent.substring(0, 50) + "...");

    // 3. 计算文件的 Hash (这是关键！论文说只对 Hash 签名)
    // 使用 Keccak256 (Web3 标准)
    const fileHash = web3.utils.sha3(fileContent);
    console.log("🔐 文件数字指纹 (Hash):", fileHash);

    // 4. 获取合约并发送请求
    const oracle = await OracleContract.deployed();
    const fee = web3.utils.toWei("0.0012", "ether");

    console.log("--------------------------------------------");
    console.log("🚀 发起 IBOS 链式签名流程...");
    
    // 发送交易：请求大家对这个文件的 Hash 进行签名确认
    const tx = await oracle.validateBlock(fileHash, { value: fee });

    console.log("✅ 请求已发送至区块链!");
    console.log("🔗 交易哈希:", tx.tx);
    console.log("============================================");
    console.log("👀 请观察 Node 1, 2, 3 的窗口...");
    console.log("   当所有节点都打印 '✅ 验证成功' 时，");
    console.log("   意味着该文件已通过全网签名验证！");

  } catch (error) {
    console.error("❌ 发生错误:", error);
  }
  
  callback();
};
