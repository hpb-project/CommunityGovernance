const { expect } = require("chai");

describe("Proxy", function () {
  it("set and get addr", async function () {
    const BlockSet = await ethers.getContractFactory("BlockSet");
    const blockSet = await BlockSet.deploy();
    await blockSet.deployed();

    const Proxy = await ethers.getContractFactory("Proxy");
    const proxy = await proxy.deploy();
    await proxy.deployed();

    var txset = await proxy.setcontract(blockSet.address);
    await txset.wait();

    var addr = await proxy.getcontract();
    expect(addr == blockSet.address);
  });

  it("set value and get from proxy", async function () {
      const BlockSet = await ethers.getContractFactory("BlockSet");
      const blockSet = await BlockSet.deploy();
      await blockSet.deployed();

      const Proxy = await ethers.getContractFactory("Proxy");
      const proxy = await proxy.deploy();
      await proxy.deployed();

      var txset = await proxy.setcontract(blockSet.address);
      await txset.wait();

      const addTx = await blockSet.addProposal("test1",1000);
      await addTx.wait();

      var get = await proxy.getValue("test1");
      expect(get == 1000);

      get = await blockSet.getValue("test1");
      expect(get == 1000);
  });

});
