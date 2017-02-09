## Prepare

### go get 
go get github.com/golang/protobuf

### edit yml

- add     
    volumes:
      - /home/eric/go/src/github.com/CebEcloudTime:/go/src/github.com/CebEcloudTime
      - /home/eric/go/src/github.com/golang:/go/src/github.com/golang

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
        "message": "56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
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
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"registerBank",
    "args":["cebbank","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDenFYaVVyRFZzanhIZzdyM1RMc1NSaGlaTApqeE1nczBIU3JsRlV3c0s1eFVCcGFwSHdSazlHeGJrMWtOd2tkSzdPeHlselJjbmxzbEd5VnhhU21KYzNqbmpvClVRVFphbVRaemViMzNNZC9oYUVkN3BhSXFkSS9pZ1Z6TEtpaStXcFJkUE02VlFaV0pHam01eEMwMHhWUWw3aWcKOExBUXV6ek1OMnNCRkhsSmN3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","Bc4gxFGOGlGGZwpzz4PP+Q/Ji+DW/8oaWT1uZ4sLhA0WW4MboEsF3gsSGdvJCaswU6X+cChfBBYLw5FKHIg2kjXZXXOzD70pBLBdtdAIuObh0nKw09BzJ+r8XknTRRhIa1zLJ01lhh/UtYobPF3KrNw01WT8UUJeLnEs6I781XM="]
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
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"queryBank",
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

#### registerAccount donor01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["donor01","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERS9HZ3VWbUk0QmRnUU5oYThPMmtYK3dHQgpNSkxQbGNPZllLMUdsQkpsZzlLUm9RNmZrWGd5aXRCb2FORzI5dVpnV3loVjZSMmhzUXpmRyt3cHE0VmsyQkEyClBxZEtFcVRyQVhMU0EvWkw0bUpPVEx4bjV0Qm5QL0p5OGRGaitXSjBlSWtzOGZhMUxORS9QVnR6UDlEZDV2bmEKbEZvTXNJa09maDlMcWxmaFRRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","gG981xuGEiadVSP17Mrr6LDhTqhX2AfKjAVheiiiwlILT2XJB+ifQuf+3/yUhzZGAt/CreiFLj0jP3TN54tpurj93hmWdiPTxU5jHr45QhFWCejxyYxvz7QlZctWCUftpdOpz6NB9GkinVN6Qc1+5V8yiBlT5xnHIkqiPBA1Uxo="]
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
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
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

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1\",\"rsaPublicKey\":\"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERS9HZ3VWbUk0QmRnUU5oYThPMmtYK3dHQgpNSkxQbGNPZllLMUdsQkpsZzlLUm9RNmZrWGd5aXRCb2FORzI5dVpnV3loVjZSMmhzUXpmRyt3cHE0VmsyQkEyClBxZEtFcVRyQVhMU0EvWkw0bUpPVEx4bjV0Qm5QL0p5OGRGaitXSjBlSWtzOGZhMUxORS9QVnR6UDlEZDV2bmEKbEZvTXNJa09maDlMcWxmaFRRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=\"}"
      },
      "id": 5
    }

#### registerAccount channel01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["channel01","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERGVNNzJVRnkvK1VTaUpTU3BkWlF2Tmt2NgpySDAvQ25raUJMRmRBUEVIUUR6bTY0Wks3bWR0bUE2ayt1alJtbENXbjhIUVRSK0xHcGp4WGxhVFc1a1Z2a2xVCks4cUdkRUUzNW5JTW0zQnc2ZldsQUlvV2I4dUxEQjlVQmpIVGxIb1dlaEE0SzV3WEh6ZmpqaWQ0U2dBNEhwbzUKdlhlazdzc3RESEpBZUhuNXdRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","AbHdSWk/vmQVgBIihWJpLWLkAPxsd/4gHzyQlWyY3AHM4FZiiVimi+wDDayCSEqp0PzSF4iMa8zKI2AEWc4A17NimHkBquD08BagxDeuyKplMGjDEAXzvxNnskZUFim0rEBEFiNJd5YbY1gTWn2/xwGZTRbuAmQ34svNGyTL0TU="]
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
        "message": "2565feb1-67c3-4c97-805e-52288dbc41b7"
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
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }     

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"channel01:cc68d3ccb7be906dbf9d63e4ae3e9605c8ddd63fb710603e3cb45976932db9f4\",\"rsaPublicKey\":\"LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDbFBibUthUk5FVm05NHUwSmZ1anVwd1ZVSgo5MVc3YmQyODc4NHdDbS9HUi85Wm1jRzZkRlNRUC9yaTQ5RHExR0hLekJEd0lvRzVXOGJmZndLeERiOTIzOWsrCmxwaHV1VjloZ0I3TlliQ25VUUo0WmkvUUlQZDFqR2pSem1DYlcvajJOdDJid1JJL3loNHNMbVhMb2Qzb2FuczcKcjR5ZzhnN2k4blRZUHlLWllRSURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=\"}"
      },
      "id": 5
    }      

