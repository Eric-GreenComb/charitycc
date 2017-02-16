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
        "message": "eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
      },
      "id": 1
    }

## REST Interface

### account  register/query
#### registerBank

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"registerBank",
    "args":["cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDenFYaVVyRFZzanhIZzdyM1RMc1NSaGlaTApqeE1nczBIU3JsRlV3c0s1eFVCcGFwSHdSazlHeGJrMWtOd2tkSzdPeHlselJjbmxzbEd5VnhhU21KYzNqbmpvClVRVFphbVRaemViMzNNZC9oYUVkN3BhSXFkSS9pZ1Z6TEtpaStXcFJkUE02VlFaV0pHam01eEMwMHhWUWw3aWcKOExBUXV6ek1OMnNCRkhsSmN3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","QPssbLyWzxZKd2ib//dZpBq0CVXGUuujnTkpMp+rsCPUqTvvwTvhrPomwaeMZagQzBzPiltKdrHZ5RiLqEGOFWsB4PJLNrAKRnVme1cHhpZ0Qwi1TP/HYgSflW5sSZ3FUM9pm75G9FuPkQpmO0bnkG1X5GKLWr3n2XRDcAlkPVg="]
    },
    "secureContext":  "jim"
    },
    "id": 3
    }

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "47fb89bf-0721-4c27-a874-3d88583c315e"
      },
      "id": 3
    }    

#### queryBank

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924"]
    },
    "secureContext":  "jim"
    },
    "id": 5
    }

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"cebbank:71266a9913375add341f1f7a4a606a2ee4e24efaf59bb2a3eb4b71aaaf5fb3fe\",\"rsaPublicKey\":\"LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDK0oxWjRiTGFPOFNobTFyOXFOOHB6RjJxbwp1c08rSnhoOWpGcDdQTkcwMERGTUR5RUhWNU9JajhhdFlVdzBwRkFTUG95dldHMGhlRzdCVU00bVRpWHNZOFF3ClpYMjZ6L2I2bk05Q3ZtM2xlell1NjgwT2NQQnNIczdLb0RZTkUrUGM2c0EwRGVDVnUxNm5aeGpQbFRwbUxMdmkKb2Z3bS9ReVlKZlVsNElRazFRSURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=\"}"
      },
      "id": 5
    }    

#### registerFund fund01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"registerFund",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEanZlcE94YkI1SGlQbVgwcnJTUThJMG90MQp2c0ljNno0RThtVnhubEtlMCt3MVpOT1J1UEdwSmE1TENNZWNqNDk4dllDcEljbk1IMkJCZjJQMk4wcnVkOW5PCllOVWZiOHlxdk5aODZNNHZBRVM4d2QxT1MzYmY1bm12Ry92Ykl5aWQwSmdIK0tRNU1LVDFZSjNGRE5TcWYrbHQKRHVTOWFUeDhtb3BKODlSME93SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","SKzK1DsdKfHMEljNPV9mtFAQfDwld2Q3cwB/yhR6wnXoZlTKpi5PPLRL+ZIYDiJgYm3rRTOmjnWA+XAjQyhKL1BukUM/l/VLBVApcLhYlM8iqhattDEOsUraEzk2RXO6hlr+bljaR4JEqW2+FdCp23eoiSdjNum6HeA/UIOedwE="]
    },
    "secureContext":  "jim"
    },
    "id": 3
    }

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "47fb89bf-0721-4c27-a874-3d88583c315e"
      },
      "id": 3
    }    

#### queryAccount fund01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"]
    },
    "secureContext":  "jim"
    },
    "id": 5
    }

result:

        

#### queryFund fund01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"queryFund",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"]
    },
    "secureContext":  "jim"
    },
    "id": 5
    }

result:

         
#### registerChannel channel01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"registerChannel",
    "args":["channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERGVNNzJVRnkvK1VTaUpTU3BkWlF2Tmt2NgpySDAvQ25raUJMRmRBUEVIUUR6bTY0Wks3bWR0bUE2ayt1alJtbENXbjhIUVRSK0xHcGp4WGxhVFc1a1Z2a2xVCks4cUdkRUUzNW5JTW0zQnc2ZldsQUlvV2I4dUxEQjlVQmpIVGxIb1dlaEE0SzV3WEh6ZmpqaWQ0U2dBNEhwbzUKdlhlazdzc3RESEpBZUhuNXdRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","jYpb/0yeO8TZ6fMPezgr4c4eRG9heObZcFVG/nRk1YSaDrW5Y2d9yaPSsjOc5p3OurFWJOQvC2xL1A0RstsPNKpIqjiickLZ/HqLXDr/4NrK4VZT2yeGwHU+nYasntaRDMwHhj1JD/zdC3sGxt8a/zQmoPmNybS7bZ0JM5vc+sQ="]
    },
    "secureContext":  "jim"
    },
    "id": 3
    }

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "47fb89bf-0721-4c27-a874-3d88583c315e"
      },
      "id": 3
    }    

