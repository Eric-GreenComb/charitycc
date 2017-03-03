## 准备fabric docker images
### pull 0.6-dp
sudo docker pull yeasy/hyperledger-fabric:0.6-dp

### tag latest
sudo docker tag yeasy/hyperledger-fabric:0.6-dp hyperledger/fabric-peer:latest

sudo docker tag yeasy/hyperledger-fabric:0.6-dp hyperledger/fabric-baseimage:latest

sudo docker tag yeasy/hyperledger-fabric:0.6-dp hyperledger/fabric-membersrvc:latest

### get compose-files
sudo git clone https://github.com/yeasy/docker-compose-files

## 准备CC代码

### 下载CC及引用代码 
go get github.com/golang/protobuf

go get github.com/CebEcloudTime/charitycc

### 编辑 4-peers-with-explorer.yml

- 在vp0中加入挂载,左侧/home/eric/go/src/...改为自己的目录
    
    
    volumes:
      - /home/eric/go/src/github.com/CebEcloudTime:/go/src/github.com/CebEcloudTime
      - /home/eric/go/src/github.com/golang:/go/src/github.com/golang

## 运行 CC

### docker-compose fabric 
sudo docker-compose -f 4-peers-with-explorer.yml up

### run explorer
浏览器访问: http://localhost:9090,确定是否4个节点都启动,如果没有都启动,重启docker-compose命令

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

result:
    
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "e6d49912bd32e9b445a91f56673cb007e90231bb53443f0e09c32c44064af763d7b6188f30f759182a1d44da25784822d4af031f277f05b041d235822bf5e35d"
      },
      "id": 1
    }

## REST Interface

### execute /CebEcloudTime/tools/fabric/deploy/deploy.sh

### query

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"551a10fd8076366e4cf2815293a21bb34e65442adf20e646ffeb83f30a925bbb3def7f593959f75cd069889496dd39b874b69a7ea3e6c983ee96a5987200426c"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924"]
    },
    "secureContext":  "jim"
    },
    "id": 5
    }

#### addr

cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924

fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7

channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a

donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1

donor02:b931578857dbe9cb37d0e59cd9cc5fb3758758d0cee1e905e25fe62f9e9e2688

donor03:a65557075090886383f6c67220e22b095dcfe1f3a6fcb47f6fb622a3ce0f1c2d

donor04:55e79613ab432a03a9c894e3efbbe5d3f0a445feea82ff49c98cb03302ea6541

donor05:f799bc349d140248bb1ea2bd1f9a820ae99161b9bc59c0ac212f6d276e9a2468

donor06:293db3e941705aaabe9505e4a513ceac8b51b8ae1407c28b7eb8ff88193eb2f6

donor07:0abc5a0def68504b23bd4fafde9fa089315f60ae6077535643ee6705643115f1

donor08:dc0e0c13816ba9270f23fb67338c21e68459dc977c85dc8ef7f4eff0e5df6189

donor09:e98462be3aebacb5bf8acf6b0d3ef46331016b4263b787a0867dea25613fb6ae

donor10:79289520e19c9d80ca228d22ca2c11dc775348cdfece34cfc32f99dd5ef89243

smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8

smartcontract02:f14e694150c33750690d2c3baad7bb3406799263f7de5920a088072da5797800

smartcontract03:9564142592ce9c66edb005f8098a725f72325d8a0411028ab78911f5d6c14718

smartcontract04:facd57477f723c07a85296c20f53ebad70c8ac4edbd67094ea3121bc8d1e8994

smartcontract05:5e8f644f5d79ad3d515f86ac545bf665595c9558e236b3c87965d4ff122a8286

bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a

bargain02:527b016abdea151f014f3fd2f840f5a02b66ebde6ffc4172d2b2e50b4c46c423

bargain03:83a9c5019541d6b8f9170bfd59e79b77a217e3cab2989d6db706b0e08240ebb3

bargain04:6a9eae00ae31d44ae7ad40cf2a66d7eee85ae624a1a737581fa7ddcf1dd20230

bargain05:0723e1bcdf3ec08dd94a83f50aab1114dc420c215fb869ae518ec755c320dab1

