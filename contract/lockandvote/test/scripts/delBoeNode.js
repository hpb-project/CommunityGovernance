// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const proxyAddr = "0x60a3698bE1493da2065E6F84B2E77B5b5D201D5D" // mainnet
//const proxyAddr = "0x2B21E06f11eFd1272691B62428E31513bF3B6F31"  // testnet
async function main() {
  const proxy = await ethers.getContractAt("Proxy",proxyAddr);
  console.log("proxy deployed at:", proxy.address);
  var [node,lock,vote] = await proxy.getcontract();
  console.log("current node address is:", node);
  console.log("current lock address is:", lock);
  console.log("current vote address is:", vote);

  var coinbase = "0x6e52d862253493cd86a353cb3b46665a172854ff";

  const nodeContract = await ethers.getContractAt("HpbNodes",node);

  var delTx = await nodeContract.deleteBoeNode(coinbase);
  console.log("del node", coinbase, " txhash: ",delTx.hash)
  await delTx.wait();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
