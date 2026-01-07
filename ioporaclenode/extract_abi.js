const fs = require('fs');
// 读取 Truffle 编译出的 JSON
const contract = JSON.parse(fs.readFileSync('../ioporaclecontracts/build/contracts/OracleContract.json', 'utf8'));
// 只写入 ABI 部分
fs.writeFileSync('OracleContract.abi', JSON.stringify(contract.abi));
console.log("ABI extracted to OracleContract.abi");