bargain06:990490e80e68525fc7c8cf276adec3087c261876f85934e7e553eaf0071ccb4f

bargain07:765b5cf305978f41ec5a59e854f7b3f5eaa5a3889fd6b1c387f9447a2608d921

bargain08:ea218af5f53e8e6ff824caab6fc521b5fa01fbce3525a46aa456cb63c884a522

bargain09:0fb2a0cdddddcfa33052c81b5e8ebbb05da19b69880487a49f44d2cfb37a7773

bargain10:9d681edb9d219cbb9885cc49d9bf76060febfa5e404e8c3a3fc269753250b62d


**===================================================================**
             

### changeCoin cebbank -> cebbank,donor01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"e6d49912bd32e9b445a91f56673cb007e90231bb53443f0e09c32c44064af763d7b6188f30f759182a1d44da25784822d4af031f277f05b041d235822bf5e35d"
    },
    "ctorMsg": {
    "function":"changeCoin",
    "args":["cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924","eyJ2ZXJzaW9uIjoxNzAxMDEsInRpbWVzdGFtcCI6MTQ4NzI0MzA3OSwidHhpbiI6W3sic291cmNlVHhIYXNoIjoiOTVmM2EwZThmYTE0NDBiMjIwNWIwZjM5NTI1NTFmYjAzMGIxNGI2NjU4MzIzYjI0ZDRkZjc0N2ExODJkZTYyMSIsImFkZHIiOiJjZWJiYW5rOjI5NzMxZDBlNmM2Y2E5Y2I5ODVlYWJmOWZlNzE2ZDE2NDRjNjI0Y2FlNTI2NWMzNmM5YjdhNDY3MDIwMDM5MjQifV0sInR4b3V0IjpbeyJ2YWx1ZSI6OTk5OTkwMDAwMDAwMCwiYWRkciI6ImNlYmJhbms6Mjk3MzFkMGU2YzZjYTljYjk4NWVhYmY5ZmU3MTZkMTY0NGM2MjRjYWU1MjY1YzM2YzliN2E0NjcwMjAwMzkyNCIsInNpZ24iOiJSSDlCWlJHRFlIVmcvSVZhb2FHWGhXYjBhbEQ3dEdUYVVNa05XM3l6dGprYjc3ck9TVlEvWkU3VG1yVy9sa04vanJ2N21tOXlydktmd0wyQkZEZzBhUmU3b2IvWm5NZ3FyUWNnek5TdUt4VUFvN1Ayc1U2a2FFS1hQdXhRekp5Z0hDQnRvOEFmQ1RMVENKSVlkaUdLWGExWlY1WG1zbXI1Y01zbWh2d25iWlU9In0seyJ2YWx1ZSI6MTAwMDAwMDAwLCJhZGRyIjoiZG9ub3IwMToyNzVlNzRiMGUzNDBmNTQxMzU0OTZlNDZkODI5YjI1YWY2OTk5ODRlNjc4N2Y5YTdiMTMxOTFhZDk5MWExZWIxIiwic2lnbiI6IlFERXRXSk54NjdCdVdyRnRTZFlOdndNUS9uVDFrWlFkOWRnbUdGMkN3YU9RZ202blBBeXhpR0ZwRjVKUnF1aEozcnZOa3pZTU1CSy9hR285OXFJalY2dVpLZEcxbi9Hc3dKM1M5ZFpQeXFKWEMwMGtNeHZrTk8xS2U5NFJxd2ZvWVE5M085R0VkbXRoQU5BN3F2RzVKM2s2dmZ1SXAwSU1MRHpxcnZEY3RMUT0ifV0sIklucHV0RGF0YSI6ImRvbm9ydXVpZCIsImZvdW5kZXIiOiJlcmljIn0=","Yjw95coln9u0mooy07Atva3tluAhKJwwPAa74RzG0UCy3MBrB94pPvJgv1HG8BQNifco2Q+SoKvSjO/zt8PDoAjhz4Gr35SdetFAPihmrFB9C8crSETsLdgL596yf+ZgsFOkBP+3zHHj4nO38oDfcqV3tjYUjMXU73SykkySlLk="]
    },
    "secureContext":    "jim"
    },
    "id": 3
    } 

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "8911e7be-4e43-468e-9154-871698aaf53b"
      },
      "id": 3
    }

