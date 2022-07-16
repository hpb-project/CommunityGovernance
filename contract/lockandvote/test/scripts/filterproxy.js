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
  var filterproxy = "0x115eB6E6DA16B8b951E1d7ef03AF466C883a362C";

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

async function rmInvalidNode(contract, address) {
  var tx = await contract.removeInvalidNode(address,{gasLimit:1000000, gasPrice:1000000000});
  await tx.wait();
  console.log("rm invalid node ", address, "hash", tx.hash);
}

async function testGetAllHpbNodes(contractMap) {

	//await rmInvalidNode(contractMap.get("filterproxy"),"0x94724bb354b29f0797019506763e14b9d586e7be");
	await addInvalidNode(contractMap.get("filterproxy"),"0x94724bb354b29f0797019506763e14b9d586e7be");
	await addInvalidNode(contractMap.get("filterproxy"),"0xd866a6884a5b4b848fe8d02a2dbc08002385be31");
	//await addInvalidNode(contractMap.get("filterproxy"),"0x99ae5303e81a4371d9043a8554e365a84566fe86");
	//await addInvalidNode(contractMap.get("filterproxy"),"0x6f75fc8088b2a1d77f1da71507f812513080084b");
	//await addInvalidNode(contractMap.get("filterproxy"),"0x86306221af2e82e1203a29188964cca48c49dd78");
	{
		var [coinbases,cid1s,cid2s,hids] = await contractMap.get("filterproxy").getAllHpbNodes();
		console.log("filter proxy get hpbnode length is ", coinbases.length);
		//for(var i = 0; i < coinbases.length; i++) {
		//	console.log("filter proxy value at index ", i, " = ", coinbases[i], cid1s[i], cid2s[i], hids[i]);
		//}  
	}
	{
		var [coinbases,cid1s,cid2s,hids] = await contractMap.get("hpbproxy").getAllHpbNodes();
		console.log("hpb proxy get hpbnode length is ", coinbases.length);
		//for(var i = 0; i < coinbases.length; i++) {
		//	console.log("hpb proxy value at index ", i, " = ", coinbases[i], cid1s[i], cid2s[i], hids[i]);
		//}
	}

}

async function testFetchAllVoteResult(contractMap) {

	// await addInvalidNode(contractMap.get("filterproxy"),"0xAe43e6D43FD29100fAca6d2AeD1dFBF5EdcBD41B");
	// await addInvalidNode(contractMap.get("filterproxy"),"0xF583c835961e669556448B80E4e0e491c7eE6C26");
	{
		var [coinbases,votes] = await contractMap.get("filterproxy").fetchAllVoteResult();
		console.log("filter proxy get vote length is", coinbases.length);
		//for(var i = 0; i < coinbases.length; i++) {
		//	console.log("filter proxy value at index ", i, " = ", coinbases[i], votes[i]);
		//}  
	}
	{
		var [coinbases,votes] = await contractMap.get("hpbproxy").fetchAllVoteResult();
		console.log("hpb proxy get vote length is", coinbases.length);
		//for(var i = 0; i < coinbases.length; i++) {
		//	console.log("hpb proxy value at index ", i, " = ", coinbases[i], votes[i]);
		//}
	}
}

async function testFetchAllHolderAddrs(contractMap) {

	// await addInvalidNode(contractMap.get("filterproxy"),"0xAe43e6D43FD29100fAca6d2AeD1dFBF5EdcBD41B");
	// await addInvalidNode(contractMap.get("filterproxy"),"0xF583c835961e669556448B80E4e0e491c7eE6C26");
	{
		var [coinbases,holders] = await contractMap.get("filterproxy").fetchAllHolderAddrs();
		console.log("filter proxy get holders length is", coinbases.length);
		//for(var i = 0; i < coinbases.length; i++) {
		//	console.log("filter proxy value at index ", i, " = ", coinbases[i], holders[i]);
		//}  
	}
	{
		var [coinbases,holders] = await contractMap.get("hpbproxy").fetchAllHolderAddrs();
		console.log("hpb proxy get holders length is", coinbases.length);
		//for(var i = 0; i < coinbases.length; i++) {
		//	console.log("hpb proxy value at index ", i, " = ", coinbases[i], holders[i]);
		//}
	}
}

async function testMaxCount(contractMap) {
	var filtercontract = contractMap.get("filterproxy");
	var blackminers = await filtercontract.getBlackList();
	var curcount = await filtercontract.curblackcount();
	console.log("got blockminers length is", blackminers.length, "and curcount is", curcount);
}

async function updateProxy(contractMap) {

	var blocknumber = await web3.eth.getBlockNumber();
	while (blocknumber < 15168205) {
		console.log("current block number is ", blocknumber.toString(), "still wait");
		blocknumber = await web3.eth.getBlockNumber();
	}
	var nonce = await web3.eth.getTransactionCount("0xb3cde0060052Fc7Ee8BfdEC5c8E251880eF1Dcbf");
	console.log("user nonce is ", nonce.toString(), "new nonce is ", nonce + 1)
	console.log("current block number is ", blocknumber.toString(), "goto send tx");
	var tx1 = contractMap.get("hpbproxy").setnodecontract(contractMap.get("filterproxy").address, {nonce:nonce+0});
	var tx2 = contractMap.get("hpbproxy").setvotecontract(contractMap.get("filterproxy").address, {nonce:nonce+1});
	var tx3 = contractMap.get("hpbproxy").setlockcontract(contractMap.get("filterproxy").address, {nonce:nonce+2});
	//var tx1 = contractMap.get("filterproxy").setnodecontract(contractMap.get("filterproxy").address, {nonce:nonce+0});
	//var tx2 = contractMap.get("filterproxy").setvotecontract(contractMap.get("filterproxy").address, {nonce:nonce+1});
	//var tx3 = contractMap.get("filterproxy").setlockcontract(contractMap.get("filterproxy").address, {nonce:nonce+2});
	await tx1;
	await tx2;
	await tx3;
	console.log("update proxy succeed");
}


async function main() {
	var contractMap = await initaddr();
	contractMap = await deployFilterProxy(contractMap);

    await testGetAllHpbNodes(contractMap);
    await testFetchAllVoteResult(contractMap);
    await testFetchAllHolderAddrs(contractMap);
    await testMaxCount(contractMap);
    //await updateProxy(contractMap);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

  
