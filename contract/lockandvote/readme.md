动态选举合约初始化步骤
1、部署节点合约(nodes.sol)
2、部署质押合约(lock.sol)
3、部署投票合约(vote.sol)
4、在节点合约中设置质押合约地址，设置函数：setLockContract
5、在质押合约中设置节点合约地址，设置函数：setNodeContract
6、在投票合约中设置节点合约地址，设置函数：setNodeContract

测试流程
1、添加boe节点信息，addBoeNode 或者 addBoeNodeBatch ，更新接口updateBoeNode
    查询boenodes接口
2、先质押  stake ，取消质押 withdraw

3、在投票vote 或者batchVote


