const { expect } = require("chai");

async function addOneBoe(hpbNodes, coinbase) {
  let cid1 = "0xaf4de7c3e01a02111cd9acbddbb8badfaf9f1132d41ee29884b2bcf080568c34";
  let cid2 = "0xe9082d775021ebca0faa25a753dd1e636d3b4f7f69336cf44cffea03b68c783c";
  let hid  = "0x0a5fe66818ae1088a40039d5531789feb8e60e2196db83816c9f38df4a5c3ad7";

  await hpbNodes.addBoeNode(coinbase, cid1, cid2, hid);
}

async function delOneBoe(hpbNodes, coinbase) {
  await hpbNodes.deleteBoeNode(coinbase);
}

async function stakeSelf(hpbLock, account) {
  await hpbLock.connect(account).stake(account.address, { value: ethers.utils.parseEther('30000') });
  //console.log("stake %s with %s", account.address, account.address);
}


describe("FilterProxy Testing MaxLimit", function () {
  let accounts;
  let nodelist = ["0xD866a6884a5B4B848fE8D02a2DbC08002385BE31", "0x94724bB354B29f0797019506763e14B9d586e7Be", "0x99ae5303e81a4371d9043a8554E365A84566fe86", 
  "0x6f75fC8088B2a1d77F1da71507f812513080084B", "0x86306221aF2E82e1203A29188964cca48c49DD78", "0x6F2a171a2E9Fb0C858e7f4f144cD28C606c671fd", "0xad535aF4077ffa81D47b9F9Fc23987152907F3c7", 
  "0xb4270C32B5D4a6B16F31883DdDe380d1F5d83AF3", "0x0926F045887401883a8d9c6e9878B1b900A5be19", "0xf8282f019c753fb2410a50E2C4789bf96b706114",
  "0x405dD40AA9298448a58Fa679ec4295eAC8472a0d", "0xAe43e6D43FD29100fAca6d2AeD1dFBF5EdcBD41B"];

  beforeEach(async function () {
    // const Filter = await hre.ethers.getContractFactory("FilterProxy");
    // filter = await Filter.deploy();
    // await filter.deployed();
    // console.log("deploy filter at ", filter.address);

    accounts = await hre.ethers.getSigners();
    console.log("account is ", accounts[0].address);
  });

  // test add max limit
  it("test max limit", async function () {
    const Filter = await hre.ethers.getContractFactory("FilterProxy");
    const filter = await Filter.deploy();
    await filter.deployed();
    var tx = await filter.addInvalidNode(nodelist[0]);
    await tx.wait();
    await expect(filter.addInvalidNode(nodelist[0])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'already exist in blacklist'")
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[1]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[2]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[3]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[4]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[5]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[6]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[7]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[8]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[9]);
    await tx.wait();
    await expect(filter.addInvalidNode(nodelist[10])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'exceed max black miner count'");
    await expect(filter.removeInvalidNode(nodelist[10])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'unknown black miner'");
  });
  it("test remove unknown", async function () {
    const Filter = await hre.ethers.getContractFactory("FilterProxy");
    const filter = await Filter.deploy();
    await filter.deployed();
    var tx = await filter.addInvalidNode(nodelist[0]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(1);
    await expect(filter.removeInvalidNode(nodelist[10])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'unknown black miner'");
    expect(await filter.curblackcount()).to.equal(1);

    tx = await filter.removeInvalidNode(nodelist[0]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(0);
  });

  it("test add already exist", async function () {
    const Filter = await hre.ethers.getContractFactory("FilterProxy");
    const filter = await Filter.deploy();
    await filter.deployed();
    var tx = await filter.addInvalidNode(nodelist[0]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(1);
    await expect(filter.addInvalidNode(nodelist[0])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'already exist in blacklist'")
    expect(await filter.curblackcount()).to.equal(1);
  });

  it("test max limit", async function () {
    const Filter = await hre.ethers.getContractFactory("FilterProxy");
    const filter = await Filter.deploy();
    await filter.deployed();
    var tx = await filter.addInvalidNode(nodelist[0]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(1);
    await expect(filter.addInvalidNode(nodelist[0])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'already exist in blacklist'")
    expect(await filter.curblackcount()).to.equal(1);
    tx = await filter.addInvalidNode(nodelist[1]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(2);
    tx = await filter.addInvalidNode(nodelist[2]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(3);
    tx = await filter.addInvalidNode(nodelist[3]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(4);
    tx = await filter.addInvalidNode(nodelist[4]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(5);
    tx = await filter.addInvalidNode(nodelist[5]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(6);
    tx = await filter.addInvalidNode(nodelist[6]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(7);
    tx = await filter.addInvalidNode(nodelist[7]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(8);
    tx = await filter.addInvalidNode(nodelist[8]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(9);
    tx = await filter.addInvalidNode(nodelist[9]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(10);
    await expect(filter.addInvalidNode(nodelist[10])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'exceed max black miner count'");
    await expect(filter.removeInvalidNode(nodelist[10])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'unknown black miner'");

    await filter.removeInvalidNode(nodelist[0]);
    expect(await filter.curblackcount()).to.equal(9);
    await filter.removeInvalidNode(nodelist[1]);
    expect(await filter.curblackcount()).to.equal(8);
    await filter.removeInvalidNode(nodelist[2]);
    expect(await filter.curblackcount()).to.equal(7);
    await filter.removeInvalidNode(nodelist[3]);
    expect(await filter.curblackcount()).to.equal(6);
    await filter.removeInvalidNode(nodelist[4]);
    expect(await filter.curblackcount()).to.equal(5);
    await filter.removeInvalidNode(nodelist[5]);
    expect(await filter.curblackcount()).to.equal(4);

    tx = await filter.addInvalidNode(nodelist[0]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(5);
    await expect(filter.addInvalidNode(nodelist[0])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'already exist in blacklist'")
    expect(await filter.curblackcount()).to.equal(5);
    tx = await filter.addInvalidNode(nodelist[1]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(6);
    tx = await filter.addInvalidNode(nodelist[2]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(7);
    tx = await filter.addInvalidNode(nodelist[3]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(8);
    tx = await filter.addInvalidNode(nodelist[4]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(9);
    tx = await filter.addInvalidNode(nodelist[5]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(10);

    await expect(filter.addInvalidNode(nodelist[10])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'exceed max black miner count'");
    await expect(filter.removeInvalidNode(nodelist[10])).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'unknown black miner'");
    var miners = await filter.getBlackList();
    for(var i = 0; i < miners.length; i++) {
      console.log("black miner has ", miners[i]);
    }
  });
  
  it("test set limit", async function () {
    const Filter = await hre.ethers.getContractFactory("FilterProxy");
    const filter = await Filter.deploy();
    await filter.deployed();
    await expect(filter.setMaxLimit(100)).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'can not limit more than 50'")
    await expect(filter.setMaxLimit(9)).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'can not limit little than 10'")
    await filter.setMaxLimit(15);

    var tx = await filter.addInvalidNode(nodelist[0]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[1]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[2]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[3]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[4]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[5]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[6]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[7]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[8]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[9]);
    await tx.wait();
    tx = await filter.addInvalidNode(nodelist[10]);
    await tx.wait();
    expect(await filter.curblackcount()).to.equal(11);
    await expect(filter.setMaxLimit(10)).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'can not limit little than curblackcount'")

    await filter.removeInvalidNode(nodelist[0]);
    expect(await filter.curblackcount()).to.equal(10);
    await expect(filter.setMaxLimit(10)).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'can not limit little than curblackcount'")
    await filter.removeInvalidNode(nodelist[1]);
    expect(await filter.curblackcount()).to.equal(9);
    await filter.setMaxLimit(10);

  });


});