#### registerAccount fund01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["fund01","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEanZlcE94YkI1SGlQbVgwcnJTUThJMG90MQp2c0ljNno0RThtVnhubEtlMCt3MVpOT1J1UEdwSmE1TENNZWNqNDk4dllDcEljbk1IMkJCZjJQMk4wcnVkOW5PCllOVWZiOHlxdk5aODZNNHZBRVM4d2QxT1MzYmY1bm12Ry92Ykl5aWQwSmdIK0tRNU1LVDFZSjNGRE5TcWYrbHQKRHVTOWFUeDhtb3BKODlSME93SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","ljsig5W3p7GHEYg1VWjxd7QMgXefOWi3CbqcbI41RhErkYBdbRRlUCX6HImVrRXi3pQFCcQ8Wo0rm0uAHuhDHFqV3CgmNVcc5kzPTC+WiQgMzsbeOrmH/S3htnkBcHIrpWkUAS9w5wU6ANCSATyYJenH7vSMcT2N9MBzyQZ1J8M="]
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
        "message": "4ca7daaa-a753-451d-b473-99443bc20333"
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
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }   

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"fund01:da336a2a2d81cdf570b5d3d70926144997e5e277f49897c6885acacba8b1cc13\",\"rsaPublicKey\":\"LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDZ0dDTEdVN1hmbFQ2UjNQdW1NN3dPRjJmdQowbUhmREl6dlA3SkUxUEJvdXR3NzkvT0F6NWJnWXIrQUtwM2VhSFhlemlIb3JTQTZFUFNDUGNoTm01ZzFsOCt4CjQ0ODJ0dm4zWVA3YWdxTmJISHI2MjNwOEprOXEwWFVhS3U0OXVEMDNjLy9EcnM2bjVjWE9QdVF4VVFkQ3VnYTkKT1FVR0RDTEs4RElUVzdXU1J3SURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=\"}"
      },
      "id": 5
    }        

#### registerAccount smartcontract01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["smartcontract01","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEMG80Q3dIdE9ISzEvUXZZQzlCL0o4NTZmKwoyQTFtRGE5elpiTE5yZXhHdVlPZ1Yvb1UwTDdTaVBnSUF3Tm9HSGRENmtQakxkK2N3aEVaMHdQRnluZjhicDh5CithK00wRUJjeStkK1VQQ1MrQmNlYjdrWDVOK3VYUm9za3VzWWx3MTVRc3UyeDBrZ3hXVUJIUXZiOGxEV0xuNVUKckptSVVUdG5XWkExMHB0em13SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","jy1NvbURF+UaieDQzbc7xojJJApvQlFL9XnASM6gkDO4KcoAN4pk+mnEGtbmx5GPXVX9mPR+8pSJHcsjGPZPYPAP68lYOcoZCN3oYmDnTuJA0Vh1OU+DN2oQHknuzujz8t28imeRBkHAgbeIRv5Gke1dRRjtGAbE1DbBkl+/on0="]
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
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
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

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"treaty01:e4811d624ebbacbaadc59bf0192c4229a9a4e05e1fa269dde345bf2bfb3f9bdd\",\"rsaPublicKey\":\"LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDeHV2RXZ5OTd4OUpiR3ZVMzlBMTZIZ3FTNgo1WDhGaGFMNzh5Wm1QSE9DZnNjY3paZk4ydE5uaHlqR3dvb1Zwbmd5bFJ5d3Bib29BekFoeVNLcVZROGdpNjBHCnRJa01QWEZ0RHZyVFJOUFZhd2MwdTloVVY5SnhSdkx1c0JjSzBmSGtkUGhuamlTa25nNGk3Mk51OUJFUlFveGMKNTQzOVU2aEpLZE45SUdta2lRSURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=\"}"
      },
      "id": 5
    }      

