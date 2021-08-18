// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const proxyAddr = "0x60a3698bE1493da2065E6F84B2E77B5b5D201D5D"
async function main() {
  const proxy = await ethers.getContractAt("Proxy",proxyAddr);
  console.log("proxy deployed at:", proxy.address);
  var [node,lock,vote] = await proxy.getcontract();
  console.log("current node address is:", node);
  console.log("current lock address is:", lock);
  console.log("current vote address is:", vote);

  const nodeContract = await ethers.getContractAt("HpbNodes",node);
  //node.addBoeNode.sendTransaction("0xad535af4077ffa81d47b9f9fc23987152907f3c7","0x4fb5ef40bebdc0202a816b1fff642a3bf7e6edfed741ef34d3cdd43c63bc6615","0xf7603641b281c35e0c9f6f1d5b758c09b1d1d65651a4bbf507004212b8543065","0x37094ba5328e9259c99517a3765bb3138d5a9ed3f90f49ab88dc99b914f7dfb2",{from:hpb.coinbase[1],gas:5000000})
  var addtx = await nodeContract.addBoeNode("0xad535af4077ffa81d47b9f9fc23987152907f3c7","0x4fb5ef40bebdc0202a816b1fff642a3bf7e6edfed741ef34d3cdd43c63bc6615","0xf7603641b281c35e0c9f6f1d5b758c09b1d1d65651a4bbf507004212b8543065","0x37094ba5328e9259c99517a3765bb3138d5a9ed3f90f49ab88dc99b914f7dfb2");
  await addtx.wait();

  var boeaddrs = await nodeContract.getAllBoesAddrs();
  console.log("total boeaddrs count ", boeaddrs.length);
  for (i=0;i<boeaddrs.length;i++) {
    console.log("have boe ", boeaddrs[i])
  }
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