query bank:

     

query donor01:

     
### donated donor01 -> smartcontract01,fund01,channel01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"e6d49912bd32e9b445a91f56673cb007e90231bb53443f0e09c32c44064af763d7b6188f30f759182a1d44da25784822d4af031f277f05b041d235822bf5e35d"
    },
    "ctorMsg": {
    "function":"donated",
    "args":["donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1","donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1","15F332E9906ED10294CC634747ADD787","smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8","eyJ2ZXJzaW9uIjoxNzAxMDEsInRpbWVzdGFtcCI6MTQ4NzMwMzI3MiwidHhpbiI6W3siaWR4IjoxLCJzb3VyY2VUeEhhc2giOiI3MDBlN2Y1MWRhNjczNjE5ZTYwMmI1ZDQ1ZjBjMzBiMTBmMDM0MTAxODFjNTY5ZWJkZjI1ODNmZGJlYjExMzcxIiwiYWRkciI6ImRvbm9yMDE6Mjc1ZTc0YjBlMzQwZjU0MTM1NDk2ZTQ2ZDgyOWIyNWFmNjk5OTg0ZTY3ODdmOWE3YjEzMTkxYWQ5OTFhMWViMSJ9XSwidHhvdXQiOlt7InZhbHVlIjoxMDAwMDAwMDAsImFkZHIiOiJzbWFydGNvbnRyYWN0MDE6MWQ1NGE4NzEzOTIzYWYxNzE4ZThlZWFiZWMzZTRkODU5NmRiYmRmMmRhM2Y2OWVhMjNhZWI4YzdhNWFiNzNkOCIsImF0dHIiOiJkb25vcjAxOjI3NWU3NGIwZTM0MGY1NDEzNTQ5NmU0NmQ4MjliMjVhZjY5OTk4NGU2Nzg3ZjlhN2IxMzE5MWFkOTkxYTFlYjEsMTVGMzMyRTk5MDZFRDEwMjk0Q0M2MzQ3NDdBREQ3ODcifV0sIklucHV0RGF0YSI6IjE1RjMzMkU5OTA2RUQxMDI5NENDNjM0NzQ3QURENzg3IiwiZm91bmRlciI6ImNoYW5uZWwwMSJ9","rNZPqFIxUT+jcCEmJtuMnE5oOy8eC797DDduHLihG0UlwEHrFLKXS2YB1jY5BJPmVlBnMO1pwHl5BmiWCb8Hr8ID1HFbdVi2VjPWbdEPHm7TOVmQ0XzHyjPgdmO8nxoVolh59zZHRn2OcKyvBhOW1w4EKPilxM4y6W/KurJNhdQ=",""]
    },
    "secureContext":    "jim"
    },
    "id": 3
    } 

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "d2fdcbdf-ea92-4304-a6bd-a753f64384ff"
      },
      "id": 3
    }  

query smartcontract01