#### queryAccount channel01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"]
    },
    "secureContext":  "jim"
    },
    "id": 5
    }

result:

        

#### queryChannel channel01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"queryChannel",
    "args":["channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"]
    },
    "secureContext":  "jim"
    },
    "id": 5
    }

result:

#### registerDonor donor01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"registerDonor",
    "args":["channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a","donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERS9HZ3VWbUk0QmRnUU5oYThPMmtYK3dHQgpNSkxQbGNPZllLMUdsQkpsZzlLUm9RNmZrWGd5aXRCb2FORzI5dVpnV3loVjZSMmhzUXpmRyt3cHE0VmsyQkEyClBxZEtFcVRyQVhMU0EvWkw0bUpPVEx4bjV0Qm5QL0p5OGRGaitXSjBlSWtzOGZhMUxORS9QVnR6UDlEZDV2bmEKbEZvTXNJa09maDlMcWxmaFRRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","Pj027Mhn5cpVTzsCs/NHTk8ZgrceatzxcCsU2LX4zK2eDEtupd5vxAIBNMNOScCWXu699gRdqSgkzHCHGguuidC9Y5l+Yr01fOylAxGvuuRNMTn+CitR5JZF9g0S/x+2HHc+rNljIT2Fqd1H5oHWSwEVbdTqTYqNrirgxWcl3kE="]
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
        "message": "86dbe9e4-c267-468b-a646-79cde649b074"
      },
      "id": 3
    }    

#### queryAccount donor01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }   

result:


#### queryDonor donor01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"queryDonor",
    "args":["donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }   

result:
            
#### registerSmartContract smartcontract01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"registerSmartContract",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7","smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEMG80Q3dIdE9ISzEvUXZZQzlCL0o4NTZmKwoyQTFtRGE5elpiTE5yZXhHdVlPZ1Yvb1UwTDdTaVBnSUF3Tm9HSGRENmtQakxkK2N3aEVaMHdQRnluZjhicDh5CithK00wRUJjeStkK1VQQ1MrQmNlYjdrWDVOK3VYUm9za3VzWWx3MTVRc3UyeDBrZ3hXVUJIUXZiOGxEV0xuNVUKckptSVVUdG5XWkExMHB0em13SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","eyJhZGRyIjoic21hcnRjb250cmFjdDAxOjFkNTRhODcxMzkyM2FmMTcxOGU4ZWVhYmVjM2U0ZDg1OTZkYmJkZjJkYTNmNjllYTIzYWViOGM3YTVhYjczZDgiLCJuYW1lIjoi5a6B5aSP6KW/6YOo5Zyw5Yy65q+N5Lqy5rC056qW6aG555uuIiwiZGV0YWlsIjoi5a6B5aSP6KW/6YOo5Zyw5Yy65q+N5Lqy5rC056qW6aG555uuIiwiZ29hbCI6MTAwMDAwMDAwMDAwMCwicGFydHlBIjoi5p+Q5Z+66YeR5LyaIiwicGFydHlCIjoi5p+Q5Zyw5Yy6IiwiZnVuZEFkZHIiOiJmdW5kMDE6MjVhYjU4MGEyMDkzNzc2Y2EyZTFkZDE3NzVlOTZkZmVjNWYxZmZiY2M5NTY1MTI5MzUxY2IzMzBjZjA3MTJkNyIsImZ1bmROYW1lIjoi5p+Q5Z+66YeR5LyaIiwiZnVuZE1hbmFuZ2VyRmVlIjozLCJjaGFubmVsQWRkciI6ImNoYW5uZWwwMTo5YzhiNDNjZTk0ODAxMGVmYzNiN2QxMDJhYWU1MDIxNjVjY2Q1ZTA3MTRhM2U3NjVmZTFhOGY0NDQ5MzY3ODVhIiwiY2hhbm5lbE5hbWUiOiLmn5BDaGFubmVsTmFtZSIsImNoYW5uZWxGZWUiOjIsImNyZWF0ZVRpbWVzdGFtcCI6MTQ4NzIzOTU5NiwiZm91bmRhdGlvbiI6ImZ1bmQwMToyNWFiNTgwYTIwOTM3NzZjYTJlMWRkMTc3NWU5NmRmZWM1ZjFmZmJjYzk1NjUxMjkzNTFjYjMzMGNmMDcxMmQ3In0=","GTWAUlSNgxQQDrKURYQ3L3NL3WMYuoIaDwZVdxGTKubmeDjJ0181uYjxLLq0p6yhYl2e+2P86WfJ8elKl6EU9uA73LvYurq0tyVEBQ54rE6Wm/C1WVFpJBVKpQySqpzJ2E05mWoG3zHtqOUVGjGyaZjKaZBtmJEyiuUxCZ6uj/0="]
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
        "message": "f157c2f5-9c60-491d-9d07-62315876012b"
      },
      "id": 3
    } 

