## prepare

### edit yml

- add     
    volumes:
      - /home/eric/go/src/github.com/CebEcloudTime:/go/src/github.com/CebEcloudTime

### deploy chaincode

    {
      "jsonrpc": "2.0",
      "method": "deploy",
      "params": {
        "type": 1,
        "chaincodeID":{
          "path": "github.com/CebEcloudTime/charitycc"
        },
        "ctorMsg": {
          "function":"deploy",
          "args":[]
        }
      },
      "id": 1
    }


## Interface

### registerBank

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"0e16716224f7bb30361911ebc2eee510df1cfe2c757e42b18f2bd4227c4d9996176161a6da811a09450b277c65fd72324b378b2f2d4279552afde2695d8e43dc"
    },
    "ctorMsg": {
    "function":"registerBank",
    "args":["cebbank","LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDK0oxWjRiTGFPOFNobTFyOXFOOHB6RjJxbwp1c08rSnhoOWpGcDdQTkcwMERGTUR5RUhWNU9JajhhdFlVdzBwRkFTUG95dldHMGhlRzdCVU00bVRpWHNZOFF3ClpYMjZ6L2I2bk05Q3ZtM2xlell1NjgwT2NQQnNIczdLb0RZTkUrUGM2c0EwRGVDVnUxNm5aeGpQbFRwbUxMdmkKb2Z3bS9ReVlKZlVsNElRazFRSURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=","TV1xznZQSfCliJBdTvKNhn5s1M8Ji9tuET68hrMQHpbd5LO83K7yHyMJQQo7lX4iV0qVvYyK5p8zeJwVVyOeNQd1jRpi83XmfLqWHzhddf3zK81xOpzz+2GVuZolyKS2dzXyheJ7holm3PyxUwZtL5+qWfQhAXBKLqN0LrU5pcE="]
    },
    "secureContext":  "jim"
    },
    "id": 3
    }

### queryBank

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"0e16716224f7bb30361911ebc2eee510df1cfe2c757e42b18f2bd4227c4d9996176161a6da811a09450b277c65fd72324b378b2f2d4279552afde2695d8e43dc"
    },
    "ctorMsg": {
    "function":"queryBank",
    "args":["cebbank:71266a9913375add341f1f7a4a606a2ee4e24efaf59bb2a3eb4b71aaaf5fb3fe"]
    },
    "secureContext":  "jim"
    },
    "id": 5
    }

### coinbase

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"0e16716224f7bb30361911ebc2eee510df1cfe2c757e42b18f2bd4227c4d9996176161a6da811a09450b277c65fd72324b378b2f2d4279552afde2695d8e43dc"
    },
    "ctorMsg": {
    "function":"coinbase",
    "args":["cebbank:71266a9913375add341f1f7a4a606a2ee4e24efaf59bb2a3eb4b71aaaf5fb3fe","eyJ2ZXJzaW9uIjoxNzAxMDAxLCJ0aW1lc3RhbXAiOjE0ODUxNTcyOTIsInR4b3V0IjpbeyJ0eEhhc2giOiJ0eElESGFzaCIsInZhbHVlIjoxMDAwMDAwMDAwMDAwLCJhZGRyIjoiY2ViYmFuazo3MTI2NmE5OTEzMzc1YWRkMzQxZjFmN2E0YTYwNmEyZWU0ZTI0ZWZhZjU5YmIyYTNlYjRiNzFhYWFmNWZiM2ZlIiwic2lnbiI6ImIyOXVmUFlJU3MxQzdYWmMxbjhXcFdmaktUMUE0S1lJZ3JhRjdsbDBsM081NitZQThHdnk3WjdFbzZiSkUyOUZLRVZxUVBDaGY1RmVJOWg2QjFGdG9nbktvMFBVaDdMaUhaZnRsOVdlVjJCMVJMcmxwMGRlTlVvMU5TRUw3Qk03bVdLYi85ZjVQVlJ2WnVYY2lQaEdQVFpZaCtYWkdiT3JNMnJ4N0ZHYVorUT0ifV0sImZvdW5kZXIiOiJlcmljIn0=","jT9HO1vO3blGUSZ1tE1usKmzgdJhHl6Ia9moGpfYEvhSx8gMkclzU19AAv1/6k1g2BhLJw4MDqVMj8MBzm8NgHjOcD5oD0/+hOrAy+nIdOw9CZeSAnu5lOBKgb+XRqqq/WBJRFhpgwaJQAkYwNXQjtz56D7E2oedX3rtYxTPWwY="]
    },
    "secureContext":  "jim"
    },
    "id": 3
    } 
 

## Process

### 捐款流程

1. 客户端组织捐款信息(捐款人,捐款合约,金额...),签名并提交捐款

2. 链码校验签名,错误则直接返回

3. 链码代理购币,银行钱包向客户账户钱包转账等金额的公益币

4. 链码读取捐款合约,获取分配比例

5. 链码根据合约,生成捐款Tx信息  
    TX{
        TXIN: 银行转账的txout
        TXOUT:
              基金账户txout(txout的Attr设为客户的account addr)
              渠道账户txout(txout的Attr设为客户的account addr)
              合约账户txout(txout的Attr设为客户的account addr)

6. 执行UTXO模型的转账,并记录tx等信息

7. 根据TXOUT中客户的account addr,向客户的业务地址中存入三笔txout交易信息

8. 链码返回区块UUID


### 提款流程

1. 基金管理系统组织提款信息(合约账户,用途合同,金额...),签名并提交提款

2. 链码校验签名,错误则直接返回

3. 链码根据合约账户addr获取合约账户信息

4. 账户中余额不足,则直接返回

5. 循环遍历合约账户中的txout

6. 如果txout值小于等于提款,则生成提款Tx信息,并减少提款值  
    TX{
        TXIN: 合约账户的txout(txout的Attr为客户的account addr)
        TXOUT:
              合同账户txout(txout的Attr设为合约账户txout的account addr)

7. 执行UTXO模型的转账,并记录tx等信息

8. 根据TXOUT中合同账户txout的account addr,向客户的业务地址中存入txout交易信息

9. 如果txout值大于提款,则生成提款Tx信息,并减少提款值  
    TX{
        TXIN: 合约账户的txout(txout的Attr为客户的account addr)
        TXOUT:
              合同账户txout(txout的Attr设为合约账户txout的account addr)
              合约账户找零txout(txout的Attr设为合约账户txout的account addr)

10. 执行UTXO模型的转账,并记录tx等信息

11. 根据TXOUT中合同账户txout的account addr,向客户的业务地址中存入txout交易信息

12. 链码返回区块UUID







           