### drawed smartcontract01 -> smartcontract01,bargain01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"e6d49912bd32e9b445a91f56673cb007e90231bb53443f0e09c32c44064af763d7b6188f30f759182a1d44da25784822d4af031f277f05b041d235822bf5e35d"
    },
    "ctorMsg": {
    "function":"drawed",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7","68926CF4F6EE035A2DC2E0B606D012A2","eyJ2ZXJzaW9uIjoxNzAxMDEsInRpbWVzdGFtcCI6MTQ4Nzc1OTgyMywidHhpbiI6W3sic291cmNlVHhIYXNoIjoiZDg2YWQzYmJhYzBlMjU5MTJmOGY3YTZkZTZjMDIzNDk3ZTkwNjRjNDljMGNjNTRlOTYyMjBmMWIxODA5ZjBhZiIsImFkZHIiOiJzbWFydGNvbnRyYWN0MDE6MWQ1NGE4NzEzOTIzYWYxNzE4ZThlZWFiZWMzZTRkODU5NmRiYmRmMmRhM2Y2OWVhMjNhZWI4YzdhNWFiNzNkOCJ9XSwidHhvdXQiOlt7InZhbHVlIjo5MDAwMDAwMCwiYWRkciI6ImJhcmdhaW4wMTo4ZmNjNThlYTdlZDIxMmY3YzFiYTM1OWQxNWJlYTE0NGU2N2MzOTAwNDRkOTUzNzk3NTQ4Y2Y2N2ZkNjI1MzRhIiwiYXR0ciI6ImRvbm9yMDE6Mjc1ZTc0YjBlMzQwZjU0MTM1NDk2ZTQ2ZDgyOWIyNWFmNjk5OTg0ZTY3ODdmOWE3YjEzMTkxYWQ5OTFhMWViMSwxNUYzMzJFOTkwNkVEMTAyOTRDQzYzNDc0N0FERDc4NyIsInNpZ24iOiJCeVB4T1A2Zi9EQnpRc2RWTmE1UURRY2tzVW5mUmdFTTF5aklYMWhqRGR1NnJ3Z2lQamRqbk10ZTJSMUl6YmcrQkR6amRJbmVpWWxybXR3eVUrUzFIczJkcUdGalRWc3lBVSsxeUs5eGZpdmYwcFEzN3IyeFUxMk00TnVkQkZKSmVvVUs3LzdjMEl0RTB1Mmh0UUprUWlxNUFISkt2S2pDVHBmNjRWV0Fnbkk9In1dLCJpbnB1dERhdGEiOiI2ODkyNkNGNEY2RUUwMzVBMkRDMkUwQjYwNkQwMTJBMiIsImZvdW5kZXIiOiJmdW5kMDE6MjVhYjU4MGEyMDkzNzc2Y2EyZTFkZDE3NzVlOTZkZmVjNWYxZmZiY2M5NTY1MTI5MzUxY2IzMzBjZjA3MTJkNyJ9","oAzS1tTlpmmwABfZgV5HzYJMCL+SGY9GoadP7/slxBYSH6PJuRAyhhXx9Zb2IvsYDS/tJfAnpT1fHA8IGP9tcAJ7P/BM7M7hUY6tTVOc1G1/8cQeBHq+O9tq21ttxmpLCQmOPONH3q9AeClXn+zDpRTfVlbIZQ92xglHsxt5CmQ=",""]
    },
    "secureContext":    "jim"
    },
    "id": 3
    } 

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "f13a2555-ac16-4e6d-b5b0-d38ec439966a"
      },
      "id": 3
    }  

query smartcontract01

    

query bargain01





























       


              

         

        

### destroycoinbase bargain01 -> 0

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"e6d49912bd32e9b445a91f56673cb007e90231bb53443f0e09c32c44064af763d7b6188f30f759182a1d44da25784822d4af031f277f05b041d235822bf5e35d"
    },
    "ctorMsg": {
    "function":"destroycoinbase",
    "args":["bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a","cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924","mnp0J0fQIgFsZd/+Bn/sb/V2H6s1ZzHyvf+kkfpH+VP5IsuPriucO06/JUov7RsJ7t/OcaOYCbv+3bZSveMTczisXz827GkYvudbcojjm3SXQo92HbNTNLQJZNskEVQZBp2eDUW4mZMRER7OUKY7qlzEEZUeKale/mT0xeLnjAo="]
    },
    "secureContext":    "jim"
    },
    "id": 3
    } 

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "f31b25c4-791b-4723-b4f6-24d4d1777096"
      },
      "id": 3
    }  

query bargain01

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a\",\"rsaPublicKey\":\"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDM3R4R3hnWUtKcjlZK21iMTROcENKS0NtUgo4QmNRek9MSjNEK3EvUFo1Zk9xUzUzdFhvVko2QUZtNEwyelZLYUFkMWNOS0s4L2t3RktsV1E5YmJLZ1ZOV25zCnc4MjM0N05yRzQxaWZocFZ3dThJVHJSOGlMOC9pR3lMdnh4SGg0OWpmQ3RIWkFHV3hrWkFsVDBwdkRZYTNJMW0KTjZMKytZWCt2WW42WTdOOFB3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=\"}"
      },
      "id": 5
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







           