#### registerAccount bargain01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["bargain01","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDM3R4R3hnWUtKcjlZK21iMTROcENKS0NtUgo4QmNRek9MSjNEK3EvUFo1Zk9xUzUzdFhvVko2QUZtNEwyelZLYUFkMWNOS0s4L2t3RktsV1E5YmJLZ1ZOV25zCnc4MjM0N05yRzQxaWZocFZ3dThJVHJSOGlMOC9pR3lMdnh4SGg0OWpmQ3RIWkFHV3hrWkFsVDBwdkRZYTNJMW0KTjZMKytZWCt2WW42WTdOOFB3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=","dIP428HgM+sGIpUmi9CXwgOQ0cmoMDc3E46GXN4tyo3qo/eJLISL3Bl2qTTi0KVXjs0S522TygQs3y3/NAyiwmAiL5ZeT/81ckeQf9ErUE+LreykkA5kD0ZFpEmtuChFfkz1DXYRr6o/IXzrfl4MZhJtiBwSzU6Vc2bhpOwvdko="]
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
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
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

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"contract01:8017d1bd15326b48b7020ff6fe434ae4999cf050025ec64871516677f1177460\",\"rsaPublicKey\":\"LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDcDlxUTBINVY4WUkwOWFGMTJvVnhjWTlVUQpuUzNVRFZwT0c2OXo3a0lzK3o3WUJXR0hlVWtKSjk2U2JrVVU1WWtoU0JhbXZoTXBoUW9STS90MlI5YTJpem5FCkNyWmladVpHZmxJdEFtU0lxOXB5Ris5M1UwYVdQWTcxc1VqSCtnQ0orcFBPSFJwSWFoRWtHZ2t1QTN1NU5wZncKUTBJQnJGOHdNWi9NSnhNSFl3SURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=\"}"
      },
      "id": 5
    }           

### smartcontract register/query
#### registerSmartContract smartcontract01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"registerSmartContract",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7","eyJhZGRyIjoic21hcnRjb250cmFjdDAxOjFkNTRhODcxMzkyM2FmMTcxOGU4ZWVhYmVjM2U0ZDg1OTZkYmJkZjJkYTNmNjllYTIzYWViOGM3YTVhYjczZDgiLCJuYW1lIjoi5a6B5aSP6KW/6YOo5Zyw5Yy65q+N5Lqy5rC056qW6aG555uuIiwiZGV0YWlsIjoi5a6B5aSP6KW/6YOo5Zyw5Yy65q+N5Lqy5rC056qW6aG555uuIiwiZ29hbCI6MTAwMDAwMDAwMDAwMCwicGFydHlBIjoi5p+Q5Z+66YeR5LyaIiwicGFydHlCIjoi5p+Q5Zyw5Yy6IiwiZnVuZE1hbmFuZ2VyRmVlIjozLCJjaGFubmVsRmVlIjoyLCJjcmVhdGVUaW1lc3RhbXAiOjE0ODY2Mzc2NDQsImZvdW5kYXRpb24iOiJmdW5kMDE6MjVhYjU4MGEyMDkzNzc2Y2EyZTFkZDE3NzVlOTZkZmVjNWYxZmZiY2M5NTY1MTI5MzUxY2IzMzBjZjA3MTJkNyJ9","eKQd9bW7N/h2miHVLOZZI+d4nkxHi/EtkNPjOwEKU9TP9BZV4ZoviXG3iZjrStMpmc/QOTsq2Rmmbacy9nFPrPXoDUXkae2SypAejF8VfrsOG9i6DgTF6SIj4P8G4TV6LitQIAd67CLOR8N5HdEezhGVmIrWL9gDoK/yDBAglE0="]
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
        "message": "c5ed2f60-80f7-4e53-b2a1-26781d6ce6f8"
      },
      "id": 3
    }      

#### querySmartContract smartcontract01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
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

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8\",\"name\":\"宁夏西部地区母亲水窖项目\",\"detail\":\"宁夏西部地区母亲水窖项目\",\"goal\":1000000000000,\"partyA\":\"某基金会\",\"partyB\":\"某地区\",\"fundManangerFee\":3,\"channelFee\":2,\"createTimestamp\":1486637644,\"foundation\":\"fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7\"}"
      },
      "id": 5
    }      

