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
        owner = newOwner;
        addAdmin(newOwner);
        deleteAdmin(owner);
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
contract HpbNodesInterface {

    function addStage() public returns(bool);

    function addHpbNodeBatch(
        address payable[] memory coinbases, 
        bytes32[] memory cid1s,
        bytes32[] memory cid2s,
        bytes32[] memory hids
    ) public returns(bool);

}
library Address {

  /**
   * Returns whether the target address is a contract
   * @dev This function will return false if invoked during the constructor of a contract,
   * as the code is not actually created until after the constructor finishes.
   * @param addr address to check
   * @return whether the target address is a contract
   */
    function isContract(
        address addr
    ) internal view returns (bool) {
        uint256 size;
        // XXX Currently there is no better way to check if there is a contract in an address
        // than to check the size of the code at that address.
        // See https://ethereum.stackexchange.com/a/14016/36603
        // for more details about how this works.
        // TODO Check this again before the Serenity release, because all addresses will be
        // contracts then.
        // solium-disable-next-line security/no-inline-assembly
        assembly { size := extcodesize(addr) }
        return size > 0;
    }

}
/***
 * HPB节点信息
 *  */
contract HpbNodes is Ownable,HpbNodesInterface{
    using Address for address;
    string public name = "HPB Nodes Service";
	event ReceivedHpb(
        address payable indexed sender, 
        uint amount
    );
    event ChangeStage(
        uint indexed stageNum
    );
    event AddHpbNode(
	    uint indexed stageNum,
        address indexed coinbase,
        bytes32 cid1,
        bytes32 cid2,
        bytes32 indexed hid
	);
	event UpdateHpbNode(
	    uint indexed stageNum,
        address indexed coinbase,
        bytes32 cid1,
        bytes32 cid2,
        bytes32 indexed hid
	);
	event DeleteHpbNode(
	    address indexed coinbase
	);
	// 接受HPB转账
    function () payable external {
        emit ReceivedHpb(
            msg.sender, 
            msg.value
        );
    }
    // 销毁合约，并把合约余额返回给合约拥有者
    function kill() onlyOwner payable public returns(bool) {
        selfdestruct(owner);
        return true;
    }

    function withdraw(
        uint _value
    ) onlyOwner payable public returns(bool) {
        require(address(this).balance >= _value);
        owner.transfer(_value);
        return true;
    }
    struct HpbNode{
        address payable coinbase;
        bytes32 cid1;
        bytes32 cid2;
        bytes32 hid;
    }
    
    struct NodeStage{
        uint blockNumber;
        HpbNode[] hpbNodes;
        mapping (address => uint) hpbNodesIndexMap;
    }
    
    NodeStage[] public nodeStages;
    uint public currentStageNum = 0;
    
    constructor () payable public {
        owner = msg.sender;
        addAdmin(owner);
        _changeStage(0);
    }

	function addStage(
    ) onlyAdmin public returns(bool){
       return _changeStage(currentStageNum+1);
    }
  
   function _changeStage(
        uint stageNum
   ) internal returns(bool){
        currentStageNum=stageNum;
        nodeStages.length++;
        nodeStages[currentStageNum].hpbNodesIndexMap[msg.sender]=0;
        nodeStages[currentStageNum].hpbNodes.push(HpbNode(msg.sender,0,0,0));
        nodeStages[currentStageNum].blockNumber=block.number;
        emit ChangeStage(stageNum);
        return true;
    }
    /**
	 * 添加节点信息
	 */
    function addHpbNode(
        address payable coinbase,
        bytes32 cid1,
        bytes32 cid2,
        bytes32 hid
    ) onlyAdmin public returns(bool){
        uint index = nodeStages[currentStageNum].hpbNodesIndexMap[coinbase];
        // 必须地址还未使用
        require(index == 0);
        index = nodeStages[currentStageNum].hpbNodes.length;
        nodeStages[currentStageNum].hpbNodesIndexMap[coinbase]=index;
        nodeStages[currentStageNum].hpbNodes.push(HpbNode(
            coinbase,
            cid1,
            cid2,
            hid
        ));
        emit AddHpbNode(currentStageNum,coinbase,cid1,cid2,hid);
        return true;
    }
    
    function addHpbNodeBatch(
        address payable[] memory coinbases,
        bytes32[] memory cid1s,
        bytes32[] memory cid2s,
        bytes32[] memory hids
    ) onlyAdmin public returns(bool){
        for(uint i = 0;i<coinbases.length;i++){
            require(addHpbNode(coinbases[i],cid1s[i],cid2s[i],hids[i]));
        }
        return true;
    }
	
    function updateHpbNode(
        address payable coinbase,
        bytes32 cid1,
        bytes32 cid2,
        bytes32 hid
    ) onlyAdmin public returns(bool){
        uint index = nodeStages[currentStageNum].hpbNodesIndexMap[coinbase];
        // 必须地址存在
        require(index != 0);
        nodeStages[currentStageNum].hpbNodes[index].coinbase=coinbase;
        nodeStages[currentStageNum].hpbNodes[index].cid1=cid1;
        nodeStages[currentStageNum].hpbNodes[index].cid2=cid2;
        nodeStages[currentStageNum].hpbNodes[index].hid=hid;
        emit UpdateHpbNode(currentStageNum,coinbase,cid1,cid2,hid);
        return true;
    }
    
    function updateHpbNodeBatch(
        address payable[] memory coinbases,
        bytes32[] memory cid1s,
        bytes32[] memory cid2s,
        bytes32[] memory hids
    ) onlyAdmin public returns(bool){
        for(uint i = 0;i<coinbases.length;i++){
            require(updateHpbNode(coinbases[i],cid1s[i],cid2s[i],hids[i]));
        }
        return true;
    }
    
    function deleteHpbNode(
        address payable coinbase
    ) onlyAdmin public returns(bool){
        uint index = nodeStages[currentStageNum].hpbNodesIndexMap[coinbase];
        // 必须地址存在
        require(index != 0);
        delete nodeStages[currentStageNum].hpbNodes[index];
        for (uint i = index;i<nodeStages[currentStageNum].hpbNodes.length-1;i++){
            nodeStages[currentStageNum].hpbNodesIndexMap[nodeStages[currentStageNum].hpbNodes[i+1].coinbase]=i;
            nodeStages[currentStageNum].hpbNodes[i] = nodeStages[currentStageNum].hpbNodes[i+1];
        }
        delete nodeStages[currentStageNum].hpbNodes[nodeStages[currentStageNum].hpbNodes.length-1];
        nodeStages[currentStageNum].hpbNodes.length--;
        nodeStages[currentStageNum].hpbNodesIndexMap[coinbase]=0;
        delete nodeStages[currentStageNum].hpbNodesIndexMap[coinbase];
        emit DeleteHpbNode(coinbase);
        return true;
    }
    
    function deleteHpbNodeBatch(
        address payable[] memory coinbases
    ) onlyAdmin public returns(bool){
        for(uint i = 0;i<coinbases.length;i++){
            require(deleteHpbNode(coinbases[i]));
        }
        return true;
    }
   
    function getAllHpbNodesByStageNum(
        uint _stageNum
    )  public view returns (
        address payable[] memory,
        bytes32[] memory,
        bytes32[] memory,
        bytes32[] memory
    ){
        require(_stageNum<=currentStageNum);
        NodeStage memory nodeStage=nodeStages[_stageNum];
        require(nodeStage.hpbNodes.length>1);
        uint cl=nodeStage.hpbNodes.length-1;
        address payable [] memory _coinbases=new address payable[](cl);
        bytes32[] memory _cid1s=new bytes32[](cl);
        bytes32[] memory _cid2s=new bytes32[](cl);
        bytes32[] memory _hids=new bytes32[](cl);
        for (uint i = 1;i<=cl;i++){
            _coinbases[i-1]=nodeStage.hpbNodes[i].coinbase;
            _cid1s[i-1]=nodeStage.hpbNodes[i].cid1;
            _cid2s[i-1]=nodeStage.hpbNodes[i].cid2;
            _hids[i-1]=nodeStage.hpbNodes[i].hid;
        }
        return (_coinbases,_cid1s,_cid2s,_hids);
    }
    
    function getAllHpbNodes(
    )  public view returns (
        address payable[] memory,
        bytes32[] memory,
        bytes32[] memory,
        bytes32[] memory
    ){
        return getAllHpbNodesByStageNum(currentStageNum);
    }
}