// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const { boelist } = require('./boelist.json');
const proxyAddr = "0x60a3698bE1493da2065E6F84B2E77B5b5D201D5D" // mainnet
//const proxyAddr = "0x2B21E06f11eFd1272691B62428E31513bF3B6F31"  // testnet
async function main() {
  const proxy = await ethers.getContractAt("Proxy",proxyAddr);
  console.log("proxy deployed at:", proxy.address);
  var [node,lock,vote] = await proxy.getcontract();
  console.log("current node address is:", node);
  console.log("current lock address is:", lock);
  console.log("current vote address is:", vote);

  console.log("boelist length = ", boelist.length);
  //return

  const nodeContract = await ethers.getContractAt("HpbNodes",node);
  var addrs = await nodeContract.getAllBoesAddrs();
  console.log("addrs length = ", addrs.length);
  var [eCoinbase, eCid1s, eCid2s, eHids] = await nodeContract.getAllBoes();
  var existNodes = new Map();
  console.log("got from contract length = ", eCoinbase.length);
  
  for (i=0; i < eCoinbase.length; i++) {
    var coinbase = eCoinbase[i].toLowerCase()
    var node = {"coinbase":coinbase, "cid1":eCid1s[i], "cid2":eCid2s[i], "hid":eHids[i]};
    existNodes.set(coinbase, node);
  }
  existNodes.forEach(function(value, key){
    console.log("has ", key);
  }, existNodes);
  if (existNodes.has("0xad535aF4077ffa81D47b9F9Fc23987152907F3c7")) {
    console.log("hasssssssss");
  } else {
    console.log("not hassssss");
  }

  for(i=0; i < boelist.length; i++) {
    var coinbase = boelist[i].coinbase;
    var cid1 = boelist[i].cid1;
    var cid2 = boelist[i].cid2;
    var hid  = boelist[i].hid;

    if(existNodes.has(coinbase.toLowerCase())) {
      console.log("ignore existed coinbase ", coinbase);
      continue;
    }
    
    var addTx = await nodeContract.addBoeNode(coinbase, cid1, cid2, hid);
    console.log("add node", coinbase, " txhash: ",addTx.hash)
    await addTx.wait();  
  }

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