#### querySmartContracts smartcontract01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"querySmartContracts",
    "args":["smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "[{\"addr\":\"smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8\",\"name\":\"宁夏西部地区母亲水窖项目\",\"detail\":\"宁夏西部地区母亲水窖项目\",\"goal\":1000000000000,\"partyA\":\"某基金会\",\"partyB\":\"某地区\",\"fundManangerFee\":3,\"channelFee\":2,\"createTimestamp\":1486637644,\"foundation\":\"fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7\"}]"
      },
      "id": 5
    }        

### bargain register/query
#### registerBargain bargain01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"registerBargain",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7","eyJhZGRyIjoiYmFyZ2FpbjAxOjhmY2M1OGVhN2VkMjEyZjdjMWJhMzU5ZDE1YmVhMTQ0ZTY3YzM5MDA0NGQ5NTM3OTc1NDhjZjY3ZmQ2MjUzNGEiLCJuYW1lIjoi5a6B5aSP6KW/6YOo5Zyw5Yy65q+N5Lqy5rC056qW6aG555uuWFjljr9YWOadkeawtOeqliIsImRldGFpbCI6IuWugeWkj+ilv+mDqOWcsOWMuuavjeS6suawtOeqlumhueebrlhY5Y6/WFjmnZHmsLTnqpYiLCJzdGFydFRpbWUiOiIyMDE3LTAxLTAxIiwiZW5kVGltZSI6IjIwMTgtMDEtMDEiLCJwYXJ0eUEiOiLmn5Dln7rph5HkvJoiLCJwYXJ0eUIiOiLmn5DlnLDljLpYWOWOv1hY5p2R5pa95bel6ZifIiwiZGVwb3NpdEJhbmsiOiLlhYnlpKfpk7booYwiLCJiYW5rQWNjb3VudCI6IjEyOThoZmFrMDlra2xqYWRza2YifQ==","uIUWny1o4YfCUjBEB+cIq0e0lSAHy7Fqk4OxoCSSM88XW24M01wQj5d9wXCPxbGSeLIsxl6FEeS36cLzCuasHq2vsUQ3K/ySudP605sG+a2hrK5PC99IjknmA06hwMnEMFsPURLft6wRuN92UmKgEnSHejVyXj1gExpEL72xQKA="]
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
        "message": "1d7111c9-8f9f-490f-87d8-197032c38ef8"
      },
      "id": 3
    }     

#### queryBargain bargain01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"queryBargain",
    "args":["bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    }  

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a\",\"name\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"detail\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"startTime\":\"2017-01-01\",\"endTime\":\"2018-01-01\",\"partyA\":\"某基金会\",\"partyB\":\"某地区XX县XX村施工队\",\"depositBank\":\"光大银行\",\"bankAccount\":\"1298hfak09kkljadskf\"}"
      },
      "id": 5
    }    

#### queryBargains bargain01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"queryBargains",
    "args":["bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"]
    },
    "secureContext":    "jim"
    },
    "id": 5
    } 

result:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "[{\"addr\":\"bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a\",\"name\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"detail\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"startTime\":\"2017-01-01\",\"endTime\":\"2018-01-01\",\"partyA\":\"某基金会\",\"partyB\":\"某地区XX县XX村施工队\",\"depositBank\":\"光大银行\",\"bankAccount\":\"1298hfak09kkljadskf\"}]"
      },
      "id": 5
    }       

### donor register/query addContribution addTrack
#### registerDonor donor01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"registerDonor",
    "args":["channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a","eyJhZGRyIjoiZG9ub3IwMToyNzVlNzRiMGUzNDBmNTQxMzU0OTZlNDZkODI5YjI1YWY2OTk5ODRlNjc4N2Y5YTdiMTMxOTFhZDk5MWExZWIxIiwibmFtZSI6ImRvbm9yVXNlcjAxIiwidG90YWwiOjEwMDAwMDAwMH0=","aIJUrEL50jVOisSW/vEUjUa+mrqGV7DwqaqJ54X3+UCUVNsrF8UhG+kx8aocpKcgYNIizLcxPIWmwUiioaUEtZPorxJADpxUVhr+jeZOJcbPKZQhzSlOBiZoC9c0iSm4i/6iP/xe7hUvEg7RUKlK8B1ZcckbKa7woevRYUU7Ixk="]
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
        "message": "eaa5b462-0876-4766-b992-034a8d461895"
      },
      "id": 3
    }     

#### queryDonor donor01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
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

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1\",\"name\":\"donorUser01\",\"total\":100000000,\"trackings\":[{\"name\":\"donorUser01\",\"accountName\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"accountAddr\":\"bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a\",\"amount\":100000000,\"timestamp\":1486630616}]}"
      },
      "id": 5
    }

#### addContribution donor01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"addContribution",
    "args":["channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a","donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1","eyJuYW1lIjoiZG9ub3JVc2VyMDEiLCJzbWFydENvbnRyYWN0TmFtZSI6IuWugeWkj+ilv+mDqOWcsOWMuuavjeS6suawtOeqlumhueebriIsInNtYXJ0Q29udHJhY3RBZGRyIjoic21hcnRjb250cmFjdDAxOjFkNTRhODcxMzkyM2FmMTcxOGU4ZWVhYmVjM2U0ZDg1OTZkYmJkZjJkYTNmNjllYTIzYWViOGM3YTVhYjczZDgiLCJhbW91bnQiOjEwMDAwMDAwMCwidGltZXN0YW1wIjoxNDg2NjMwNjE2fQ==","aIJUrEL50jVOisSW/vEUjUa+mrqGV7DwqaqJ54X3+UCUVNsrF8UhG+kx8aocpKcgYNIizLcxPIWmwUiioaUEtZPorxJADpxUVhr+jeZOJcbPKZQhzSlOBiZoC9c0iSm4i/6iP/xe7hUvEg7RUKlK8B1ZcckbKa7woevRYUU7Ixk="]
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
        "message": "9c47c9ff-b2f0-41b7-8339-ed35259628ce"
      },
      "id": 3
    }         

