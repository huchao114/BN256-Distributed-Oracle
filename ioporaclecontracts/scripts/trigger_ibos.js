// 文件路径: ioporaclecontracts/scripts/trigger_ibos.js

const OracleContract = artifacts.require("OracleContract");

module.exports = async function(callback) {
  try {
    // 1. 获取已部署的合约实例
    const oracle = await OracleContract.deployed();
    
    // 2. 准备参数
    // 费用必须 >= BASE_FEE (0.001) + VALIDATOR_FEE (0.0001) = 0.0011 ETH
    // 这里我们发送 0.0012 ETH 以确保足够
    const fee = web3.utils.toWei("0.0012", "ether");
    
    // 生成一个随机的 Hash 进行测试 (加入时间戳保证每次不同)
    const testContent = "Test Document IBOS " + Date.now();
    const testHash = web3.utils.sha3(testContent);
    
    console.log("============================================");
    console.log("🚀 Starting IBOS Trigger Script");
    console.log("--------------------------------------------");
    console.log("📍 Oracle Contract:", oracle.address);
    console.log("📄 Test Hash:", testHash);
    console.log("💰 Sending Fee:", web3.utils.fromWei(fee, "ether"), "ETH");

    // 3. 调用合约: validateBlock
    // 这会触发 ValidationRequest 事件，Go 节点监听到后会开始工作
    const tx = await oracle.validateBlock(testHash, { value: fee });

    console.log("--------------------------------------------");
    console.log("✅ Transaction Successful!");
    console.log("🔗 Tx Hash:", tx.tx);
    console.log("============================================");
    console.log("👉 现在请立即查看 Node 1 (Aggregator) 的终端日志！");
    console.log("   你应该能看到 'Received ValidationRequest' 和 'Starting IBOS sequence'...");
    
  } catch (error) {
    console.error("❌ Error executing script:", error);
  }
  
  // 结束脚本
  callback();
};