#### queryAccount smartcontract01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }

result:

#### querySmartContract smartcontract01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"querySmartContract",
    "args":["smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }

result:

#### querySmartContractTrack smartcontract01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"querySmartContractTrack",
    "args":["smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }

result:


#### registerBargain bargain01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"registerBargain",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7","bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDM3R4R3hnWUtKcjlZK21iMTROcENKS0NtUgo4QmNRek9MSjNEK3EvUFo1Zk9xUzUzdFhvVko2QUZtNEwyelZLYUFkMWNOS0s4L2t3RktsV1E5YmJLZ1ZOV25zCnc4MjM0N05yRzQxaWZocFZ3dThJVHJSOGlMOC9pR3lMdnh4SGg0OWpmQ3RIWkFHV3hrWkFsVDBwdkRZYTNJMW0KTjZMKytZWCt2WW42WTdOOFB3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","eyJhZGRyIjoiYmFyZ2FpbjAxOjhmY2M1OGVhN2VkMjEyZjdjMWJhMzU5ZDE1YmVhMTQ0ZTY3YzM5MDA0NGQ5NTM3OTc1NDhjZjY3ZmQ2MjUzNGEiLCJuYW1lIjoi5a6B5aSP6KW/6YOo5Zyw5Yy65q+N5Lqy5rC056qW6aG555uuWFjljr9YWOadkeawtOeqliIsImRldGFpbCI6IuWugeWkj+ilv+mDqOWcsOWMuuavjeS6suawtOeqlumhueebrlhY5Y6/WFjmnZHmsLTnqpYiLCJzdGFydFRpbWUiOiIyMDE3LTAxLTAxIiwiZW5kVGltZSI6IjIwMTgtMDEtMDEiLCJwYXJ0eUEiOiLmn5Dln7rph5HkvJoiLCJwYXJ0eUIiOiLmn5DlnLDljLpYWOWOv1hY5p2R5pa95bel6ZifIiwiZGVwb3NpdEJhbmsiOiLlhYnlpKfpk7booYwiLCJiYW5rQWNjb3VudCI6IjEyOThoZmFrMDlra2xqYWRza2YifQ==","ODlZNQp5fjyOKiXbchNF3vlF8E/VBoUKmxmNPPTMg/giE8yadF4Fc7OdiUwsPhHzLQbGdpaltgrw+tliPwhmV0zf5zwwmra/4BUWnAtNq+Zvjx5+G8ljSDpLCKqZQvrvEX3SfR/YEhGY/9QrPasZJrm7pc+iDQ3WmEnL0b/drR0="]
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
        "message": "9162b394-accf-4d83-b580-8b5884ea4352"
      },
      "id": 3
    }     

#### queryAccount bargain01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }  

result:

### coinbase

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"coinbase",
    "args":["cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924","eyJ2ZXJzaW9uIjoxNzAxMDEsInRpbWVzdGFtcCI6MTQ4NzI0MjgxOCwidHhpbiI6W3sic291cmNlVHhIYXNoIjoibnVsbCIsImFkZHIiOiJudWxsIn1dLCJ0eG91dCI6W3sidmFsdWUiOjEwMDAwMDAwMDAwMDAwLCJhZGRyIjoiY2ViYmFuazoyOTczMWQwZTZjNmNhOWNiOTg1ZWFiZjlmZTcxNmQxNjQ0YzYyNGNhZTUyNjVjMzZjOWI3YTQ2NzAyMDAzOTI0Iiwic2lnbiI6IklXN2EzVjV1S1Zzc29WaVF1Ym15aVNUWmhDaGUvaTFETklkQkZ6K1pmNGgrbWR3bXFObmtKeFQ3Q0hCOGlxWWdsVjJKcjNicWpHNEJQY0NMNUxVaW5HUnlQQlp2cFRNMTVNSWxOenU3S0doNlBmZUdkdnhCbzFsWm5OMU1RWklOcENldWZFWlBHcTBFNXdCaEtaQzdFeTlKdjRKWTdvaVUyU2RXZ25FL2Zmcz0ifV0sIklucHV0RGF0YSI6ImRvbm9ydXVpZCIsImZvdW5kZXIiOiJjZWJiYW5rIn0=","smZPn6MCd5b07xHFgV2XTYtH6XrkBX2VCYUacNMhrswY9tOg3PpjMUStdefs7iADwgVaUenONsfw9BNaUwhCY6P5F/nE360+/au7bAZeEVtpOhu+S9E+8zoFbleX3Bslp5o6S05nsaM+BdYwRj/XSmoQ8DGfuBeZyTow7gOSh/Q="]
    },
    "secureContext":  "jim"
    },
    "id": 3
    }

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "6899d7b7-e8ea-451e-a094-f83beee4a857"
      },
      "id": 3
    }  

