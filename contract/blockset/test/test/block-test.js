const { expect } = require("chai");

describe("BlockSet", function () {
  it("Should got 0 admins", async function () {
    const BlockSet = await ethers.getContractFactory("blockSet");
    const blockSet = await BlockSet.deploy();
    await blockSet.deployed();
	  var admins = await blockSet.getAdmins();
	  expect(admins.length == 0);
  });
  it("Should got 1 admins", async function () {
    const BlockSet = await ethers.getContractFactory("blockSet");
    const blockSet = await BlockSet.deploy();
    await blockSet.deployed();

    const addAdminTx = await blockSet.addAdmin("0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D");

    // wait until the transaction is mined
    await addAdminTx.wait();
	  var admins = await blockSet.getAdmins();
	  expect(admins.length == 1);
  });
  it("Should delete to 0 admins", async function () {
    const BlockSet = await ethers.getContractFactory("blockSet");
    const blockSet = await BlockSet.deploy();
    await blockSet.deployed();

    const addAdminTx = await blockSet.addAdmin("0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D");

    // wait until the transaction is mined
    await addAdminTx.wait();
	  var admins = await blockSet.getAdmins();
	  expect(admins.length == 1);

    const delAdminTx = await blockSet.deleteAdmin("0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D");

    // wait until the transaction is mined
    await delAdminTx.wait();
	  var admins = await blockSet.getAdmins();
	  expect(admins.length == 0);
  });
  it("Should change owner", async function () {
    const BlockSet = await ethers.getContractFactory("blockSet");
    const blockSet = await BlockSet.deploy();
    await blockSet.deployed();

    const changeOwnerTx= await blockSet.changeOwner("0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D");

    // wait until the transaction is mined
    await changeOwnerTx.wait();
	  var owner = await blockSet.getOwner();
	  console.log("got owner",owner);
	  //expect(owner.equal("0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D"));
  });
  it("Should change threshold", async function () {
    const BlockSet = await ethers.getContractFactory("blockSet");
    const blockSet = await BlockSet.deploy();
    await blockSet.deployed();
	  var initThreshold = await blockSet.getThreshold();
	  expect(initThreshold==1);

    const addAdminTx1 = await blockSet.addAdmin("0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D");
    await addAdminTx1.wait();
    const addAdminTx2 = await blockSet.addAdmin("0xcfCb3cB30C940614e18A3b53Ae9fa91a8592f417");
    await addAdminTx2.wait();
	  var admins = await blockSet.getAdmins();
	  //console.log(admins);
	  expect(admins.length == 2);

    const setTx= await blockSet.setThreshold(2);
    await setTx.wait();
	  var threshold = await blockSet.getThreshold();
	  expect(threshold==2);
  });
  it("Should got value ", async function () {
    const BlockSet = await ethers.getContractFactory("blockSet");
    const blockSet = await BlockSet.deploy();
    await blockSet.deployed();
	  const addProposalTx = await blockSet.addProposal("test1",1000);
	  await addProposalTx.wait();

	  var test1 = await blockSet.getValue("test1");
	  expect(test1==1000);
  });
  it("Should revert", async function () {
    const BlockSet = await ethers.getContractFactory("blockSet");
    const blockSet = await BlockSet.deploy();
    await blockSet.deployed();
	  var initThreshold = await blockSet.getThreshold();
	  expect(initThreshold==1);

    const addAdminTx1 = await blockSet.addAdmin("0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D");
    await addAdminTx1.wait();
    const addAdminTx2 = await blockSet.addAdmin("0xcfCb3cB30C940614e18A3b53Ae9fa91a8592f417");
    await addAdminTx2.wait();
    var admins = await blockSet.getAdmins();
    expect(admins.length == 2);

    const setTx= await blockSet.setThreshold(2);
    await setTx.wait();
    var threshold = await blockSet.getThreshold();
    expect(threshold==2);

    const addProposalTx = await blockSet.addProposal("test1",1000);
    await addProposalTx.wait();
  
    await expect(blockSet.getValue("test1")).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'not valid'")
  });

    it("Should change threshold", async function () {
        const BlockSet = await ethers.getContractFactory("blockSet");
        const blockSet = await BlockSet.deploy();
        await blockSet.deployed();
        var initThreshold = await blockSet.getThreshold();
        expect(initThreshold==1);

        // add proposal
        const addProposalTx = await blockSet.addProposal("test1",1000);
        await addProposalTx.wait();

        var test1 = await blockSet.getValue("test1");
        expect(test1==1000);

        const addAdminTx1 = await blockSet.addAdmin("0x69DC6E2990C73B658DcbAB841c630F082AAAeD2D");
        await addAdminTx1.wait();
        const addAdminTx2 = await blockSet.addAdmin("0xcfCb3cB30C940614e18A3b53Ae9fa91a8592f417");
        await addAdminTx2.wait();
        var admins = await blockSet.getAdmins();
        //console.log(admins);
        expect(admins.length == 2);

        const setTx= await blockSet.setThreshold(2);
        await setTx.wait();

        var threshold = await blockSet.getThreshold();
        expect(threshold==2);

        // reset proposal.
        const resetProposalTx = await blockSet.resetProposal("test1",2000);
        await resetProposalTx.wait();
        // get proposal got old value.
        test1 = await blockSet.getValue("test1");
        expect(test1==1000);
    });
});
