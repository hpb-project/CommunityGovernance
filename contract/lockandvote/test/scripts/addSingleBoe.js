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

  
  //{"coinbase":"0x5ce2d9e6d30d83c02efd38f711d5692146ed9b9e", "hid":"0xb95ff98c9ef73f274edf763915962a1d3efe21c789abdd755815eeb877bcfaf6","cid1":"0xd1e5bf293d0f37a45b78a6bb32463db5a96b17d78c65f608b2363fd40c04f5ce", "cid2":"0x732cc4cab67de2e063a1aff745c8a83ef723d1e66b24a20e91e2e8add28b7671"},
  var coinbase = "0x5ce2d9e6d30d83c02efd38f711d5692146ed9b9e";
  var cid1 = "0xd1e5bf293d0f37a45b78a6bb32463db5a96b17d78c65f608b2363fd40c04f5ce";
  var cid2 = "0x732cc4cab67de2e063a1aff745c8a83ef723d1e66b24a20e91e2e8add28b7671";
  var hid  = "0xb95ff98c9ef73f274edf763915962a1d3efe21c789abdd755815eeb877bcfaf6";

  if(existNodes.has(coinbase.toLowerCase())) {
    console.log("ignore existed coinbase ", coinbase);
    return
  }
    
  var addTx = await nodeContract.addBoeNode(coinbase, cid1, cid2, hid);
  console.log("add node", coinbase, " txhash: ",addTx.hash)
  await addTx.wait();  
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