### changeCoin cebbank -> cebbank,donor01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
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
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"donated",
    "args":["","","",""]
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

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8\",\"balance\":99500000,\"rsaPublicKey\":\"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEMG80Q3dIdE9ISzEvUXZZQzlCL0o4NTZmKwoyQTFtRGE5elpiTE5yZXhHdVlPZ1Yvb1UwTDdTaVBnSUF3Tm9HSGRENmtQakxkK2N3aEVaMHdQRnluZjhicDh5CithK00wRUJjeStkK1VQQ1MrQmNlYjdrWDVOK3VYUm9za3VzWWx3MTVRc3UyeDBrZ3hXVUJIUXZiOGxEV0xuNVUKckptSVVUdG5XWkExMHB0em13SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=\",\"txouts\":{\"ed21c857da821e827380c8faafa99d94d4102a08ff1a7e64b05183c42f6db228:0\":{\"value\":99500000,\"addr\":\"smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8\",\"attr\":\"donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1\",\"sign\":\"Uq5MiA4JgZwaH4bfx1GwWEGVBD+LNc1IoaaVvLQeF4zNTQlj7M+QxvYst2X9m7zAaP+WuM/ufm08NMjNc4VF8fglAQJIFxxF1O6/QxLRO+04pHUlaUAIKsVvtdQ5vgLSno5a0pCNmnZMkh9KuFVPiPB/L0Cqj7Ihl8lV+jciBb0=\"}}}"
      },
      "id": 5
    }              

