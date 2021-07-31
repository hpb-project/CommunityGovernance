const { expect } = require("chai");

async function addOneBoe(hpbNodes, coinbase) {
  let cid1 = "0xaf4de7c3e01a02111cd9acbddbb8badfaf9f1132d41ee29884b2bcf080568c34";
  let cid2 = "0xe9082d775021ebca0faa25a753dd1e636d3b4f7f69336cf44cffea03b68c783c";
  let hid  = "0x0a5fe66818ae1088a40039d5531789feb8e60e2196db83816c9f38df4a5c3ad7";

  await hpbNodes.addBoeNode(coinbase, cid1, cid2, hid);
}

async function updateOneBoe(hpbNodes, coinbase) {
  let cid1 = "0xadfaf9f1132d41ee29884b2bcf080568c34af4de7c3e01a02111cd9acbddbb8b";
  let cid2 = "0xb4f7f69336cf44cffea03b68c783ce9082d775021ebca0faa25a753dd1e636d3";
  let hid  = "0x9feb8e60e2196db83816c9f38df4a5c3ad70a5fe66818ae1088a40039d553178";

  await hpbNodes.updateBoeNode(coinbase, cid1, cid2, hid);
}

function checkBoeUpdated(cid1,cid2,hid) {
  let ucid1 = "0xadfaf9f1132d41ee29884b2bcf080568c34af4de7c3e01a02111cd9acbddbb8b";
  let ucid2 = "0xb4f7f69336cf44cffea03b68c783ce9082d775021ebca0faa25a753dd1e636d3";
  let uhid  = "0x9feb8e60e2196db83816c9f38df4a5c3ad70a5fe66818ae1088a40039d553178";

  //console.log("cid1 equal ? ", cid1, ucid1, cid1==ucid1);

  return (ucid1 == cid1 && ucid2 == cid2 && uhid == hid)
}

async function delOneBoe(hpbNodes, coinbase) {
  await hpbNodes.deleteBoeNode(coinbase);
}


describe("HpbNodes", function () {
  it("add boeNode should successful", async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    const hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();
    let coinbase = ethers.utils.getAddress("0x7cc4e0cc326756155d4afb2e780bdd13a17d4434");

    await addOneBoe(hpbNodes, coinbase);

    expect(await hpbNodes.isBoeNode(coinbase)).to.true;
    expect(await hpbNodes.isLockNode(coinbase)).to.false;
  });

  // it("delete boeNode should be reverted", async function () {
  //   const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
  //   const hpbNodes = await HpbNodes.deploy();
  //   await hpbNodes.deployed();
  //   let coinbase = ethers.utils.getAddress("0x7cc4e0cc326756155d4afb2e780bdd13a17d4434");
  
  //   await expect(await hpbNodes.deleteBoeNode(coinbase)).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'not exist'");

  //   //await expect().to.revertedWith('/not exist/');
  // });

  it("delete boeNode should successful", async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    const hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();
    let coinbase = ethers.utils.getAddress("0x7cc4e0cc326756155d4afb2e780bdd13a17d4434");

    await addOneBoe(hpbNodes, coinbase);
    expect(await hpbNodes.isBoeNode(coinbase)).to.true;

    await delOneBoe(hpbNodes, coinbase);
    expect(await hpbNodes.isBoeNode(coinbase)).to.false;
  });

  it("update boeNode should successful", async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    const hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();
    let coinbase = ethers.utils.getAddress("0x7cc4e0cc326756155d4afb2e780bdd13a17d4434");

    await addOneBoe(hpbNodes, coinbase);
    expect(await hpbNodes.isBoeNode(coinbase)).to.true;

    await updateOneBoe(hpbNodes, coinbase);
    
    var [coinbases,cid1s,cid2s,hids] = await hpbNodes.getAllBoes();
    expect(coinbases.length == 1);
    // console.log("get cid1s=", cid1s);
    // console.log("get cid2s=", cid2s);
    // console.log("get hids =", hids);

    expect(checkBoeUpdated(cid1s[0], cid2s[0], hids[0])).to.true;
    expect(await hpbNodes.isBoeNode(coinbase)).to.true;

  });

  it("set and get lockContract", async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    const hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();

    const HpbLocks = await hre.ethers.getContractFactory("HpbLock");
    const hpbLocks = await HpbLocks.deploy();
    await hpbLocks.deployed();
    //console.log("hpbLocks deployed to:", hpbLocks.address);

    await hpbNodes.setLockContract(hpbLocks.address);
    expect(await hpbNodes.getLockContract()).to.equal(hpbLocks.address);
  });


});


describe("HpbNodes Holders tests", function () {
  it("bind boe and holder addr", async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    const hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();
    const [boeaddr,holderaddr] = await hre.ethers.getSigners();

    await addOneBoe(hpbNodes, boeaddr.address); // add boe
    // bind holder
    await hpbNodes.setHolderAddr(holderaddr.address);

    // verify boeaddr
    await hpbNodes.connect(holderaddr).setHolderBoeAddr(boeaddr.address);

    // check
    let [boes,holders] = await hpbNodes.fetchAllHolderAddrs();
    expect(boes[0]).to.equal(boeaddr.address);
    expect(holders[0]).to.equal(holderaddr.address);

  });

});