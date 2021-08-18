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

  var addtx = await proxy.addAdmin("0xc1238F78df709AFaA8BFD9F80F1747c5b8177479");
  await addtx.wait();

  var admins = await proxy.getAdmins();
  console.log("total admins count ", admins.length)
  for (i=0;i<admins.length;i++) {
    console.log("have admin ", admins[i])
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