#### addTrack donor01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"addTrack",
    "args":["channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a","donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1","eyJuYW1lIjoiZG9ub3JVc2VyMDEiLCJhY2NvdW50TmFtZSI6IuWugeWkj+ilv+mDqOWcsOWMuuavjeS6suawtOeqlumhueebrlhY5Y6/WFjmnZHmsLTnqpYiLCJhY2NvdW50QWRkciI6ImJhcmdhaW4wMTo4ZmNjNThlYTdlZDIxMmY3YzFiYTM1OWQxNWJlYTE0NGU2N2MzOTAwNDRkOTUzNzk3NTQ4Y2Y2N2ZkNjI1MzRhIiwiYW1vdW50IjoxMDAwMDAwMDAsInRpbWVzdGFtcCI6MTQ4NjYzMDYxNn0=","v1GVprhF5KLSPIquJ1IEo593DHzs4x2uTL3fgt2kI+3izfbe/Q43wHZ5Y6YSWfDQu1VmVPA7Sl85dHimdEeJURVvL/3EkIxGJNbdUVvGeKM9ASU2jUXTJ+SFgygglnOf+X+GWZVAYFy47ZKMki7xxUOTTLOUNMXWvtZE6frw4N8="]
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
        "message": "9781df01-1c4c-45ce-add6-f147fabd43ef"
      },
      "id": 3
    }                      

### coinbase

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"coinbase",
    "args":["cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924","eyJ2ZXJzaW9uIjoxNzAxMDEsInRpbWVzdGFtcCI6MTQ4NjYzMzY2NiwidHhpbiI6W3sic291cmNlVHhIYXNoIjoibnVsbCIsImFkZHIiOiJudWxsIn1dLCJ0eG91dCI6W3sidmFsdWUiOjEwMDAwMDAwMDAwMDAsImFkZHIiOiJjZWJiYW5rOjI5NzMxZDBlNmM2Y2E5Y2I5ODVlYWJmOWZlNzE2ZDE2NDRjNjI0Y2FlNTI2NWMzNmM5YjdhNDY3MDIwMDM5MjQiLCJzaWduIjoiT0pIZjJZSENMTFpaTVVtTkFkemJ4aHQ4RVJjSHRVVWtaOVJ6SmRtWFhjWXN3NHM1WCt6SEtramhhRSs0Njh2K0N4aG1ZY3hrWm5LWXR3N0V5eDVYcythMjJBWTZYRm1GQmZaYUFoRWh0TmJUcC94SWtiVThiYzBrSlNSMml2ZlZmeUdad0hoM1FPUU5ib3BFdXVTQkZLUytiaUUzYkswaGx0d251T0JUcGlrPSJ9XSwiSW5wdXREYXRhIjoiZG9ub3J1dWlkIiwiZm91bmRlciI6ImNlYmJhbmsifQ==","h4Tlf+YCl//ga/2HNPs2DbH/rvdg7DhXv2GSSi903WNwqqKY+qiMwDvT095WvCKkoHl8Kdne67o92HlX8u0oC8tL7OpANxX6a3kMmw9x3RbSbaLKL+DQS1wenS7jpYC8Xfpmo3hRbS2PoowOoBQysZa+UxVdsg+cBNmF4WpA2C8="]
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

