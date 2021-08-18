// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

function wait(ms) {
  return new Promise(resolve =>setTimeout(() =>resolve(), ms));
};

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  // We get the contract to deploy
  const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
  const HpbLocks = await hre.ethers.getContractFactory("HpbLock");
  const HpbVotes = await hre.ethers.getContractFactory("HpbVote");

  const hpbNodes = await HpbNodes.deploy();
  await hpbNodes.deployed();
  console.log("hpbNodes address:", hpbNodes.address);

  const hpbLocks = await HpbLocks.deploy();
  await hpbLocks.deployed();
  console.log("hpbLocks address:", hpbLocks.address);

  const hpbVotes = await HpbVotes.deploy();
  await hpbVotes.deployed();
  console.log("hpbVotes address:", hpbVotes.address);
  
  var t1 = await hpbNodes.setLockContract(hpbLocks.address);
  await t1.wait();

  var t2 = await hpbLocks.setNodeContract(hpbNodes.address);
  await t2.wait();

  var t3 = await hpbVotes.setNodeContract(hpbNodes.address);
  await t3.wait();

  const Proxy = await hre.ethers.getContractFactory("Proxy");
  const proxy = await Proxy.deploy();
  await proxy.deployed();
  console.log("proxy address:", proxy.address);
  var t4 = await proxy.setnodecontract(hpbNodes.address);
  await t4.wait();
  var t5 = await proxy.setvotecontract(hpbVotes.address);
  await t5.wait();
  var t6 = await proxy.setlockcontract(hpbLocks.address);
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
