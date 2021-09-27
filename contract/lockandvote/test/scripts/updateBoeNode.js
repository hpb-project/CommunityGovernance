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

  var coinbase = "0x74288994a9a79f9e34dd7662b54df5964c156fac";
  var cid1 = "0x78efef8b21c6acf9fc9b7164876f1b343496ef9a7f746bf7dff418152f7556c2";
  var cid2 = "0xbcfe2789958cf79e915512c131bfb840fd4aac79acc793a601996e4d2d2e3d1c";
  var hid = "0xb173cd1773797db7d60b5b7406ecbbb091b735a79b9e07371f5265659a9f9370";

  const nodeContract = await ethers.getContractAt("HpbNodes",node);

  var updateTx = await nodeContract.updateBoeNode(coinbase, cid1, cid2, hid);
  console.log("update node", coinbase, " txhash: ",updateTx.hash)
  await updateTx.wait();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