### changeCoin cebbank -> donor01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"changeCoin",
    "args":["cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924","eyJ2ZXJzaW9uIjoxNzAxMDEsInRpbWVzdGFtcCI6MTQ4NjYzODI0NiwidHhpbiI6W3sic291cmNlVHhIYXNoIjoiYWRiZTE5NjI0OWRiYWRhYWNhOWNhMzVjYTFmYzMzNTUwOTY5NjkxZDA4ODlmNzI5MTY1OGQ0YzBiZjkzN2Y0YyIsImFkZHIiOiJjZWJiYW5rOjI5NzMxZDBlNmM2Y2E5Y2I5ODVlYWJmOWZlNzE2ZDE2NDRjNjI0Y2FlNTI2NWMzNmM5YjdhNDY3MDIwMDM5MjQifV0sInR4b3V0IjpbeyJ2YWx1ZSI6OTk5OTAwMDAwMDAwLCJhZGRyIjoiY2ViYmFuazoyOTczMWQwZTZjNmNhOWNiOTg1ZWFiZjlmZTcxNmQxNjQ0YzYyNGNhZTUyNjVjMzZjOWI3YTQ2NzAyMDAzOTI0Iiwic2lnbiI6IkttUWxjQVJCOHcwMUlsT1M5ZGNCcEs4dStPZ1dMQVFuTDVuMXlpNDhYaGpvU1ZaMDJFNU03NlJvUFRRQURIQm91RVdLVnpYWTlMUFhUNzdjY0tTUWpGeURXQm9YNUxFVUpnNXV2ZXJWalZYanl2RFlVUk9EV1FCS21IZURyNFlVd3AyT0FuWGpFUnRXQWZHNCsrS2pMN05wbG1WV2VrZTMzZk92K2NJNjMwVT0ifSx7InZhbHVlIjoxMDAwMDAwMDAsImFkZHIiOiJkb25vcjAxOjI3NWU3NGIwZTM0MGY1NDEzNTQ5NmU0NmQ4MjliMjVhZjY5OTk4NGU2Nzg3ZjlhN2IxMzE5MWFkOTkxYTFlYjEiLCJzaWduIjoiUURFdFdKTng2N0J1V3JGdFNkWU52d01RL25UMWtaUWQ5ZGdtR0YyQ3dhT1FnbTZuUEF5eGlHRnBGNUpScXVoSjNydk5rellNTUJLL2FHbzk5cUlqVjZ1WktkRzFuL0dzd0ozUzlkWlB5cUpYQzAwa014dmtOTzFLZTk0UnF3Zm9ZUTkzTzlHRWRtdGhBTkE3cXZHNUozazZ2ZnVJcDBJTUxEenFydkRjdExRPSJ9XSwiSW5wdXREYXRhIjoiZG9ub3J1dWlkIiwiZm91bmRlciI6ImVyaWMifQ==","pE5Eq1umWQoOdb4/kxucjZpyKz4n4oq7uvbRZWL8gfBEc2O2+PapwqflGl3EBfoAH26z8nLmxmUqwfilhK67T2krYoqGYyY+wnjfd3eypZ/FOqpm7IplnZVYzrdX3ybjvF1DlYSN83z1Dvv348yY19cclZdHzI/TBwdCwxsV6e0="]
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

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924\",\"balance\":999900000000,\"rsaPublicKey\":\"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDenFYaVVyRFZzanhIZzdyM1RMc1NSaGlaTApqeE1nczBIU3JsRlV3c0s1eFVCcGFwSHdSazlHeGJrMWtOd2tkSzdPeHlselJjbmxzbEd5VnhhU21KYzNqbmpvClVRVFphbVRaemViMzNNZC9oYUVkN3BhSXFkSS9pZ1Z6TEtpaStXcFJkUE02VlFaV0pHam01eEMwMHhWUWw3aWcKOExBUXV6ek1OMnNCRkhsSmN3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=\",\"txouts\":{\"17e0e8d20d130dc751cc5203db00e49b52568ce1edc3d8f6116a6b4cef78e08a:0\":{\"value\":999900000000,\"addr\":\"cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924\",\"sign\":\"KmQlcARB8w01IlOS9dcBpK8u+OgWLAQnL5n1yi48XhjoSVZ02E5M76RoPTQADHBouEWKVzXY9LPXT77ccKSQjFyDWBoX5LEUJg5uverVjVXjyvDYURODWQBKmHeDr4YUwp2OAnXjERtWAfG4++KjL7NplmVWeke33fOv+cI630U=\"}}}"
      },
      "id": 5
    } 

query donor01:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1\",\"balance\":100000000,\"rsaPublicKey\":\"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERS9HZ3VWbUk0QmRnUU5oYThPMmtYK3dHQgpNSkxQbGNPZllLMUdsQkpsZzlLUm9RNmZrWGd5aXRCb2FORzI5dVpnV3loVjZSMmhzUXpmRyt3cHE0VmsyQkEyClBxZEtFcVRyQVhMU0EvWkw0bUpPVEx4bjV0Qm5QL0p5OGRGaitXSjBlSWtzOGZhMUxORS9QVnR6UDlEZDV2bmEKbEZvTXNJa09maDlMcWxmaFRRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=\",\"txouts\":{\"17e0e8d20d130dc751cc5203db00e49b52568ce1edc3d8f6116a6b4cef78e08a:1\":{\"value\":100000000,\"addr\":\"donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1\",\"sign\":\"QDEtWJNx67BuWrFtSdYNvwMQ/nT1kZQd9dgmGF2CwaOQgm6nPAyxiGFpF5JRquhJ3rvNkzYMMBK/aGo99qIjV6uZKdG1n/GswJ3S9dZPyqJXC00kMxvkNO1Ke94RqwfoYQ93O9GEdmthANA7qvG5J3k6vfuIp0IMLDzqrvDctLQ=\"}}}"
      },
      "id": 5
    }  

