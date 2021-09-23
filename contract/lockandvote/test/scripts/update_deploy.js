// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

function wait(ms) {
  return new Promise(resolve =>setTimeout(() =>resolve(), ms));
};

const proxyAddr = "0x60a3698bE1493da2065E6F84B2E77B5b5D201D5D"

async function main() {
  // get deployed proxy contract.
  const proxy = await ethers.getContractAt("Proxy",proxyAddr);
  console.log("proxy deployed at:", proxy.address);

  // get current node/lock/vote contract address.
  var [node,lock,vote] = await proxy.getcontract();
  console.log("current node address is:", node);
  console.log("current lock address is:", lock);
  console.log("current vote address is:", vote);
  
  // deploy new hpbNodes contract.
  const HpbNodes = await ethers.getContractFactory("HpbNodes");
  const newHpbNodes = await HpbNodes.deploy();
  await newHpbNodes.deployed();
  console.log("new hpbNodes address:", newHpbNodes.address);

  const locksContract = await ethers.getContractAt("HpbLock", lock);
  const votesContract = await ethers.getContractAt("HpbVote", vote);

  // set lock contract.
  var t1 = await newHpbNodes.setLockContract(locksContract.address);
  await t1.wait();

  // update nodes contract.
  var t2 = await locksContract.setNodeContract(newHpbNodes.address);
  await t2.wait();

  var t3 = await votesContract.setNodeContract(newHpbNodes.address);
  await t3.wait();

  // reset address to proxy.
  var t4 = await proxy.setnodecontract(newHpbNodes.address);
  await t4.wait();
  var t5 = await proxy.setvotecontract(votesContract.address);
  await t5.wait();
  var t6 = await proxy.setlockcontract(locksContract.address);
  await t6.wait();

  console.log("after set contracts to proxy, deploy finished.")
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
