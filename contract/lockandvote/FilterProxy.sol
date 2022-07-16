pragma solidity ^0.5.1;
import "./owner.sol";
import "hardhat/console.sol";

interface IHpbNodes {
    function fetchAllHolderAddrs() external view returns(address [] memory, address [] memory);
    function getAllHpbNodes() external view returns (address payable [] memory, bytes32[] memory, bytes32[] memory, bytes32[] memory);
}

interface IHpbVote {
    function fetchAllVoteResult() external view returns (address payable[] memory, uint[] memory);
}

interface IHpbLock {
    function fetchLockInfoByNodeAddr(address nodeAddr) external view returns (address, uint);
}


contract FilterProxy is Ownable{
    mapping (address=>bool) public blacklist;
    address [] blackminer;

    uint256 public maxblackcount = 10;
    uint256 public curblackcount = 0;
    
    IHpbNodes hpbnode;
    IHpbVote  hpbvote;
    IHpbLock  hpblock;

    constructor () public {
        owner = msg.sender;
        addAdmin(owner);
    }

    function _addBlackMiner(address addr) internal {
        if (blackminer.length > curblackcount) {
            blackminer[curblackcount] = addr;
        } else {
		    blackminer.push(addr);
        }
	    curblackcount += 1;
    }

    function _rmBlackMiner(address addr) internal {
        for(uint256 i = 0; i < curblackcount; i++) {
            if(blackminer[i] == addr) {
                blackminer[i] = blackminer[curblackcount-1];
                blackminer[curblackcount-1] = address(0);
                curblackcount -= 1;
                break;
            }
        }
    }

    function setnodecontract(address payable nodeaddr) onlyAdmin public{
        hpbnode = IHpbNodes(nodeaddr);
    }

    function setvotecontract(address payable voteaddr) onlyAdmin public{
        hpbvote = IHpbVote(voteaddr);
    }

    function setlockcontract(address payable lockaddr) onlyAdmin public{
        hpblock = IHpbLock(lockaddr);
    }

    function setMaxLimit(uint256 limit) onlyAdmin public{
        require(limit >= 10, "can not limit little than 10");
        require(limit <= 50, "can not limit more than 50");
        require(limit > curblackcount, "can not limit little than curblackcount");
        maxblackcount = limit;
    }

    function addInvalidNode(address addr) onlyAdmin public {
        require(blacklist[addr] == false, "already exist in blacklist");
        require(curblackcount < maxblackcount,"exceed max black miner count");
        blacklist[addr] = true;
        _addBlackMiner(addr);
    }

    function removeInvalidNode(address addr) onlyAdmin public {
	    require(blacklist[addr] == true, "unknown black miner");
	    _rmBlackMiner(addr);
	    blacklist[addr] = false;
    }

    function getBlackList() public view returns(address []memory) {
	address [] memory retlist = new address[](curblackcount);
	for(uint256 i = 0; i < curblackcount; i++) {
		retlist[i] = blackminer[i];
        }
	return retlist;
    }

    function getcontract() public view returns(address,address,address){
        return (address(hpbnode),address(hpblock),address(hpbvote));
    }

    function getAllHpbNodes(
    )  public view returns (
        address payable[] memory,
        bytes32[] memory,
        bytes32[] memory,
        bytes32[] memory
    ) {

        uint256 count = 0;
        (address payable [] memory _coinbase, bytes32[] memory _cid1s, bytes32[] memory _cid2s, bytes32[] memory _hids) = hpbnode.getAllHpbNodes();

        for(uint256 i=0; i < _coinbase.length; i++) {
            if(blacklist[_coinbase[i]] == false) {
	    	count += 1;
            }
        }
        address payable [] memory _filter_coinbase = new address payable[](count);
        bytes32[] memory _filter_cid1s = new bytes32[](count);
        bytes32[] memory _filter_cid2s = new bytes32[](count);
        bytes32[] memory _filter_hids = new bytes32[](count);

        uint256 index = 0;
        for(uint256 i=0; i < _coinbase.length; i++) {
            if(blacklist[_coinbase[i]] == false) {
                _filter_coinbase[index] = _coinbase[i];
                _filter_cid1s[index] = _cid1s[i];
                _filter_cid2s[index] = _cid2s[i];
                _filter_hids[index] = _hids[i];
	        index += 1;
            }
        }
        
        return (_filter_coinbase, _filter_cid1s, _filter_cid2s, _filter_hids);
    }

    function fetchAllVoteResult() public view returns (
        address payable[] memory, 
        uint[] memory
    ){
        uint256 count = 0;
        (address payable [] memory _coinbase, uint[] memory _votes) = hpbvote.fetchAllVoteResult();

        for(uint256 i=0; i < _coinbase.length; i++) {
            if(blacklist[_coinbase[i]] == false) {
	    	    count += 1;
            }
        }
        address payable [] memory _filter_coinbase = new address payable[](count);
        uint[] memory _filter_votes = new uint[](count);
        
        uint256 index = 0;
        for(uint256 i=0; i < _coinbase.length; i++) {
            if(blacklist[_coinbase[i]] == false) {
                _filter_coinbase[index] = _coinbase[i];
                _filter_votes[index] = _votes[i];
	            index += 1;
            }
        }
        
        return (_filter_coinbase, _filter_votes);
    }

    function fetchAllHolderAddrs() public view returns(
        address [] memory,
        address [] memory) {
        uint256 count = 0;
        (address [] memory _coinbase, address [] memory _holder) = hpbnode.fetchAllHolderAddrs();

        for(uint256 i=0; i < _coinbase.length; i++) {
            if(blacklist[_coinbase[i]] == false) {
	    	    count += 1;
            }
        }
        address [] memory _filter_coinbase = new address [](count);
        address [] memory _filter_holder   = new address[](count);
        
        uint256 index = 0;
        for(uint256 i=0; i < _coinbase.length; i++) {
            if(blacklist[_coinbase[i]] == false) {
                _filter_coinbase[index] = _coinbase[i];
                _filter_holder[index] = _holder[i];
	            index += 1;
            }
        }
        
        return (_filter_coinbase, _filter_holder);
    }

}
