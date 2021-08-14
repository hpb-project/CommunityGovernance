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


describe("HpbVote Testing", function () {
  let hpbNodes;
  let hpbLocks;
  let hpbVote;
  let accounts;
  let boeNumer = 6;
  let lockedNum = 5;

  beforeEach(async function () {
    const HpbNodes = await hre.ethers.getContractFactory("HpbNodes");
    hpbNodes = await HpbNodes.deploy();
    await hpbNodes.deployed();

    const HpbLocks = await hre.ethers.getContractFactory("HpbLock");
    hpbLocks = await HpbLocks.deploy();
    await hpbLocks.deployed();

    const HpbVote = await hre.ethers.getContractFactory("HpbVote");
    hpbVote = await HpbVote.deploy();
    await hpbVote.deployed();

    await hpbNodes.setLockContract(hpbLocks.address);
    await hpbLocks.setNodeContract(hpbNodes.address);
    await hpbVote.setNodeContract(hpbNodes.address);
    accounts = await hre.ethers.getSigners();

    var i;
    for (i = 0; i < boeNumer; i++ ) {
      await addOneBoe(hpbNodes, accounts[i].address); // add account to boenode
    }

    for (i = 0; i < lockedNum; i++ ) {
      await stakeSelf(hpbLocks, accounts[i]);         // lock node.
    }
  });

  // test vote under minLimit
  it("test vote under minLimit", async function () {
    let voted = ethers.utils.parseEther('1');
    let minlimit = ethers.utils.parseEther('100');
    await hpbVote.setMinLimit(minlimit);
    await expect(hpbVote.vote(accounts[1].address, voted)).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'num too low'");
  });

  // test vote not boenode
  it("test vote not boenode", async function () {
    let voted = ethers.utils.parseEther('100');
    await expect(hpbVote.vote(accounts[boeNumer].address, voted)).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'unlock node'");
  });

  // test vote not locknode
  it("test vote not locknode", async function () {
    let voted = ethers.utils.parseEther('100');
    await expect(hpbVote.vote(accounts[lockedNum].address, voted)).to.revertedWith("VM Exception while processing transaction: reverted with reason string 'unlock node'");
  });

  // test vote work right.
  it("test vote success", async function () {
    let owner = accounts[0].address;
    let candidate = accounts[1].address;
    let voted = ethers.utils.parseEther('100');
    await hpbVote.vote(candidate, voted);
    // check candidate vote number.
    expect(await hpbVote.fetchVoteNumForCandidate(candidate)).to.equal(voted);
    // check voter voted number.
    expect(await hpbVote.fetchVoteNumForVoter(owner)).to.equal(voted);
    // check voter voted to the candidate number.
    expect(await hpbVote.fetchVoteNumForVoterToCandidate(owner, candidate)).to.equal(voted);

    // fetchVoteInfoForVoter
    let [candidaters,voteNums] = await hpbVote.fetchVoteInfoForVoter(owner);
    expect(candidaters.length).equal(1);
    expect(candidaters[0]).equal(candidate);
    expect(voteNums.length).equal(1);
    expect(voteNums[0]).equal(voted);

    // fetchAllVoters
    let [allVotes,allVoted] = await hpbVote.fetchAllVoters();
    expect(allVotes.length).equal(1);
    expect(allVotes[0]).equal(owner);
    expect(allVoted.length).equal(1);
    expect(allVoted[0]).equal(voted);
  });

  // test cancel vote with unexist addr
  it("test cancel vote with unexist addr", async function () {
    let unexist = ethers.utils.getAddress("0x1111e0cc326756155d4afb2e780bdd13a17d4434");
    await expect(hpbVote.cancelVoteForCandidate(unexist)).to.be.reverted;
  });

  // test cancel vote with unvoted addr
  it("test cancel vote with unvoted addr", async function () {
    let owner = accounts[0];
    let candidate = accounts[1].address;
    let unvoted = accounts[2].address;
    let voted = ethers.utils.parseEther('100');
    await hpbVote.vote(candidate, voted);
    await hpbVote.cancelVoteForCandidate(unvoted);

    // fetchAllVoters
    let [allVotes,allVoted] = await hpbVote.fetchAllVoters();
    expect(allVotes.length).equal(1);
    expect(allVotes[0]).equal(owner.address);
    expect(allVoted.length).equal(1);
    expect(allVoted[0]).equal(voted);
  });

  // test cancel vote with exist addr
  it("test cancel vote with exist addr", async function () {
    let owner = accounts[0];
    let candidate = accounts[1].address;
    let voted = ethers.utils.parseEther('100');
    await hpbVote.vote(candidate, voted);
    await hpbVote.cancelVoteForCandidate(candidate);

    // fetchAllVoters
    let [allVoters,allVoted] = await hpbVote.fetchAllVoters();
    expect(allVoters.length).equal(0);
    expect(allVoted.length).equal(0);
  });

  // test cancel all vote
  it("test cancel all vote", async function () {
    let owner = accounts[0];
    let candidates = [accounts[1].address, accounts[2].address];
    let voted = ethers.utils.parseEther('100');
    let totalVoted = 0;
    var base = 1e18;
    // first vote.
    for (const toVote of candidates) {
      await hpbVote.vote(toVote, voted);
      totalVoted += 100;
    }
    for (const toVote of candidates) {
      // check candidate vote number.
      expect(await hpbVote.fetchVoteNumForCandidate(toVote)).to.equal(voted);
      // check voter voted to the candidate number.
      expect(await hpbVote.fetchVoteNumForVoterToCandidate(owner.address, toVote)).to.equal(voted);
    }
    // check voter voted number.
    let getVoted = await hpbVote.fetchVoteNumForVoter(owner.address);
    expect(getVoted/base).to.equal(totalVoted);
    //expect(await hpbVote.fetchVoteNumForVoter(owner.address)).to.equal();

    // fetchVoteInfoForVoter
    let [candidaters,voteNums] = await hpbVote.fetchVoteInfoForVoter(owner.address);
    expect(candidaters.length).equal(candidates.length);
    expect(candidaters[0]).equal(candidates[0]);
    expect(voteNums.length).equal(candidates.length);
    expect(voteNums[0]).equal(voted);

    // fetchAllVoters
    let [allVoters,allVoted] = await hpbVote.fetchAllVoters();
    expect(allVoters.length).equal(1);
    expect(allVoters[0]).equal(owner.address);
    expect(allVoted.length).equal(1);
    expect(allVoted[0]/base).equal(totalVoted);
    
    // do cancel all vote.
    await hpbVote.cancelVote();

    // after check.
    for (const toVote of candidates) {
      // check candidate vote number.
      expect(await hpbVote.fetchVoteNumForCandidate(toVote)).to.equal(0);
      // check voter voted to the candidate number.
      expect(await hpbVote.fetchVoteNumForVoterToCandidate(owner.address, toVote)).to.equal(0);
    }
    // check voter voted number.
    expect(await hpbVote.fetchVoteNumForVoter(owner.address)).to.equal(0);

    // fetchVoteInfoForVoter
    let [ncandidaters,nvoteNums] = await hpbVote.fetchVoteInfoForVoter(owner.address);
    expect(ncandidaters.length).equal(0);
    expect(nvoteNums.length).equal(0);

    // fetchAllVoters
    let [nallVoters,nallVoted] = await hpbVote.fetchAllVoters();
    expect(nallVoters.length).equal(0);
    expect(nallVoted.length).equal(0);
  });

  // test fetchVoteInfoForCandidate
  it("test fetchVoteInfoForCandidate", async function () {
    let voters = [accounts[0],accounts[1],accounts[2]];
    let candidates = [accounts[2].address, accounts[3].address, accounts[4].address];
    let voted = ethers.utils.parseEther('100');
    let totalVoted = [0,0,0];
    var base = 1e18;
    // first vote.
    for(var i=0; i < voters.length; i++) {
      for(const can of candidates) {
        await hpbVote.connect(voters[i]).vote(can, voted);
        totalVoted[i] += 100;  
      }
    }
    for(var i=0; i < candidates.length; i++) {
      let [allVoters,allVoted] = await hpbVote.fetchVoteInfoForCandidate(candidates[i]);
      expect(allVoters.length).equal(voters.length);
      expect(allVoted.length).equal(voters.length);
      for(const nvoted of allVoted) {
        expect(nvoted/base).equal(100);
      }
    }
  });
  // test fetchAllVoteResult
  it("test fetchAllVoteResult", async function () {
    let voters = [accounts[0],accounts[1],accounts[2]];
    let candidates = [accounts[2].address, accounts[3].address, accounts[4].address];
    let allLockedVote = new Map();
    let voted = ethers.utils.parseEther('100');
    
    var base = 1e18;

    for(var i=0; i < lockedNum; i++) {
      allLockedVote.set(accounts[i].address, 0);
    }

    // first vote.
    for(var i=0; i < voters.length; i++) {
      for(var j=0;j < candidates.length; j++) {
        let cvoted = allLockedVote.get(candidates[j]);
        cvoted += 100;
        await hpbVote.connect(voters[i]).vote(candidates[j], voted);
        allLockedVote.set(candidates[j], cvoted);
      }
    }

    let [allCandidaters,allVoted] = await hpbVote.fetchAllVoteResult();
    expect(allCandidaters.length).equal(lockedNum);
    expect(allVoted.length).equal(lockedNum);
    for(var i =0; i < allCandidaters.length; i++) {
      let cvoted = allLockedVote.get(allCandidaters[i]);
      expect(allVoted[i]/base).equal(cvoted);
    }
  });
  

  it("set and get nodeContract", async function () {
    await hpbVote.setNodeContract(hpbNodes.address);
    expect(await hpbVote.getNodeContract()).to.equal(hpbNodes.address);
  });

  it("set gasLeftLimit", async function () {
    const gasLeftLimit = 2222222222;
    await hpbVote.setGasLeftLimit(gasLeftLimit);
  });

  it("set and get vote minLimit", async function () {
    let voteMinLimit = ethers.utils.parseEther('2');
    await hpbVote.setMinLimit(voteMinLimit);
    expect(await hpbVote.minLimit()).to.equal(voteMinLimit);
  });

});