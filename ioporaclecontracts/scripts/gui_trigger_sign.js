
const fs = require('fs');
const OracleContract = artifacts.require("OracleContract");
module.exports = async function(callback) {
  try {
    const filePath = process.env.TARGET_FILE_PATH;
    const fileContent = fs.readFileSync(filePath, 'utf8');
    const fileHash = web3.utils.sha3(fileContent);
    const fee = web3.utils.toWei("0.0012", "ether");
    const oracle = await OracleContract.deployed();
    console.log("--------------------------------------------");
    console.log("🚀 GUI Trigger: Validating File");
    console.log("📄 Hash:", fileHash);
    console.log("--------------------------------------------");
    const tx = await oracle.validateBlock(fileHash, { value: fee });
    console.log("✅ Transaction Sent! Hash:", tx.tx);
    console.log("👉 Waiting for signatures...");
  } catch (error) { console.error("❌ Error:", error); }
  callback();
};
