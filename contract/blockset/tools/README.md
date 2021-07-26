## 使用说明
### 1. 修改环境变量
修改`env.sh` 文件中变量 

    PK : 操作者的账户私钥
    url: rpc地址
    addr: 合约账户地址

### 2. 各脚本功能

#### 2.1. addAdmin.sh : 合约owner有权操作,添加管理员

**./addAdmin.sh 0xaf2f041e5c4937a1bc6a1f085e087620b59b9c21**

#### 2.2. delAdmin.sh : 合约owner有权操作,删除管理员
   
**./delAdmin.sh 0xaf2f041e5c4937a1bc6a1f085e087620b59b9c21**

#### 2.3. reset.sh    : 合约owner有权操作,重置key对应的块号
   
**./reset.sh testkey 100**

#### 2.4. proposal.sh : 提议key对应的块号
   
**./proposal.sh testkey 10**

#### 2.5. vote.sh     : 对提案的key进行投票
   
**./vote.sh testkey**

#### 2.6. get.sh      : 查询 key 对应的块号
   
**./get.sh testkey**

