// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

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
  console.log("hpbNodes deployed to:", hpbNodes.address);

  const hpbLocks = await HpbLocks.deploy();
  await hpbLocks.deployed();
  console.log("hpbLocks deployed to:", hpbLocks.address);

  const hpbVotes = await HpbVotes.deploy();
  await hpbVotes.deployed();
  console.log("hpbVotes deployed to:", hpbVotes.address);
  
  await hpbNodes.setLockContract(hpbLocks.address);

  await hpbLocks.setNodeContract(hpbNodes.address);

  await hpbVotes.setNodeContract(hpbNodes.address);

  const Proxy = await hre.ethers.getContractFactory("Proxy");
  const proxy = await Proxy.deploy();
  await proxy.deployed();
  console.log("proxy deployed to:", proxy.address);

  await proxy.setnodecontract(hpbNodes.address);
  await proxy.setvotecontract(hpbVotes.address);
  await proxy.setlockcontract(hpbLocks.address);
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
