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

function wait(ms) {
  return new Promise(resolve =>setTimeout(() =>resolve(), ms));
};

describe("HpbLock", function () {
  it("deploy hpbnodes and hpblock should successful", async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    const hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();

    const HpbLocks = await hre.ethers.getContractFactory("HpbLock");
    const hpbLocks = await HpbLocks.deploy();
    await hpbLocks.deployed();

    await hpbLocks.setNodeContract(hpbNodes.address);
    expect(await hpbLocks.getNodeContract()).to.equal(hpbNodes.address);

    await hpbNodes.setLockContract(hpbLocks.address);
    expect(await hpbNodes.getLockContract()).to.equal(hpbLocks.address);

    // add boeNodes
    const accounts = await hre.ethers.getSigners();
    //console.log("accounts = ", accounts)
    let boeNodes = accounts[0].address;
    await addOneBoe(hpbNodes, boeNodes);
    expect(await hpbNodes.isBoeNode(boeNodes)).to.true
    expect(await hpbNodes.isLockNode(boeNodes)).to.false
  });

  it("lock boeNode and check", async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    const hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();

    const HpbLocks = await hre.ethers.getContractFactory("HpbLock");
    const hpbLocks = await HpbLocks.deploy();
    await hpbLocks.deployed();

    await hpbLocks.setNodeContract(hpbNodes.address);
    expect(await hpbLocks.getNodeContract()).to.equal(hpbNodes.address);

    await hpbNodes.setLockContract(hpbLocks.address);
    expect(await hpbNodes.getLockContract()).to.equal(hpbLocks.address);

    const [owner] = await hre.ethers.getSigners();
    await addOneBoe(hpbNodes, owner.address);
    expect(await hpbNodes.isBoeNode(owner.address)).to.true;
    expect(await hpbNodes.isLockNode(owner.address)).to.false;
    // call stake with value.
    await hpbLocks.stake(owner.address, { value: ethers.utils.parseEther('30000') });
  });

  it("stake and withdraw", async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    const hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();

    const HpbLocks = await hre.ethers.getContractFactory("HpbLock");
    const hpbLocks = await HpbLocks.deploy();
    await hpbLocks.deployed();

    await hpbLocks.setNodeContract(hpbNodes.address);
    expect(await hpbLocks.getNodeContract()).to.equal(hpbNodes.address);

    await hpbNodes.setLockContract(hpbLocks.address);
    expect(await hpbNodes.getLockContract()).to.equal(hpbLocks.address);

    const accounts = await hre.ethers.getSigners();
    let boeNodes = accounts[0].address;
    await addOneBoe(hpbNodes, boeNodes);
    expect(await hpbNodes.isBoeNode(boeNodes)).to.true;
    expect(await hpbNodes.isLockNode(boeNodes)).to.false;
    
    // todo: call stake with value.
    expect(await hpbLocks.stake(boeNodes, { value: ethers.utils.parseEther('30000') }));
    expect(await hpbNodes.isLockNode(boeNodes)).to.true;

    // withdraw
    await expect(hpbLocks.withdraw(boeNodes)).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'not enough limit time'");
    await wait(10000); // need set `lockLimitTime = 10` in lock.sol.
    await hpbLocks.withdraw(boeNodes);
    expect(await hpbNodes.isLockNode(boeNodes)).to.false;
  });

  it("fetchLockInfo", async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    const hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();

    const HpbLocks = await hre.ethers.getContractFactory("HpbLock");
    const hpbLocks = await HpbLocks.deploy();
    await hpbLocks.deployed();

    await hpbLocks.setNodeContract(hpbNodes.address);
    expect(await hpbLocks.getNodeContract()).to.equal(hpbNodes.address);

    await hpbNodes.setLockContract(hpbLocks.address);
    expect(await hpbNodes.getLockContract()).to.equal(hpbLocks.address);

    const accounts = await hre.ethers.getSigners();
    let boeNodes = accounts[0].address;

    const [lockAddr,amount] = await hpbLocks.fetchLockInfoByNodeAddr(boeNodes);
    expect(lockAddr).equal("0x0000000000000000000000000000000000000000");
    expect(amount).equal(0);

    await addOneBoe(hpbNodes, boeNodes);
    expect(await hpbNodes.isBoeNode(boeNodes)).to.true;
    expect(await hpbNodes.isLockNode(boeNodes)).to.false;
    
    // todo: call stake with value.
    expect(await hpbLocks.stake(boeNodes, { value: ethers.utils.parseEther('30002') }));

    const [nlockAddr,namount] = await hpbLocks.fetchLockInfoByNodeAddr(boeNodes);
    expect(nlockAddr).equal(boeNodes);
    //console.log("namount is ", namount);
    expect(namount).equal("30000000000000000000000");

  });
});