pragma solidity ^0.5.1;
import "./safemath.sol";
import "./owner.sol";


contract NodeBallotProx is Ownable{
    using SafeMath for uint256;
    
    address payable public hpbNodesAddress;
    address payable public contractSimpleProxyAddress;
    
    address payable[] public nodeBallotAddrs;
    mapping (address => uint) public nodeBallotIndex;
    address payable public nodeBallotAddrInService;


    function _getHpbNodesInterface() internal view returns  (
        HpbNodesInterface
    ) {
        require(hpbNodesAddress != address(0));
        return HpbNodesInterface(hpbNodesAddress);
    }

    function _getContractSimpleProxyInterface() internal view returns (
        ContractSimpleProxyInterface
    ) {
        require(contractSimpleProxyAddress != address(0));
        return ContractSimpleProxyInterface(contractSimpleProxyAddress);
    }

    function getLastestBallotAddrAndIndex() public view returns(
        address payable,
        uint
    ) {
        return (
            nodeBallotAddrs[nodeBallotAddrs.length - 1],
            nodeBallotAddrs.length - 1
        );
    }
   
    function setHolderAddr(
        address payable _coinBase,
        address payable _holderAddr
    ) public returns(bool) {
        return setHolderAddrByIndex(nodeBallotAddrs.length - 1, _coinBase,_holderAddr);
    }
    function setHolderAddrByIndex(
        uint index, 
        address payable _coinBase,
        address payable _holderAddr
    )onlyAdmin public returns(bool) {
        return _getVoteInterface(index).setHolderAddr(_coinBase,_holderAddr);
    }
    
    function batchSetHolderAddr(
        address payable[] memory _coinBases,
        address payable[] memory _holderAddrs
    ) public returns(bool) {
        return batchSetHolderAddrByIndex(nodeBallotAddrs.length - 1, _coinBases,_holderAddrs);
    }

    function batchSetHolderAddrByIndex(
        uint index, 
        address payable[] memory _coinBases,
        address payable[] memory _holderAddrs
    ) public returns(bool) {
        for (uint i=0;i < _coinBases.length;i++) {
            require(setHolderAddrByIndex(index, _coinBases[i],_holderAddrs[i]));
        }
        return true;
    }

    /**
     * 投票  
     */
    function  vote(
        address payable candidateAddr, 
        uint num
    ) public returns(bool) {
        return voteByIndex(nodeBallotAddrs.length - 1, candidateAddr, num) ;
    }

    function  voteByIndex(
        uint index, 
        address payable candidateAddr, 
        uint num
    ) public returns(bool) {
        return _getVoteInterface(index).vote(msg.sender, candidateAddr, num);
    }

    /**
     * 用于批量投票  For non locked voting
     */
    function  batchVote(
        address payable[] memory candidateAddrs, 
        uint[] memory nums
    ) public returns(bool) {
        return batchVoteByIndex(nodeBallotAddrs.length - 1, candidateAddrs, nums);
    }

    function  batchVoteByIndex(
        uint index, 
        address payable[] memory candidateAddrs, 
        uint[] memory nums
    ) public returns(bool) {
        return _getVoteInterface(index).batchVote(msg.sender, candidateAddrs, nums);
    }

    function refreshVoteForAll() public returns(bool) {
        return refreshVoteForAllByIndex(nodeBallotAddrs.length - 1);
    }

    function refreshVoteForAllByIndex(
        uint index
    ) public returns(bool) {
        return _getVoteInterface(index).refreshVoteForAll();
    }

    /**
     * 撤回对某个候选人的投票 Withdraw a vote on a candidate.
     */
    function cancelVoteForCandidate(
        address payable candidateAddr, 
        uint num
    ) public returns(bool) {
        return cancelVoteForCandidateByIndex(nodeBallotAddrs.length - 1, candidateAddr, num);
    }

    function cancelVoteForCandidateByIndex(
        uint index, 
        address payable candidateAddr, 
        uint num
    ) public returns(bool) {
        return _getVoteInterface(index).cancelVoteForCandidate(msg.sender, candidateAddr, num);
    }

    function  batchCancelVoteForCandidate(
        address payable[] memory candidateAddrs, 
        uint[] memory nums
    ) public returns(bool) {
        return batchCancelVoteForCandidateByIndex(nodeBallotAddrs.length - 1, candidateAddrs, nums);
    }

    function  batchCancelVoteForCandidateByIndex(
        uint index, 
        address payable[] memory candidateAddrs, 
        uint[] memory nums
    ) public returns(bool) {
        for (uint i=0;i < candidateAddrs.length;i++) {
            require(_getVoteInterface(index).cancelVoteForCandidate(msg.sender, candidateAddrs[i], nums[i]));
        }
        return true;
    }

	function fetchAllVoterWithBalance() public view returns (
        address payable[] memory, 
        uint[] memory,
        uint[] memory
    ) {
        address payable[] memory _addrs;
        uint[] memory _voteNumbers;
        (_addrs,_voteNumbers)=fetchAllVoters();
        uint cl=_addrs.length ;
        uint[] memory _voteBalances=new uint[](cl);
        for (uint i=0;i <cl;i++) {
            _voteBalances[i] = _addrs[i].balance;
        }
        return (_addrs, _voteNumbers,_voteBalances);
    }
	function getToRefreshResult() public view returns (
	    address payable[] memory
	) {
        address payable[] memory _addrs;
        uint[] memory _voteNumbers;
        (_addrs,_voteNumbers)=fetchAllVoters();
        uint j=0;
        for (uint i=0;i <_addrs.length;i++) {
            if(_voteNumbers[i]>_addrs[i].balance){
                j++;
            }
        }
        address payable[] memory addrs=new address payable[](j);
        uint n=0;
        for (uint k=0;k <_addrs.length;k++) {
            if(_voteNumbers[k]>_addrs[k].balance){
                addrs[n]=_addrs[k];
                n++;
            }
        }
        return addrs;
    }
	function getToRefreshResultByIndex(
	    uint index
	) public view returns (
	    address payable[] memory
	) {
        address payable[] memory _addrs;
        uint[] memory _voteNumbers;
        (_addrs,_voteNumbers)=fetchAllVotersByIndex(index);
        uint j=0;
        for (uint i=0;i <_addrs.length;i++) {
            if(_voteNumbers[i]>_addrs[i].balance){
                j++;
            }
        }
        address payable[] memory addrs=new address payable[](j);
        uint n=0;
        for (uint k=0;k <_addrs.length;k++) {
            if(_voteNumbers[k]>_addrs[k].balance){
                addrs[n]=_addrs[k];
                n++;
            }
        }
        return addrs;
    }
    
    /**
     * 获取所有投票人的详细信息
     */
    function fetchAllVoters() public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        return fetchAllVotersByIndex(nodeBallotAddrs.length - 1);
    }

    function fetchAllVotersByIndex(
        uint index
    ) public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        return _getFetchVoteInterface(index).fetchAllVoters();
    }

    /**
     * 获取所有投票人的投票情况
     */
    function fetchVoteInfoForVoter(
        address payable voterAddr
    ) public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        return fetchVoteInfoForVoterByIndex(nodeBallotAddrs.length - 1, voterAddr);
    }

    function fetchVoteInfoForVoterByIndex(
        uint index, 
        address payable voterAddr
    ) public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        return _getFetchVoteInterface(index).fetchVoteInfoForVoter(voterAddr);
    }

    /**
     * 获取某个候选人的总得票数
     * Total number of votes obtained from candidates
     */
    function fetchVoteNumForCandidate(
        address payable candidateAddr
    ) public view returns (uint) {
        return fetchVoteNumForCandidateByIndex(nodeBallotAddrs.length - 1, candidateAddr);
    }

    function fetchVoteNumForCandidateByIndex(
        uint index, 
        address payable candidateAddr
    ) public view returns (uint) {
        return _getFetchVoteInterface(index).fetchVoteNumForCandidate(candidateAddr);
    }

    /**
     * 获取某个投票人已投票数
     * Total number of votes obtained from voterAddr
     */
    function fetchVoteNumForVoter(
        address payable voterAddr
    ) public view returns (uint) {
        return fetchVoteNumForVoterByIndex(nodeBallotAddrs.length - 1, voterAddr);
    }
    function fetchVoteNumForVoterWithBalance(
        address payable voterAddr
    ) public view returns (uint,uint) {
        return (fetchVoteNumForVoter(voterAddr),voterAddr.balance);
    }
    function fetchVoteNumForVoterWithBalance(
        address payable[] memory voterAddrs
    ) public view returns (
        uint[] memory,
        uint[] memory
    ) {
        uint cl=voterAddrs.length;
        uint[] memory _voteNumbers=new uint[](cl);
        uint[] memory _voteBalances=new uint[](cl);
        for(uint i=0;i<cl;i++){
            _voteNumbers[i]=fetchVoteNumForVoter(voterAddrs[i]);
            _voteBalances[i]=voterAddrs[i].balance;
        }
        return (_voteNumbers,_voteBalances);
    }

    function fetchVoteNumForVoterByIndex(
        uint index, 
        address payable voterAddr
    ) public view returns (uint) {
        return _getFetchVoteInterface(index).fetchVoteNumForVoter(voterAddr);
    }

    /**
     * 获取某个候选人被投票详细情况
     */
    function fetchVoteInfoForCandidate(
        address payable candidateAddr
    ) public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        return fetchVoteInfoForCandidateByIndex(nodeBallotAddrs.length - 1, candidateAddr);
    }

    function fetchVoteInfoForCandidateByIndex(
        uint index, 
        address payable candidateAddr
    ) public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        return _getFetchVoteInterface(index).fetchVoteInfoForCandidate(candidateAddr);
    }
	function fetchVoteNumForVoterToCandidate(
        address payable voterAddr,
        address payable candidateAddr
    ) public view returns (uint){
        return fetchVoteNumForVoterToCandidateByIndex(nodeBallotAddrs.length - 1,voterAddr,candidateAddr);
    }
	function fetchVoteNumForVoterToCandidateByIndex(
	    uint index,
        address payable voterAddr,
        address payable candidateAddr
    ) public view returns (uint){
        return _getFetchVoteInterface(index).fetchVoteNumForVoterToCandidate(voterAddr,candidateAddr);
    }

    function getHolderAddr(
        address payable _coinBase
    )  public view returns (
        address payable
    ){
        return getHolderAddrByIndex(nodeBallotAddrs.length - 1,_coinBase);
    }
    function getHolderAddrByIndex(
        uint index,
        address payable _coinBase
    )  public view returns (
        address payable
    ){
        return _getFetchVoteInterface(index).getHolderAddr(_coinBase);
    }

    function fetchAllHolderAddrsByIndex(
        uint index
    )  public view returns (
        address payable[] memory,
        address payable[] memory
    ){
        address payable[] memory coinbases=_getFetchVoteInterface(index).getAllCoinBases();
        address payable[] memory holderAddrs=new address payable[](coinbases.length);
        for (uint i =0;i <coinbases.length;i++) {
            holderAddrs[i] =getHolderAddrByIndex(index,coinbases[i]);
        }
        return (coinbases,holderAddrs);
    }

    function fetchAllHolderAddrs(
    )  public view returns (
        address[] memory,
        address[] memory
    ){
        return fetchAllHolderAddrsByIndex(nodeBallotAddrs.length - 1);
    }


    function fetchAllVoteResult() public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        return fetchAllVoteResultByIndex(nodeBallotAddrs.length - 1);
    }
    
    function fetchAllVoteResultByIndex(
        uint index
    ) public view returns (
        address payable [] memory, 
        uint[] memory
    ) {
        return _getFetchVoteInterface(index).fetchAllVoteResult();
    }
  
    function fetchResultForNodes() public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        if (FetchVoteInterface(nodeBallotAddrInService).isRunUpStage()) {
            return (new address payable[](0),new uint[](0));
        }
        address payable[] memory _candidateAddrs;
        uint[] memory _nums;
        (_candidateAddrs,_nums) = FetchVoteInterface(nodeBallotAddrInService).fetchAllVoteResult();
        if (_nums.length < 1) {
            return (new address payable[](0),new uint[](0));
        } else {
            return (_candidateAddrs,_nums);
        }
    }

	function fetchAllHolderAddrsForNodes(
    )  public view returns (
        address payable[] memory,
        address payable[] memory
    ){
        if (FetchVoteInterface(nodeBallotAddrInService).isRunUpStage()) {
            return (new address payable[](0),new address payable[](0));
        }
        address payable[] memory _coinbaseAddrs;
        address payable[] memory _holderAddrs;
        uint ballotIndex=nodeBallotIndex[nodeBallotAddrInService];
        (_coinbaseAddrs,_holderAddrs)=fetchAllHolderAddrsByIndex(ballotIndex);
        if (_coinbaseAddrs.length < 1) {
            return (new address payable[](0),new address payable[](0));
        } else {
            return (_coinbaseAddrs,_holderAddrs);
        }
    }
    function addHpbNodeBatch(
        address payable[] memory coinbases, 
        bytes32[] memory cid1s,
        bytes32[] memory cid2s,
        bytes32[] memory hids
    )onlyAdmin public returns(bool) {
        return _getHpbNodesInterface().addHpbNodeBatch(coinbases, cid1s, cid2s, hids);
    }

    function addContractMethod(
        address _contractAddr,
        bytes4 _methodId
    )onlyAdmin public returns(bool){
        return _getContractSimpleProxyInterface().addContractMethod(_contractAddr,_methodId);
    }
    function updateInvokeContract(
        uint invokeIndex,
        uint contractIndex,
        uint methodIdIndex
    )onlyAdmin public returns(bool){
        return _getContractSimpleProxyInterface().updateInvokeContract(invokeIndex,contractIndex,methodIdIndex);
    }
  
    function getContractIndexAndMethodIndex(
        address _contractAddr,
        bytes4 _methodId
    )public view returns (uint,uint){
       return _getContractSimpleProxyInterface().getContractIndexAndMethodIndex(_contractAddr,_methodId);
    }
    
    function getInvokeContract(
        uint invokeIndex
    ) public view returns (address,bytes4){
       return _getContractSimpleProxyInterface().getInvokeContract(invokeIndex);   
    }
}
}