pragma solidity ^0.5.1;
contract AdminInterface {
    function addCandidate(
        address payable _candidateAddr
    ) public returns(bool);

    function deleteCandidate(
        address payable _candidateAddr
    ) public returns(bool);

    function setCapacity(
        uint _capacity
    ) public returns(bool);

	function addCoinBase(
        address payable _coinBase
    )  public returns(bool);
    
    function initHolderAddr(
        address payable _coinBase,
        address payable _holderAddr
    ) public returns(bool);
    
    function calVoteResult() public returns(bool);
}
contract VoteInterface {
    /**
     * 投票  
     */
    function vote(
        address payable voterAddr, 
        address payable candidateAddr, 
        uint num
    ) public returns(bool);

    /**
     * 用于批量投票
     */
    function batchVote(
        address payable voterAddr, 
        address payable[] memory candidateAddrs, 
        uint[] memory nums
    ) public returns(bool);
    
    function updateCoinBase(
        address payable _coinBase,
        address payable _newCoinBase
    ) public returns(bool);
    
    function setHolderAddr(
        address payable _coinBase,
        address payable _holderAddr
    ) public returns(bool);
    
    function updateCandidateAddr(
        address payable _candidateAddr, 
        address payable _newCandidateAddr
    ) public returns(bool);
    /**
     * 撤回对某个候选人的投票
     */
    function cancelVoteForCandidate(
        address payable voterAddr, 
        address payable candidateAddr, 
        uint num
    ) public returns(bool);

    function refreshVoteForAll() public returns(bool);
    
    function refreshVoteForVoter(address payable voterAddr) public returns(bool);

}
contract FetchVoteInterface {
    /**
     * 是否为竞选阶段
     */
	function isRunUpStage() public view returns (bool);
    /**
     * 获取所有候选人的详细信息
     */
    function fetchAllCandidates() public view returns (
        address payable[] memory
    );

    /**
     * 获取所有投票人的详细信息
     */
    function fetchAllVoters() public view returns (
        address payable[] memory, 
        uint[] memory
    );

    /**
     * 获取所有投票人的投票情况
     */
    function fetchVoteInfoForVoter(
        address payable voterAddr
    ) public view returns (
        address payable[] memory, 
        uint[] memory
    );

    /**
     * 获取某个候选人的总得票数
     */
    function fetchVoteNumForCandidate(
        address payable candidateAddr
    ) public view returns (uint);

    /**
     * 获取某个投票人已投票数
     */
    function fetchVoteNumForVoter(
        address payable voterAddr
    ) public view returns (uint);

    /**
     * 获取某个候选人被投票详细情况
     */
    function fetchVoteInfoForCandidate(
        address payable candidateAddr
    ) public view returns (
        address payable[] memory, 
        uint[] memory
    );
    /**
     * 获取某个投票人对某个候选人投票数量
     */
	function fetchVoteNumForVoterToCandidate(
        address payable voterAddr,
        address payable candidateAddr
    ) public view returns (uint);
    /**
     * 获取所有候选人的得票情况
     */
    function fetchAllVoteResult() public view returns (
        address payable[] memory,
        uint[] memory
    );
    function getHolderAddr(
        address payable _coinBase
    )  public view returns (
        address payable
    );
    function getAllCoinBases(
    ) public view returns (
        address payable[] memory
    );
}
contract Ownable {
    address payable public owner;

    modifier onlyOwner {
        require(msg.sender == owner);
        // Do not forget the "_;"! It will be replaced by the actual function
        // body when the modifier is used.
        _;
    }

    function transferOwnership(
        address payable newOwner
    ) onlyOwner  public returns(bool) {
        addAdmin(newOwner);
        deleteAdmin(owner);
        owner = newOwner;
        return true;
    }

    function getOwner() public view returns (
        address payable
    ) {
        return owner;
    }
    // 合约管理员，可以添加和删除候选人
    mapping (address  => address payable) public adminMap;

    modifier onlyAdmin {
        require(adminMap[msg.sender] != address(0));
        _;
    }

    function addAdmin(address payable addr) onlyOwner public returns(bool) {
        require(adminMap[addr] == address(0));
        adminMap[addr] = addr;
        return true;
    }

    function deleteAdmin(address payable addr) onlyOwner public returns(bool) {
        require(adminMap[addr] != address(0));
        adminMap[addr] = address(0);
        return true;
    }
}
contract Monitor {
    function doVoted(
        address payable voterAddr,
        address payable candidateAddr,
        uint num,
        uint blockNumber
    )public returns(bool);
}
library SafeMath {

    /**
     * @dev Multiplies two numbers, throws on overflow.
     */
    function mul(uint256 a, uint256 b) internal pure returns (uint256 c) {
        // Gas optimization: this is cheaper than asserting 'a' not being zero,
        // but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
        if (a == 0) {
            return 0;
        }
        c = a * b;
        assert(c / a == b);
        return c;
    }

    /**
     * @dev Integer division of two numbers, truncating the quotient.
     */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        // assert(b > 0); // Solidity automatically throws when dividing by 0
        // uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't
        // hold
        return a / b;
    }

    /**
     * @dev Subtracts two numbers, throws on overflow (i.e. if subtrahend is
     *      greater than minuend).
     */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        assert(b <= a);
        return a - b;
    }

    /**
     * @dev Adds two numbers, throws on overflow.
     */
    function add(uint256 a, uint256 b) internal pure returns (uint256 c) {
        c = a + b;
        assert(c >= a);
        return c;
    }
}
contract NodeBallot is Ownable ,AdminInterface,VoteInterface,FetchVoteInterface {
    using SafeMath for uint256;

    struct Candidate{
        address payable candidateAddr;
        bool isValid;
        uint voteNumber;
    }
    Candidate[] candidateArray;//候选者的数组
    mapping (address  => uint) candidateIndexMap;//候选者地址=>候选者数组下标
 
    struct Voter{
        address payable voterAddr;
        uint voteNumber;
        mapping (uint=> uint) candidateVoteNumberMap;//候选者数组下标=>投票数 
    }
    Voter[] voterArray;//投票者的数组
    mapping (address  => uint) voterIndexMap;//投票者地址=>投票者数组下标

    address payable[] coinBaseArray;//高性能节点地址数组
    mapping (address  => uint) coinBaseIndexMap;//高性能节点地址=>高性能节点数组下标
    mapping (address  => address payable) holderMap;//高性能节点地址=>持币地址

    uint public capacity=105;//最终获选者总数（容量，获选者数量上限）默认105个
    bool isRunUps=true;//竞选准备阶段
    uint _gasLeftLimit=500000;//对于过于复杂操作，无法一步完成，那么必须分步进行
    
    bool isCalVoteResult=false;//是否真正计算竞选结果，如果正在计算过程中中断投票直到完成竞选
    
    // Mapping from owner to operator approvals
    mapping (address  => mapping (address  => bool)) private _operatorApprovals;
    
    
    //默认操作员，这里一般是官方发布的代理合约作为操作员，并且操作源头仍然必须是本人(这个在代理合约中编写的逻辑)
    address payable private _defaultOperator;
    Monitor monitor;
    address public _defaultMonitor;
    uint public minLimit=100 finney;//最小投票数量限额：0.1HPB,可外部设置
    
    event ApprovalFor(
        bool indexed approved,
        address payable indexed operator, 
        address payable indexed owner
    );
    event CandidateAdded(
        address payable indexed candidateAddr
    );
    event AddCoinBase(
        address payable indexed coinBase
    );
    event UpdateCoinBase(
        address payable indexed coinBase,
        address payable indexed newCoinBase
    );
    event SetHolderAddr(
        address payable indexed coinBase,
        address payable indexed holderAddr
    );
    event CandidateDeleted(
        address payable indexed candidateAddr
    );
    event UpdateCandidateAddr(
        address payable indexed _candidateAddr,
        address payable indexed _newCandidateAddr
    );
    event DoVoted(// 投票  flag=1为投票,flag=0为撤票
        uint indexed flag,
        address payable indexed candidateAddr,
        address payable indexed voteAddr,
        uint num
    );
    //记录发送HPB的发送者地址和发送的金额
    event ReceivedHpb(
        address payable indexed sender, 
        uint amount
    );
    //接受HPB转账
    function () payable external {
        emit ReceivedHpb(msg.sender, msg.value);
    }
    //销毁合约，并把合约余额返回给合约拥有者
    function kill() onlyOwner public returns(bool) {
        selfdestruct(owner);
        return true;
    }
    //合约拥有者提取合约余额的一部分
    function withdraw(
        uint _value
    ) onlyOwner payable public returns(bool) {
        require(address(this).balance >= _value);
        owner.transfer(_value);
        return true;
    }
    modifier onlyApproved(
        address payable addr
    ) {
        require(!isCalVoteResult);//正在计算竞选结果中，不允许操作
        require(isApproved(addr, msg.sender));//必须经过该地址允许才能操作
        _;
    }

    /**
     *用户指定允许代理其对合约操作的权限的操作员
     */
    function setApproval(
        address payable to, 
        bool approved
    ) public returns(bool) {
        require(to != msg.sender);
        _operatorApprovals[msg.sender][to] = approved;
        emit ApprovalFor(approved, to, msg.sender);
        return true;
    }

    function isApproved(
        address payable addr, 
        address payable operator
    ) public view returns (bool) {
        if (addr == operator) {//如果是自己，默认允许
            return true;
        } else if (_operatorApprovals[addr][operator]) {
            return true;
        } else {//如果是官方代理合约，并且操作源是本人也会允许
            return tx.origin == addr && operator == _defaultOperator;
        }
    }
	/**
     *设置最小允许的投票数
     */
    function setMinLimit(
        uint _minLimit
    ) onlyAdmin public returns(bool) {
        minLimit = _minLimit;
        return true;
    }
	/**
     *对投票合约设置默认的代理操作合约，便于间接的操作投票合约
     */
    function setDefaultOperator(
        address payable defaultOperator
    ) onlyAdmin public returns(bool) {
        require(_defaultOperator != msg.sender);
        _defaultOperator = defaultOperator;
        adminMap[_defaultOperator] = _defaultOperator;
        return true;
    }
    function setDefaultMonitor(
        address payable defaultMonitor
    ) onlyAdmin public returns(bool) {
        _defaultMonitor = defaultMonitor;
        monitor=Monitor(_defaultMonitor);
        return true;
    }

    /**
     * 构造函数 初始化投票智能合约的部分依赖参数
     */
    constructor () payable public {
        owner = msg.sender;
        // 设置默认管理员
        adminMap[owner] = owner;
        voterArray.push(Voter(address(0),0));
        candidateArray.push(Candidate(address(0),false,0));
        coinBaseArray.push(address(0));
    }
	/**
     * 设置最终获选者总数
     */
    function setCapacity(
        uint _capacity
    ) onlyAdmin public returns(bool) {
        capacity = _capacity;
        return true;
    }
    
    /**
     * 是否还在竞选准备阶段
     */
	function isRunUpStage() public view returns (bool) {
        return isRunUps;
    }
    
    /**
     * 为了防止因为gas消耗超出而导致本次操作失败
     */
	function setGasLeftLimit(uint gasLeftLimit) onlyAdmin public returns (bool) {
        _gasLeftLimit=gasLeftLimit;
        return true;
    }
    /**
     * 增加候选者
     */
    function addCandidate(
        address payable _candidateAddr
    ) onlyAdmin public returns(bool) {
        require(isRunUps); //必须是竞选准备阶段
        // 必须候选人地址还未使用
        require(candidateIndexMap[_candidateAddr] == 0);
        candidateIndexMap[_candidateAddr]= candidateArray.
        	push(Candidate(_candidateAddr,true,0))-1;
        
        _operatorApprovals[_candidateAddr][owner] = true;
        if(msg.sender!= owner){
        	_operatorApprovals[_candidateAddr][msg.sender] = true;
        }
        emit CandidateAdded(_candidateAddr);
        return true;
    }
    
    function addCoinBase(
        address payable _coinBase
    ) onlyAdmin public returns(bool) {
        require(!isRunUps); //必须是竞选结束阶段
        require(coinBaseIndexMap[_coinBase] == 0);
		coinBaseIndexMap[_coinBase]=coinBaseArray.push(_coinBase)-1;
        emit AddCoinBase(_coinBase);
        return true;
    }
    
    function getAllCoinBases(
    ) public view returns (
        address payable[] memory
    ) {
        if(coinBaseArray.length<2){
            return (new address payable[](0));
        }
		uint vcl=coinBaseArray.length - 1;
        address payable[] memory _coinBases=new address payable[](vcl);
        for (uint i = 1;i <= vcl;i++) {
            _coinBases[i-1] = coinBaseArray[i]; 
        }
        return (_coinBases);
    }
    
    function updateCoinBase(
        address payable _coinBase,
        address payable _newCoinBase
    ) onlyApproved(_coinBase) public returns(bool) {
        require(!isRunUps); // 必须是竞选结束阶段
        require(coinBaseIndexMap[_coinBase]!= 0);
		coinBaseArray[coinBaseIndexMap[_coinBase]]=_newCoinBase;
		coinBaseIndexMap[_newCoinBase]=coinBaseIndexMap[_coinBase];
		coinBaseIndexMap[_coinBase]=0;
        if(candidateIndexMap[_coinBase]!= 0){
            require(updateCandidateAddr(_coinBase,_newCoinBase));
        }
		emit UpdateCoinBase(_coinBase,_newCoinBase);
		return true;
    }
    
    function setHolderAddr(
        address payable _coinBase,
        address payable _holderAddr
    ) onlyApproved(_coinBase) onlyApproved(_holderAddr) public returns(bool) {
        require(!isRunUps); //必须是竞选结束阶段
        require(coinBaseIndexMap[_coinBase]!= 0);
        require(!isHolderAddrExist(_holderAddr));//持币地址不可以重复
        holderMap[_coinBase]=_holderAddr;
		emit SetHolderAddr(_coinBase,_holderAddr);
		return true;
    }
    
    function initHolderAddr(
        address payable _coinBase,
        address payable _holderAddr
    ) onlyAdmin public returns(bool) {
        require(coinBaseIndexMap[_coinBase]!= 0);
        require(holderMap[_coinBase]==address(0));//管理员只能初始化一次，如果再次修改只能是节点自己来修改
        require(!isHolderAddrExist(_holderAddr));//持币地址不可以重复
        holderMap[_coinBase]=_holderAddr;
		emit SetHolderAddr(_coinBase,_holderAddr);
		return true;
    }
    function isHolderAddrExist(
        address payable _holderAddr
    )  public view returns(bool){
        bool isExist=false;
        for(uint i=1;i<coinBaseArray.length;i++){
            if(_holderAddr==getHolderAddr(coinBaseArray[i])){
               isExist=true;
               break;
            }
        }
        return isExist;
    }
    function getHolderAddr(
        address payable _coinBase
    )  public view returns (
        address payable
    ){
        if(coinBaseIndexMap[_coinBase]== 0){
            return address(0);
        }
        if(holderMap[_coinBase]==address(0)){
            return _coinBase;
        }
        return holderMap[_coinBase];
    }
    /**
     * 删除候选者
     * @param _candidateAddr 候选者账户地址 
     */
    function deleteCandidate(
        address payable _candidateAddr
    ) onlyAdmin public returns(bool){
        require(isRunUps);//必须是竞选准备阶段
        uint candidateIndex=candidateIndexMap[_candidateAddr];
        require(candidateIndex != 0);
        candidateArray[candidateIndex].isValid=false;
        candidateArray[candidateIndex].voteNumber=0;
        //  撤销该候选者对应的投票者关联的投票数据
        for (uint n=1;n <voterArray.length;n++) {
            if(voterArray[n].voteNumber>0){
                if(voterArray[n].candidateVoteNumberMap[candidateIndex]>0){
                   	voterArray[n].voteNumber=voterArray[n].voteNumber.sub(
                   	    voterArray[n].candidateVoteNumberMap[candidateIndex]);
                    voterArray[n].candidateVoteNumberMap[candidateIndex]=0;
                }
            }
        }
        emit CandidateDeleted(_candidateAddr);
        return true;
    }

    function updateCandidateAddr(
        address payable _candidateAddr, 
        address payable _newCandidateAddr
    ) onlyApproved(_candidateAddr) public returns(bool) {
        // 判断候选人是否已经存在 Judge whether candidates exist.
        uint candidateIndex =candidateIndexMap[_candidateAddr];
        require(candidateIndex != 0);
        candidateArray[candidateIndex].candidateAddr= _newCandidateAddr;
        candidateIndexMap[_newCandidateAddr] = candidateIndex;
        candidateIndexMap[_candidateAddr] = 0;
      
        emit UpdateCandidateAddr(_candidateAddr, _newCandidateAddr);
        return true;
    }

    /**
     * 投票
     */
    function vote(
        address payable voterAddr, 
        address payable candidateAddr,
        uint num
    ) onlyApproved(voterAddr) public returns(bool){
        // 刷新所有的投票结果
        require(refreshVoteForAll());
        uint voterIndex=voterIndexMap[voterAddr];
        // 如果从没投过票，就添加投票人
        if (voterIndex == 0) {
            voterIndex = voterArray.push(Voter(voterAddr,0))-1;
            voterIndexMap[voterAddr] = voterIndex;
        }
        require(voterAddr.balance>=num.add(voterArray[voterIndex].voteNumber));
        return _doVote(voterIndex,candidateAddr, num);
    }

 	/**
     * 执行投票 do vote
     */
    function _doVote(
        uint voterIndex,
        address payable candidateAddr, 
        uint num
    ) internal returns(bool){
        // 不少于允许的最小投票数
        require(num > minLimit);
        uint candidateIndex=candidateIndexMap[candidateAddr];
        // 候选人必须存在
        require(candidateIndex!= 0);
        require(candidateArray[candidateIndex].isValid);
        voterArray[voterIndex].candidateVoteNumberMap[candidateIndex]=
            num.add(voterArray[voterIndex].candidateVoteNumberMap[candidateIndex]);
        // 投票人已投总数累加
        voterArray[voterIndex].voteNumber=num.add(voterArray[voterIndex].voteNumber);
        // 候选者得票数累加
        candidateArray[candidateIndex].voteNumber = num.add(candidateArray[candidateIndex].voteNumber);
        emit DoVoted(1,candidateAddr,voterArray[voterIndex].voterAddr,num);
        if(_defaultMonitor!=address(0)){
	    	monitor.doVoted(
	    	    voterArray[voterIndex].voterAddr,
	    	    candidateAddr,
	    	    voterArray[voterIndex].candidateVoteNumberMap[candidateIndex],
	    	    block.number
	    	);
	    }
        return true;
    }
    /**
     * 用于批量投票 
     */
    function  batchVote(
        address payable voterAddr, 
        address payable[] memory candidateAddrs,
        uint[] memory nums
    ) onlyApproved(voterAddr) public returns(bool){
        require(candidateAddrs.length==nums.length);
        // 刷新所有的投票结果
        require(refreshVoteForAll());
        uint voterIndex=voterIndexMap[voterAddr];
        // 如果从没投过票，就添加投票人
        if (voterIndex == 0) {
            voterIndex = voterArray.push(Voter(voterAddr,0))-1;
            voterIndexMap[voterAddr] = voterIndex;
        }
        for (uint i=0;i<candidateAddrs.length;i++) {
            require(_doVote(voterIndex,candidateAddrs[i], nums[i]));
        }
        require(voterAddr.balance>=voterArray[voterIndex].voteNumber);
        return true;
    }

    function refreshVoteForAll() public returns(bool) {
        if(voterArray.length>0){
	        for (uint i=1;i<voterArray.length;i++) {
	            if(voterArray[i].voteNumber>0){
	            	require(_innerRefreshVoter(i));
	            }
	            if(gasleft()<_gasLeftLimit){
	                break;
	            }
	        }
        }
        return true;
    }
    
    function refreshVoteForVoter(
        address payable voterAddr
    ) public returns(bool){
        uint voterIndex=voterIndexMap[voterAddr];
        if (voterIndex!= 0) {
            if(voterArray[voterIndex].voteNumber>0){
                require(_innerRefreshVoter(voterIndex));
            }
        }
        return true;
    }
    
	function _innerRefreshVoter(
        uint voterIndex
    ) internal returns(bool){
    	uint balance=voterArray[voterIndex].voterAddr.balance;
        if (balance <=minLimit){//如果账户余额少于最小投票数量限额，全部撤销
            for (uint k = 1;k <candidateArray.length;k++) {
                uint oldNum= voterArray[voterIndex].candidateVoteNumberMap[k];
	            if(oldNum>0){
	                require(_cancelVote(voterIndex, k, oldNum));
	            }
            }
        }else{
            uint voteNumber=voterArray[voterIndex].voteNumber;
	        if(balance<voteNumber){//如果账户余额少于已投票数
	            for (uint k=1;k<candidateArray.length;k++) {
	                uint oldNum= voterArray[voterIndex].candidateVoteNumberMap[k];
		            if(oldNum>0){
		                uint remainNum=oldNum.mul(balance).div(voteNumber);
		                if(remainNum <= minLimit) {
		                    require(_cancelVote(voterIndex, k, oldNum));
		                }else{
		                    require(_cancelVote(voterIndex, k, oldNum.sub(remainNum)));
		                }
		            }
	            }
	        }
        } 
        return true;
    }
    /**
     * 撤回对某个候选人的投票 
     */
    function cancelVoteForCandidate(
        address payable voterAddr, 
        address payable candidateAddr,
        uint num
    ) onlyApproved(voterAddr) public returns(bool) {
        uint voterIndex=voterIndexMap[voterAddr];
        require(voterIndex != 0);
        uint candidateIndex=candidateIndexMap[candidateAddr];
        require(candidateIndex != 0);
        require(candidateArray[candidateIndex].isValid);
        uint oldNum=voterArray[voterIndex].candidateVoteNumberMap[candidateIndex];
        require(oldNum >= num);
        uint remainNum=oldNum.sub(num);
        if (remainNum <=minLimit) {
            require(_cancelVote(voterIndex,candidateIndex, oldNum));
        }else{
            require(_cancelVote(voterIndex, candidateIndex, num));
        }
        return true;
    }

    function _cancelVote(
        uint voterIndex, 
        uint candidateIndex,
        uint num
    ) internal returns(bool){
        voterArray[voterIndex].candidateVoteNumberMap[candidateIndex] =
            voterArray[voterIndex].candidateVoteNumberMap[candidateIndex].sub(num);
        voterArray[voterIndex].voteNumber = voterArray[voterIndex].voteNumber.sub(num);
        candidateArray[candidateIndex].voteNumber = candidateArray[candidateIndex].voteNumber.sub(num);
        emit DoVoted(
            0,
            candidateArray[candidateIndex].candidateAddr,
            voterArray[voterIndex].voterAddr,
            num
        );
        if(_defaultMonitor!=address(0)){
	    	monitor.doVoted(
	    	    voterArray[voterIndex].voterAddr,
	    	    candidateArray[candidateIndex].candidateAddr,
	    	    voterArray[voterIndex].candidateVoteNumberMap[candidateIndex],
	    	    block.number
	        );
	    }
        return true;
    }

    /**
     * 计算选举结果
     */
    function calVoteResult() onlyAdmin public returns(bool) {
        require(isRunUps); // 必须是竞选准备阶段
        refreshVoteForAll();
        uint candidateLength=candidateArray.length;
        if (candidateLength>capacity.add(1)) {
            isCalVoteResult=true;
            address payable[] memory _inAddrs=new address payable[](capacity);
            uint[] memory _nums=new uint[](capacity);
            uint minIndex=0;
            uint upIndex=0;
            for (uint p = 1;p < candidateLength;p++) {
                if(gasleft()<_gasLeftLimit){
                    return false;
                }
                if(candidateArray[p].isValid){
	                if (upIndex<capacity) {
	                      //candidateIsValidMap[k]
	                    // 先初始化获选者数量池 Initialize the number of pools selected first.
	                    _inAddrs[upIndex] = candidateArray[p].candidateAddr;
	                    _nums[upIndex] = candidateArray[p].voteNumber;
	                    // 先记录获选者数量池中得票最少的记录 
	                    if (_nums[upIndex] < _nums[minIndex]) {
	                        minIndex = upIndex;
	                    }
	                    upIndex++;
	                } else{
	                    if (candidateArray[p].voteNumber > _nums[minIndex]) {
	                        deleteCandidate(_inAddrs[minIndex]);
	                        _inAddrs[minIndex] = candidateArray[p].candidateAddr;
	                        _nums[minIndex] = candidateArray[p].voteNumber;
	                        //重新记下最小得票者
	                        for (uint j=0;j <capacity;j++) {
	                            if (_nums[j] < _nums[minIndex]) {
	                                minIndex = j;
	                            }
	                        }
	                    } else {
	                        deleteCandidate(candidateArray[p].candidateAddr);
	                    }
	                }
                }
            }
        }
        isCalVoteResult=false;
        isRunUps=false;
        return true;
    }
    /**
     * 获取所有候选人的详细信息
     */
    function fetchAllCandidates() public view returns (
        address payable[] memory
    ) {
        uint i=0;
        for (uint j = 1;j < candidateArray.length;j++) {
            if(candidateArray[j].isValid){
                i++;
            }
        }
        address payable[] memory _addrs=new address payable[](i);
        uint p=0;
        for (uint k = 1;k < candidateArray.length;k++) {
            if(candidateArray[k].isValid){
                _addrs[p]=candidateArray[k].candidateAddr;
                p++;
            }
        }
        return _addrs;
    }

    /**
     * 获取所有投票人的详细信息
     */
    function fetchAllVoters() public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        uint i=0;
        for (uint j = 1;j < voterArray.length;j++) {
            if(voterArray[j].voteNumber>0){
                i++;
            }
        }
        address payable[] memory _addrs=new address payable[](i);
        uint[] memory _voteNumbers=new uint[](i);
        uint p=0;
        for (uint k = 1;k < voterArray.length;k++) {
            if(voterArray[k].voteNumber>0){
                _addrs[p]=voterArray[k].voterAddr;
                _voteNumbers[p] = voterArray[k].voteNumber;
                p++;
            }
        }
        return (_addrs, _voteNumbers);
        
    }
    function fetchAllVoterAddrs() public view returns (
        address payable[] memory
    ){
        uint i=0;
        for (uint j = 1;j < voterArray.length;j++) {
            if(voterArray[j].voteNumber>0){
                i++;
            }
        }
        address payable[] memory _addrs=new address payable[](i);
        uint p=0;
        for (uint k = 1;k < voterArray.length;k++) {
            if(voterArray[k].voteNumber>0){
                _addrs[p]=voterArray[k].voterAddr;
                p++;
            }
        }
        return _addrs;
    }

    /**
     * 获取投票人的投票情况
     */
    function fetchVoteInfoForVoter(
        address payable voterAddr
    ) public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        uint voterIndex = voterIndexMap[voterAddr];
        if (voterIndex == 0) { //没投过票 
            return (new address payable[](0), new uint[](0));
        }
        if (voterArray[voterIndex].voteNumber == 0) {
            return (new address payable[](0), new uint[](0));
        }
        uint i=0;
        for (uint j = 1;j < candidateArray.length;j++) {
            if(voterArray[voterIndex].candidateVoteNumberMap[j]>0){
                i++;
            }
        }
        address payable[] memory _addrs=new address payable[](i);
        uint[] memory _nums=new uint[](i);
        uint p=0;
        for (uint k = 1;k < candidateArray.length;k++) {
            if(voterArray[voterIndex].candidateVoteNumberMap[k]>0){
                _addrs[p]=candidateArray[k].candidateAddr;
                _nums[p] =voterArray[voterIndex].candidateVoteNumberMap[k];
                p++;
            }
        }
        return (_addrs, _nums);
    }

    /**
     * 获取某个候选人的总得票数 Total number of votes obtained from candidates
     */
    function fetchVoteNumForCandidate(
        address payable candidateAddr
    ) public view returns (uint) {
        uint candidateIndex=candidateIndexMap[candidateAddr];
        require(candidateIndex != 0);
        require(candidateArray[candidateIndex].isValid);
        return candidateArray[candidateIndex].voteNumber;
    }

    /**
     * 获取某个投票人已投票数 Total number of votes obtained from voterAddr
     */
    function fetchVoteNumForVoter(
        address payable voterAddr
    ) public view returns (uint) {
        uint voterIndex = voterIndexMap[voterAddr];
        if(voterIndex == 0){//没投过票
            return 0;
        }
        return voterArray[voterIndex].voteNumber;
    }
    
    function fetchVoteNumForVoterToCandidate(
        address payable voterAddr,
        address payable candidateAddr
    ) public view returns (uint) {
        uint voterIndex = voterIndexMap[voterAddr];
        if(voterIndex == 0){//没投过票
            return 0;
        }
        uint candidateIndex=candidateIndexMap[candidateAddr];
        if(candidateIndex == 0){
            return 0;
        }
        return voterArray[voterIndex].candidateVoteNumberMap[candidateIndex];
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
        uint candidateIndex=candidateIndexMap[candidateAddr];
        require(candidateIndex != 0);
        require(candidateArray[candidateIndex].isValid);
        if(candidateArray[candidateIndex].voteNumber==0){
            return (new address payable[](0),new uint[](0));
        }
        uint i=0;
        for (uint j = 1;j < voterArray.length;j++) {
            if(voterArray[j].candidateVoteNumberMap[candidateIndex]>0){
                i++;
            }
        }
        address payable[] memory _addrs=new address payable[](i);
        uint[] memory _nums=new uint[](i);
        uint p=0;
        for (uint k = 1;k < voterArray.length;k++) {
            if(voterArray[k].candidateVoteNumberMap[candidateIndex]>0){
                _addrs[p]=voterArray[k].voterAddr;
                _nums[p] =voterArray[k].candidateVoteNumberMap[candidateIndex];
                p++;
            }
        }
        return (_addrs, _nums);
    }

    function fetchAllVoteResult() public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        uint i=0;
        for (uint j = 1;j < candidateArray.length;j++) {
            if(candidateArray[j].isValid){
                i++;
            }
        }
        address payable[] memory _addrs=new address payable[](i);
        uint[] memory _nums=new uint[](i);
        uint p=0;
        for (uint k = 1;k < candidateArray.length;k++) {
            if(candidateArray[k].isValid){
                _addrs[p]=candidateArray[k].candidateAddr;
                _nums[p] =candidateArray[k].voteNumber;
                p++;
            }
        }
        return (_addrs,_nums);
    }
}