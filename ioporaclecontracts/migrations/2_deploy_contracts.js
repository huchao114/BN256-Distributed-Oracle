const RegistryContract = artifacts.require("RegistryContract");
const OracleContract = artifacts.require("OracleContract");
const DistKeyContract = artifacts.require("DistKeyContract");

module.exports = async function (deployer) {
  // 1. 部署 DistKeyContract
  await deployer.deploy(DistKeyContract);
  const distKey = await DistKeyContract.deployed();

  // 2. 部署 RegistryContract
  await deployer.deploy(RegistryContract, DistKeyContract.address);
  const registry = await RegistryContract.deployed();

  // 3. 设置 Registry 地址到 DistKey 中
  // 这一步非常重要，必须等待它完成
  await distKey.setRegistryContract(RegistryContract.address);

  // 4. 部署 OracleContract
  await deployer.deploy(
    OracleContract,
    RegistryContract.address,
    DistKeyContract.address
  );
};
