pragma solidity ^0.5.1;

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
    ) onlyOwner public returns(bool) {
        addAdmin(newOwner);
        deleteAdmin(owner);
        owner = newOwner;
        return true;
    }
    // 合约管理员,可以提前设置解锁时间
    mapping (address  => address payable) public adminMap;

    modifier onlyAdmin {
        require(adminMap[msg.sender] != address(0));
        _;
    }
    function addAdmin(
        address payable addr
    ) onlyOwner public returns(bool) {
        require(adminMap[addr] == address(0));
        adminMap[addr] = addr;
        return true;
    }
    function deleteAdmin(
        address payable addr
    ) onlyOwner public returns(bool) {
        require(adminMap[addr] != address(0));
        adminMap[addr] = address(0);
        delete adminMap[addr];
        return true;
    }
}
contract HpbNodeLockContract is Ownable{
    using SafeMath for uint256;
    uint public beginLockBlock=0;//设置为已经锁定质押区块号
    uint public lockBlockLimit=2160000;//默认经过2160000个区块后肯定解锁，保证资金一定能解锁
    uint public lockNum=30000 ether;//默认每个账户只能锁定3万
    /**
	 ** 如果状态为1表示可以质押，
	 ** 如果状态为2表示已经锁定质押，
	 ** 如果状态是3,代表到期可以解锁，即使没达到经过2592000区块后即可解锁
	 */
    uint public lockStatus=1;
    mapping(address=>uint) public lockBal;//保存质押金额
    mapping(address=>uint) public lockStartBlock;//保存质押操作区块号
    mapping(address=>address payable) public nodeToLockAddr;//节点地址=》质押地址
    mapping(address=>uint) public nodeAddrStatus;//保存节点地址的竞选状态，1为失败，非1，比如0为成功，
    mapping(uint => address payable[]) public nodeAddrs;//保存质押对应节点地址列表，根据阶段序号来保存

    uint public currentRound=0;//当前阶段序号
    /**
	 ** 设置竞选失败的节点地址
	 ** 参数为节点地址列表
	 */
    function setNodeAddrWinStatus(
        address payable[] memory _nodeAddrs
    ) onlyAdmin public returns(bool) {
        for (uint i=0;i<_nodeAddrs.length;i++) {
            if(lockBal[_nodeAddrs[i]]==lockNum){//节点已经质押
                nodeAddrStatus[_nodeAddrs[i]]=1;
            }
        }
        return true;
    }

    /**
	 ** 质押锁仓
	 ** 必须是状态为1，可以质押，
	 */
    function stake(address payable nodeAddr) public payable returns (bool){
        require(lockStatus==1);
        require(msg.value>=lockNum);//至少锁仓30000个HPB
        require(lockBal[nodeAddr]==0);//锁仓数量0代表还没锁仓
        if(nodeToLockAddr[nodeAddr]==address(0)){
            nodeAddrs[currentRound].push(nodeAddr);
        }
        if(nodeToLockAddr[nodeAddr]!=msg.sender){
            nodeToLockAddr[nodeAddr]=msg.sender;//记录节点地址对应的质押地址
        }
        lockBal[nodeAddr]=lockNum;//记录下锁仓金额
        lockStartBlock[nodeAddr]=block.number;//记录下操作锁仓区块号
        if(msg.value>=lockNum){//如果锁仓超过30000个HPB，多余的退回
            msg.sender.transfer(msg.value.sub(lockNum));
        }
        emit LockBal(nodeAddr,msg.sender,lockNum);
        return true;
    }
    /**
	 ** 解锁质押
	 ** 必须是状态为3或者经过lockBlockLimit个区块后可解锁
	 */
    function withdraw(address payable nodeAddr) public payable returns (bool){
        require(nodeToLockAddr[nodeAddr]==msg.sender);//必须是节点地址对应的质押地址
        require(lockBal[nodeAddr]==lockNum);//必须是已经锁仓3万个HPB
        require(
            lockStatus==3||//锁仓到期
            nodeAddrStatus[nodeAddr]==1||//没有竞选成功
            block.number>lockStartBlock[nodeAddr].add(lockBlockLimit)//超过指定区块时间
        );
        lockBal[nodeAddr]=0;//重置锁仓为0，代表已经全部解锁
        lockStartBlock[nodeAddr]=0;//重置锁仓操作区块号
        msg.sender.transfer(lockNum);//解锁即可取回锁仓的金额
        emit withdrawBal(nodeAddr,msg.sender,lockNum);
        return true;
    }
    function fetchLockInfosByNodeAddrs(address payable[] memory _nodeAddrs) public view returns (
        address payable[] memory ,//锁仓地址列表
        uint[] memory
    ) {
        address payable[] memory _addrs=new address payable[](_nodeAddrs.length);
        uint[] memory _nums=new uint[](_nodeAddrs.length);
        for (uint i=0;i<_nodeAddrs.length;i++) {
            _addrs[i]=nodeToLockAddr[_nodeAddrs[i]];
            _nums[i]=lockBal[_nodeAddrs[i]];
        }
        return (_addrs,_nums);
    }

    function fetchAllLockInfos() public view returns (
    	address payable[] memory ,//节点地址列表
        address payable[] memory ,//锁仓地址列表
        uint[] memory,
        uint[] memory,//质押区块
        uint[] memory//节点是否选中状态
    ) {
    	address payable[] memory _addrs=new address payable[](nodeAddrs[currentRound].length);
        uint[] memory _nums=new uint[](nodeAddrs[currentRound].length);
        uint[] memory _blocks=new uint[](nodeAddrs[currentRound].length);
        uint[] memory _status=new uint[](nodeAddrs[currentRound].length);
        for (uint i=0;i<nodeAddrs[currentRound].length;i++) {
            _addrs[i]=nodeToLockAddr[nodeAddrs[currentRound][i]];
            _nums[i]=lockBal[nodeAddrs[currentRound][i]];
            _blocks[i]=lockStartBlock[nodeAddrs[currentRound][i]];
            _status[i]=getNodeAddrStatus(nodeAddrs[currentRound][i]);
        }

        return (nodeAddrs[currentRound],_addrs,_nums,_blocks,_status);
    }
    
    function fetchAllLockInfosByRound(uint round) public view returns (
    	address payable[] memory ,//节点地址列表
        address payable[] memory ,//锁仓地址列表
        uint[] memory,
        uint[] memory,//质押区块
        uint[] memory//节点是否选中状态
    ) {
        require(round<=currentRound);
    	address payable[] memory _addrs=new address payable[](nodeAddrs[round].length);
        uint[] memory _nums=new uint[](nodeAddrs[round].length);
        uint[] memory _blocks=new uint[](nodeAddrs[round].length);
        uint[] memory _status=new uint[](nodeAddrs[round].length);
        for (uint i=0;i<nodeAddrs[round].length;i++) {
            _addrs[i]=nodeToLockAddr[nodeAddrs[round][i]];
            _nums[i]=lockBal[nodeAddrs[round][i]];
            _blocks[i]=lockStartBlock[nodeAddrs[round][i]];
            _status[i]=getNodeAddrStatus(nodeAddrs[round][i]);
        }

        return (nodeAddrs[round],_addrs,_nums,_blocks,_status);
    }
    function fetchLockInfosByLockAddr(address payable lockAddr) public view returns (
        address payable[] memory,//节点地址列表
        uint[] memory,//金额
        uint[] memory,//质押区块
        uint[] memory//节点是否选中状态
    ) {
        uint j=0;
        for (uint i=0;i<nodeAddrs[currentRound].length;i++) {
            if(nodeToLockAddr[nodeAddrs[currentRound][i]]==lockAddr){
                j++;
            }
        }
        address payable[] memory _addrs=new address payable[](j);
        uint[] memory _nums=new uint[](j);
        uint[] memory _blocks=new uint[](j);
        uint[] memory _status=new uint[](j);
        uint k=0;
        for (uint i=0;i<nodeAddrs[currentRound].length;i++) {
            if(nodeToLockAddr[nodeAddrs[currentRound][i]]==lockAddr){
                _addrs[k]=nodeAddrs[currentRound][i];
                _nums[k]=lockBal[nodeAddrs[currentRound][i]];
                _blocks[k]=lockStartBlock[nodeAddrs[currentRound][i]];
                _status[k]=getNodeAddrStatus(nodeAddrs[currentRound][i]);
                k++;
            }
        }
        return (_addrs,_nums,_blocks,_status);
    }
    function getNodeAddrStatus(address payable _nodeAddr) public view returns (uint){
        if(lockStatus==3){
            return 1;
        }else if(block.number>lockStartBlock[_nodeAddr].add(lockBlockLimit)){
            return 1;
        }else{
            return nodeAddrStatus[_nodeAddr];
        }
    }
    function fetchLockInfosByLockAddrByRound(address payable lockAddr,uint round) public view returns (
        address payable[] memory,//节点地址列表
        uint[] memory,//金额
        uint[] memory,//质押区块
        uint[] memory//节点是否选中状态
    ) {
        require(round<=currentRound);
        uint j=0;
        for (uint i=0;i<nodeAddrs[round].length;i++) {
            if(nodeToLockAddr[nodeAddrs[round][i]]==lockAddr){
                j++;
            }
        }
        address payable[] memory _addrs=new address payable[](j);
        uint[] memory _nums=new uint[](j);
        uint[] memory _blocks=new uint[](j);
        uint[] memory _status=new uint[](j);
        uint k=0;
        for (uint i=0;i<nodeAddrs[round].length;i++) {
        	address payable _nodeAddr=nodeAddrs[round][i];
            if(nodeToLockAddr[_nodeAddr]==lockAddr){
                _addrs[k]=_nodeAddr;
                _nums[k]=lockBal[_nodeAddr];
                _blocks[k]=lockStartBlock[_nodeAddr];
                _status[k]=getNodeAddrStatus(_nodeAddr);
                k++;
            }
        }
        return (_addrs,_nums,_blocks,_status);
    }
    constructor (
        uint _lockBlockLimit,
        uint _lockNum
    ) payable public {
        owner = msg.sender;
        addAdmin(owner);
        beginLockBlock=block.number;
        if(_lockBlockLimit>0){
            lockBlockLimit=_lockBlockLimit;
        }
        if(_lockNum>0){
            lockNum=_lockNum;
        }
    }
    function () payable external {
        emit ReceivedHpb(
            msg.sender,
            msg.value
        );
        if(msg.value>0){
            require(stake(msg.sender));
        }else{
            require(withdraw(msg.sender));
        }
    }
    event ReceivedHpb(
        address payable indexed sender,
        uint amount
    );
    event LockBal(
        address indexed nodeAddr,
        address indexed addr,
        uint indexed lockNum
    );
    event withdrawBal(
        address indexed nodeAddr,
        address indexed addr,
        uint indexed lockNum
    );
    /**
     ** 必须至少经过2652000(lockBlockLimit+60000=2592000+60000=2652000)
     *  个区块才能kill锁仓合约
     ** 必须是状态是3,到期解锁
     */
    function kill() onlyOwner payable public returns(bool) {
        require(lockStatus==3);
        require(beginLockBlock.add(lockBlockLimit).add(60000)<block.number);
        selfdestruct(owner);
        return true;
    }
    /**
	 ** 设置锁仓状态
	 ** 如果状态为1表示可以质押，
	 ** 如果状态为2表示已经锁定质押，
	 ** 如果状态是3,代表到期可以解锁，即使没达到经过lockBlockLimit区块后即可提前解锁
	 */
    function setlockStatus(uint _lockStatus) onlyAdmin public returns (bool) {
        lockStatus=_lockStatus;
        if(lockStatus==2){
            beginLockBlock=block.number;
        }
        return true;
    }
    function setCurrentRound(uint _currentRound) onlyAdmin public returns (bool) {
        currentRound=_currentRound;
        return true;
    }
    function addRound() onlyAdmin public returns (bool) {
        for (uint i=0;i<nodeAddrs[currentRound].length;i++) {
            if(lockBal[nodeAddrs[currentRound][i]]==lockNum){//已经锁仓，还未解锁，延续下一轮
                nodeAddrs[currentRound+1].push(nodeAddrs[currentRound][i]);
            }
        }
        currentRound=currentRound+1;
        return setlockStatus(1);
    }
}
contract HpbNodeLockContractFactory{
  event NewHpbNodeLockContract(address indexed addr, uint _lockBlockLimit,uint _lockNum);

  function createHpbNodeLockContract(uint _lockBlockLimit,uint _lockNum) public returns(address){
    HpbNodeLockContract hpbNodeLockContract = new HpbNodeLockContract(_lockBlockLimit, _lockNum);
    require(hpbNodeLockContract.transferOwnership(msg.sender));
    emit NewHpbNodeLockContract(address(hpbNodeLockContract), _lockBlockLimit, _lockNum);
    return address(hpbNodeLockContract);
  }
}
library SafeMath {
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