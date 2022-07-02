// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

function wait(ms) {
  return new Promise(resolve =>setTimeout(() =>resolve(), ms));
};

async function deployFilterProxy(contractMap) {
  const FilterProxy = await hre.ethers.getContractFactory("FilterProxy");
  const filterProxy = await FilterProxy.deploy();
  await filterProxy.deployed();
  console.log("deploy filterProxy address:", filterProxy.address);

  var tx = await filterProxy.setnodecontract(contractMap.get("hpbnodes").address);
  await tx.wait();
  console.log("set node contract succeed");

  tx = await filterProxy.setvotecontract(contractMap.get("hpbvotes").address);
  await tx.wait();
  console.log("set vote contract succeed");

  tx = await filterProxy.setlockcontract(contractMap.get("hpblocks").address);
  await tx.wait();
  console.log("set lock contract succeed");

  contractMap.set("filterproxy", filterProxy);

  return contractMap;
}

async function initaddr() {
  var hpbNodes = "0x451d785A0379E637d17C1C0E96cA150168A5Ab9A";
  var hpbLocks = "0x74054455F954E1DDAf0694d906e4e68D63a33A18";
  var hpbVotes = "0x35a3445C0ca0B01B7CEA2F867D762f6410c9e952";
  var proxy = "0x60a3698bE1493da2065E6F84B2E77B5b5D201D5D";
  var filterproxy = "0x5D60b3bf2483FB19d0a4FE76Da8857e10AE13201";

  var contractMap = new Map();

  const HpbNodes = await hre.ethers.getContractAt("HpbNodes", hpbNodes);
	const HpbLocks = await hre.ethers.getContractAt("HpbLock", hpbLocks);
	const HpbVotes = await hre.ethers.getContractAt("HpbVote", hpbVotes);
	const HpbProxy = await hre.ethers.getContractAt("Proxy", proxy);
	const FilterProxy = await hre.ethers.getContractAt("FilterProxy", filterproxy);
	contractMap.set("hpbnodes", HpbNodes);
  contractMap.set("hpblocks", HpbLocks);
  contractMap.set("hpbvotes", HpbVotes);
  contractMap.set("hpbproxy", HpbProxy);
  contractMap.set("filterproxy", FilterProxy);
  return contractMap;
}

async function addInvalidNode(contract, address) {
  var tx = await contract.addInvalidNode(address,{gasLimit:1000000, gasPrice:1000000000});
  await tx.wait();
  console.log("add invalid node ", address);
}

async function testGetAllHpbNodes(contractMap) {

	await addInvalidNode(contractMap.get("filterproxy"),"0xAe43e6D43FD29100fAca6d2AeD1dFBF5EdcBD41B");
	await addInvalidNode(contractMap.get("filterproxy"),"0xF583c835961e669556448B80E4e0e491c7eE6C26");
	{
		var [coinbases,cid1s,cid2s,hids] = await contractMap.get("filterproxy").getAllHpbNodes();
		for(var i = 0; i < coinbases.length; i++) {
			console.log("filter proxy value at index ", i, " = ", coinbases[i], cid1s[i], cid2s[i], hids[i]);
		}  
	}
	{
		var [coinbases,cid1s,cid2s,hids] = await contractMap.get("hpbproxy").getAllHpbNodes();
		for(var i = 0; i < coinbases.length; i++) {
			console.log("hpb proxy value at index ", i, " = ", coinbases[i], cid1s[i], cid2s[i], hids[i]);
		}
	}

}

async function testFetchAllVoteResult(contractMap) {

	// await addInvalidNode(contractMap.get("filterproxy"),"0xAe43e6D43FD29100fAca6d2AeD1dFBF5EdcBD41B");
	// await addInvalidNode(contractMap.get("filterproxy"),"0xF583c835961e669556448B80E4e0e491c7eE6C26");
	{
		var [coinbases,votes] = await contractMap.get("filterproxy").fetchAllVoteResult();
		for(var i = 0; i < coinbases.length; i++) {
			console.log("filter proxy value at index ", i, " = ", coinbases[i], votes[i]);
		}  
	}
	{
		var [coinbases,votes] = await contractMap.get("hpbproxy").fetchAllVoteResult();
		for(var i = 0; i < coinbases.length; i++) {
			console.log("hpb proxy value at index ", i, " = ", coinbases[i], votes[i]);
		}
	}
}

async function testFetchAllHolderAddrs(contractMap) {

	// await addInvalidNode(contractMap.get("filterproxy"),"0xAe43e6D43FD29100fAca6d2AeD1dFBF5EdcBD41B");
	// await addInvalidNode(contractMap.get("filterproxy"),"0xF583c835961e669556448B80E4e0e491c7eE6C26");
	{
		var [coinbases,holders] = await contractMap.get("filterproxy").fetchAllHolderAddrs();
		for(var i = 0; i < coinbases.length; i++) {
			console.log("filter proxy value at index ", i, " = ", coinbases[i], holders[i]);
		}  
	}
	{
		var [coinbases,holders] = await contractMap.get("hpbproxy").fetchAllHolderAddrs();
		for(var i = 0; i < coinbases.length; i++) {
			console.log("hpb proxy value at index ", i, " = ", coinbases[i], holders[i]);
		}
	}
}



async function main() {
	var contractMap = await initaddr();
	contractMap = await deployFilterProxy(contractMap);

    await testGetAllHpbNodes(contractMap);
    await testFetchAllVoteResult(contractMap);
    await testFetchAllHolderAddrs(contractMap);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

  