### donated donor01 -> smartcontract01,fund01,channel01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"donated",
    "args":["channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a","donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1","eyJ2ZXJzaW9uIjoxNzAxMDEsInRpbWVzdGFtcCI6MTQ4NjYzODQ3OSwidHhpbiI6W3siaWR4IjoxLCJzb3VyY2VUeEhhc2giOiIxN2UwZThkMjBkMTMwZGM3NTFjYzUyMDNkYjAwZTQ5YjUyNTY4Y2UxZWRjM2Q4ZjYxMTZhNmI0Y2VmNzhlMDhhIiwiYWRkciI6ImRvbm9yMDE6Mjc1ZTc0YjBlMzQwZjU0MTM1NDk2ZTQ2ZDgyOWIyNWFmNjk5OTg0ZTY3ODdmOWE3YjEzMTkxYWQ5OTFhMWViMSJ9XSwidHhvdXQiOlt7InZhbHVlIjo5OTUwMDAwMCwiYWRkciI6InNtYXJ0Y29udHJhY3QwMToxZDU0YTg3MTM5MjNhZjE3MThlOGVlYWJlYzNlNGQ4NTk2ZGJiZGYyZGEzZjY5ZWEyM2FlYjhjN2E1YWI3M2Q4IiwiYXR0ciI6ImRvbm9yMDE6Mjc1ZTc0YjBlMzQwZjU0MTM1NDk2ZTQ2ZDgyOWIyNWFmNjk5OTg0ZTY3ODdmOWE3YjEzMTkxYWQ5OTFhMWViMSIsInNpZ24iOiJVcTVNaUE0Smdad2FINGJmeDFHd1dFR1ZCRCtMTmMxSW9hYVZ2TFFlRjR6TlRRbGo3TStReHZZc3QyWDltN3pBYVArV3VNL3VmbTA4Tk1qTmM0VkY4ZmdsQVFKSUZ4eEYxTzYvUXhMUk8rMDRwSFVsYVVBSUtzVnZ0ZFE1dmdMU25vNWEwcENObW5aTWtoOUt1RlZQaVBCL0wwQ3FqN0lobDhsVitqY2lCYjA9In0seyJ2YWx1ZSI6MjAwMDAwLCJhZGRyIjoiY2hhbm5lbDAxOjljOGI0M2NlOTQ4MDEwZWZjM2I3ZDEwMmFhZTUwMjE2NWNjZDVlMDcxNGEzZTc2NWZlMWE4ZjQ0NDkzNjc4NWEiLCJhdHRyIjoiZG9ub3IwMToyNzVlNzRiMGUzNDBmNTQxMzU0OTZlNDZkODI5YjI1YWY2OTk5ODRlNjc4N2Y5YTdiMTMxOTFhZDk5MWExZWIxIiwic2lnbiI6Ind3V005M1dyQ3RNYXBYbytiTWlYQ3dmS2ppYjZ4cUFaa3kyTWUyUzVmRWtXaGdMczVMLzFueGFRWXJvd0ZQei9HNUVTajdLQWVsMUw2aXZtRDlQeEN5TmpKWFR4RTZQZTQ3dUo1TnUwcWVvTy85YmI0NmluYW4yaWN1WTY0a1lCQmRkaHBKcXZnd0FMSGtWR3A2MGtpYnFDWUY4dDlsMWtWcEtQYk11YUJTND0ifSx7InZhbHVlIjozMDAwMDAsImFkZHIiOiJmdW5kMDE6MjVhYjU4MGEyMDkzNzc2Y2EyZTFkZDE3NzVlOTZkZmVjNWYxZmZiY2M5NTY1MTI5MzUxY2IzMzBjZjA3MTJkNyIsImF0dHIiOiJkb25vcjAxOjI3NWU3NGIwZTM0MGY1NDEzNTQ5NmU0NmQ4MjliMjVhZjY5OTk4NGU2Nzg3ZjlhN2IxMzE5MWFkOTkxYTFlYjEiLCJzaWduIjoieWdQcEZ4VCtpNXdqV3hYNGFjQVMvck9OWUpPUDlIUEZBM0Z2c2FFVWtsSE1HbXpXb1NDRFBQVUlLaTh6U3VhT1pPQTJPZHRFK2NTOFdXTGwwMjVEQUJUNjJmdHlEaVo1eHJ6dEdwbXR3TWlaV0ozKzNvRm44M1U2SlJOcWJuMngvZkdUblZseGZvSlRDNEFZNHcvcVRCaXFhVnZBSFRFUW9sby83SFcrL3U0PSJ9XSwiSW5wdXREYXRhIjoiZG9ub3J1dWlkIiwiZm91bmRlciI6ImNoYW5uZWwwMSJ9","kdJ7wPilpeEqywra+3k2qtJShe/CMIQz5If/WaX95oqT9MMSEKdFJ+7CxPd9XWrfmCIrfaUBh8m5Ma1bWFhYzWPXzvvXXGwlH030n96JETyM4YZm93fbCIqt5AujrjzxCOvxdE6vg65WIsol4fWKqDlTW+HU9VUIY7U9725aGnY="]
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
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
    },
    "ctorMsg": {
    "function":"drawed",
    "args":["fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7","eyJ2ZXJzaW9uIjoxNzAxMDEsInRpbWVzdGFtcCI6MTQ4NjYzODc3NCwidHhpbiI6W3sic291cmNlVHhIYXNoIjoiZWQyMWM4NTdkYTgyMWU4MjczODBjOGZhYWZhOTlkOTRkNDEwMmEwOGZmMWE3ZTY0YjA1MTgzYzQyZjZkYjIyOCIsImFkZHIiOiJzbWFydGNvbnRyYWN0MDE6MWQ1NGE4NzEzOTIzYWYxNzE4ZThlZWFiZWMzZTRkODU5NmRiYmRmMmRhM2Y2OWVhMjNhZWI4YzdhNWFiNzNkOCJ9XSwidHhvdXQiOlt7InZhbHVlIjo5NTAwMDAwLCJhZGRyIjoic21hcnRjb250cmFjdDAxOjFkNTRhODcxMzkyM2FmMTcxOGU4ZWVhYmVjM2U0ZDg1OTZkYmJkZjJkYTNmNjllYTIzYWViOGM3YTVhYjczZDgiLCJhdHRyIjoiZG9ub3IwMToyNzVlNzRiMGUzNDBmNTQxMzU0OTZlNDZkODI5YjI1YWY2OTk5ODRlNjc4N2Y5YTdiMTMxOTFhZDk5MWExZWIxIiwic2lnbiI6InpFUWVnTWk3QS9CNW9DbElEUVRldS9XY0twbFF0NWlOeFRnQVI1ZDFjTWFMOFhPeElLZFRQdDFnNHFhT3hjYmtscGdRNFVNUldRWjArbU12aW1OakNqdXBCMUpqamlnYWRwVHBCTTk5bGpnVml4dy8wcTl6RXRiYSsvUnh6eUVJd1BpOGwzQnRVMEw3QXNUcDh0Z1lhV0xVQ1BpMmpRWExLRWdMdWlsV1lYZz0ifSx7InZhbHVlIjo5MDAwMDAwMCwiYWRkciI6ImJhcmdhaW4wMTo4ZmNjNThlYTdlZDIxMmY3YzFiYTM1OWQxNWJlYTE0NGU2N2MzOTAwNDRkOTUzNzk3NTQ4Y2Y2N2ZkNjI1MzRhIiwiYXR0ciI6ImRvbm9yMDE6Mjc1ZTc0YjBlMzQwZjU0MTM1NDk2ZTQ2ZDgyOWIyNWFmNjk5OTg0ZTY3ODdmOWE3YjEzMTkxYWQ5OTFhMWViMSIsInNpZ24iOiJCeVB4T1A2Zi9EQnpRc2RWTmE1UURRY2tzVW5mUmdFTTF5aklYMWhqRGR1NnJ3Z2lQamRqbk10ZTJSMUl6YmcrQkR6amRJbmVpWWxybXR3eVUrUzFIczJkcUdGalRWc3lBVSsxeUs5eGZpdmYwcFEzN3IyeFUxMk00TnVkQkZKSmVvVUs3LzdjMEl0RTB1Mmh0UUprUWlxNUFISkt2S2pDVHBmNjRWV0Fnbkk9In1dLCJJbnB1dERhdGEiOiJkb25vcnV1aWQiLCJmb3VuZGVyIjoiZnVuZDAxIn0=","RLryJqxiyUhgB6u26D7xvGKUYF3YX1RP7OpSAZDuwKsTsEybsieQTgxGmOHrjgvhHmDhP/ud3J7hXy7qhqAHVX7AQUTTbq+gp2JSWXyNPVwHQKWcSF1npTqao4BuAwAoDihMotcCsMrcSDUxfdgPSMcH3nANVBZLJPSyhq0GrAY="]
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
    "name":"56cca7777e5e066445740e515a8eb1dbfa563c3493b8d9209bd4f08393fdf1c38f90632eb99154dae4812fa3d00ccb8d395953772b3d08ffa480c5e225b3732c"
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







           
