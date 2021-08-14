// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const Web3 = require("web3");

const hre = require("hardhat");
const Proxy = require("../artifacts/contracts/Proxy.sol/Proxy.json");
//var web3 = new Web3("https://114.242.26.15:8006");
//var proxy = new hre.ethers.Contract(Proxy.abi, "0x31E035c9405950cafe24Db89c353D53714AAed53"); // testnet
var web3 = new Web3('https://hpbnode.com');
var proxy = new web3.eth.Contract(Proxy.abi, "0xA83b0Ff8075EE46795f0a100fBdba1DA6F01c04d"); // mainnet

async function main() {
  console.log("proxy deployed at:", proxy.address);
  var addr = await proxy.getcontract();
  console.log("current blockset address is:", addr);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
