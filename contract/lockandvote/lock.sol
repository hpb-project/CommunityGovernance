pragma solidity ^0.5.1;
import "./nodes.sol";
import "./safemath.sol";

contract HpbLock is Ownable{
    using SafeMath for uint256;
    uint public lockNum=30000 ether;//默认每个账户只能锁定3万
    mapping(address=>uint) public lockBal;//保存质押金额
    mapping(address=>address payable) public nodeToLockAddr;//节点地址=》质押地址
 
    HpbNodes boenodes;
    
    function setNodeContract(address payable addr) onlyAdmin public {
        boenodes = HpbNodes(addr);
    }
    /**
	 ** 质押锁仓，动态选举可随时进行质押
	 */
    function stake(address payable nodeAddr) public payable returns (bool){
        require(msg.value>=lockNum,"value too low");//至少锁仓30000个HPB
        require(lockBal[nodeAddr]==0,"locked");//锁仓数量0代表还没锁仓
        require(boenodes.isBoeNode(nodeAddr),"not boe");
        lockBal[nodeAddr]=lockNum;//记录下锁仓金额
        nodeToLockAddr[nodeAddr]=msg.sender;//记录节点地址对应的质押地址
        boenodes.stake(nodeAddr);
        if(msg.value>=lockNum){//如果锁仓超过30000个HPB，多余的退回
            msg.sender.transfer(msg.value.sub(lockNum));
        }
        emit LockBal(nodeAddr,msg.sender,lockNum);
        return true;
    }
    /**
	 ** 解锁质押,动态选举中可随时解除质押
	 */
    function withdraw(address payable nodeAddr) public payable returns (bool){
        require(nodeToLockAddr[nodeAddr]==msg.sender,"node nq sender");//必须是节点地址对应的质押地址
        require(lockBal[nodeAddr]==lockNum);//必须是已经锁仓3万个HPB
        lockBal[nodeAddr]=0;//重置锁仓为0，代表已经全部解锁
        delete nodeToLockAddr[nodeAddr];
        boenodes.withdraw(nodeAddr);
        msg.sender.transfer(lockNum);//解锁即可取回锁仓的金额
        emit withdrawBal(nodeAddr,msg.sender,lockNum);
        return true;
    }
    
    function fetchLockInfoByNodeAddr(address nodeAddr) public view returns (
        address,
        uint
    ) {
        return (nodeToLockAddr[nodeAddr],lockBal[nodeAddr]);
    }
    
    constructor (
    ) payable public {
        owner = msg.sender;
        addAdmin(owner);
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

    function kill() onlyOwner payable public returns(bool) {
        selfdestruct(owner);
        return true;
    }
    
}