### drawing smartcontract01 -> smartcontract01,bargain01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
    },
    "ctorMsg": {
    "function":"drawed",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7","eyJ2ZXJzaW9uIjoxNzAxMDEsInRpbWVzdGFtcCI6MTQ4Njk1NTkxNCwidHhpbiI6W3sic291cmNlVHhIYXNoIjoiYjUwNjNmNGVjMDI0NWE4ODU5ZDU5Y2I0MDk0MzlkMzNiOTI5MGE0Y2M2ZmNmZmNjNDFiZmQwYTFjZTE2MWZkZCIsImFkZHIiOiJzbWFydGNvbnRyYWN0MDE6MWQ1NGE4NzEzOTIzYWYxNzE4ZThlZWFiZWMzZTRkODU5NmRiYmRmMmRhM2Y2OWVhMjNhZWI4YzdhNWFiNzNkOCJ9XSwidHhvdXQiOlt7InZhbHVlIjo5NTAwMDAwLCJhZGRyIjoic21hcnRjb250cmFjdDAxOjFkNTRhODcxMzkyM2FmMTcxOGU4ZWVhYmVjM2U0ZDg1OTZkYmJkZjJkYTNmNjllYTIzYWViOGM3YTVhYjczZDgiLCJhdHRyIjoiZG9ub3IwMToyNzVlNzRiMGUzNDBmNTQxMzU0OTZlNDZkODI5YjI1YWY2OTk5ODRlNjc4N2Y5YTdiMTMxOTFhZDk5MWExZWIxIiwic2lnbiI6InpFUWVnTWk3QS9CNW9DbElEUVRldS9XY0twbFF0NWlOeFRnQVI1ZDFjTWFMOFhPeElLZFRQdDFnNHFhT3hjYmtscGdRNFVNUldRWjArbU12aW1OakNqdXBCMUpqamlnYWRwVHBCTTk5bGpnVml4dy8wcTl6RXRiYSsvUnh6eUVJd1BpOGwzQnRVMEw3QXNUcDh0Z1lhV0xVQ1BpMmpRWExLRWdMdWlsV1lYZz0ifSx7InZhbHVlIjo5MDAwMDAwMCwiYWRkciI6ImJhcmdhaW4wMTo4ZmNjNThlYTdlZDIxMmY3YzFiYTM1OWQxNWJlYTE0NGU2N2MzOTAwNDRkOTUzNzk3NTQ4Y2Y2N2ZkNjI1MzRhIiwiYXR0ciI6ImRvbm9yMDE6Mjc1ZTc0YjBlMzQwZjU0MTM1NDk2ZTQ2ZDgyOWIyNWFmNjk5OTg0ZTY3ODdmOWE3YjEzMTkxYWQ5OTFhMWViMSIsInNpZ24iOiJCeVB4T1A2Zi9EQnpRc2RWTmE1UURRY2tzVW5mUmdFTTF5aklYMWhqRGR1NnJ3Z2lQamRqbk10ZTJSMUl6YmcrQkR6amRJbmVpWWxybXR3eVUrUzFIczJkcUdGalRWc3lBVSsxeUs5eGZpdmYwcFEzN3IyeFUxMk00TnVkQkZKSmVvVUs3LzdjMEl0RTB1Mmh0UUprUWlxNUFISkt2S2pDVHBmNjRWV0Fnbkk9In1dLCJJbnB1dERhdGEiOiJkb25vcnV1aWQiLCJmb3VuZGVyIjoiZnVuZDAxIn0=","dWmvPtEjTSaRIPy3m3C5wzdrpzAH9NFHlVkZHhUnnIf5aWOalAwsxtl73+UrUtfYhyWfzQXbMKbUfvqO7htfmOORsE6Jkq41uDjf0DKJBcRJN8xQ7FIzvwH0oLhPLmo24UesMPKTH2twOS+QSWbGPm2trpunDYyY0+tozs38v+M="]
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

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8\",\"balance\":9500000,\"rsaPublicKey\":\"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEMG80Q3dIdE9ISzEvUXZZQzlCL0o4NTZmKwoyQTFtRGE5elpiTE5yZXhHdVlPZ1Yvb1UwTDdTaVBnSUF3Tm9HSGRENmtQakxkK2N3aEVaMHdQRnluZjhicDh5CithK00wRUJjeStkK1VQQ1MrQmNlYjdrWDVOK3VYUm9za3VzWWx3MTVRc3UyeDBrZ3hXVUJIUXZiOGxEV0xuNVUKckptSVVUdG5XWkExMHB0em13SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=\",\"txouts\":{\"631ba3f3ff7e14ff441cb919da32b0a46d9cc7924ed5e0024660894ef2aa6074:0\":{\"value\":9500000,\"addr\":\"smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8\",\"attr\":\"donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1\",\"sign\":\"zEQegMi7A/B5oClIDQTeu/WcKplQt5iNxTgAR5d1cMaL8XOxIKdTPt1g4qaOxcbklpgQ4UMRWQZ0+mMvimNjCjupB1JjjigadpTpBM99ljgVixw/0q9zEtba+/RxzyEIwPi8l3BtU0L7AsTp8tgYaWLUCPi2jQXLKEgLuilWYXg=\"}}}"
      },
      "id": 5
    }

query bargain01 

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a\",\"balance\":90000000,\"rsaPublicKey\":\"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDM3R4R3hnWUtKcjlZK21iMTROcENKS0NtUgo4QmNRek9MSjNEK3EvUFo1Zk9xUzUzdFhvVko2QUZtNEwyelZLYUFkMWNOS0s4L2t3RktsV1E5YmJLZ1ZOV25zCnc4MjM0N05yRzQxaWZocFZ3dThJVHJSOGlMOC9pR3lMdnh4SGg0OWpmQ3RIWkFHV3hrWkFsVDBwdkRZYTNJMW0KTjZMKytZWCt2WW42WTdOOFB3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=\",\"txouts\":{\"631ba3f3ff7e14ff441cb919da32b0a46d9cc7924ed5e0024660894ef2aa6074:1\":{\"value\":90000000,\"addr\":\"bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a\",\"attr\":\"donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1\",\"sign\":\"ByPxOP6f/DBzQsdVNa5QDQcksUnfRgEM1yjIX1hjDdu6rwgiPjdjnMte2R1Izbg+BDzjdIneiYlrmtwyU+S1Hs2dqGFjTVsyAU+1yK9xfivf0pQ37r2xU12M4NudBFJJeoUK7/7c0ItE0u2htQJkQiq5AHJKvKjCTpf64VWAgnI=\"}}}"
      },
      "id": 5
    }    

### destroycoinbase bargain01 -> 0

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"eade0d093d136cbdd49f66622300dbc59405478265a8a6c8e4a05bc6e9095050036a5329ce71582dd0b69d2026612d9e1aed6cbef3472b03208f587ac658014c"
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







           
