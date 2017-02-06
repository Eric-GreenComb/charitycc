## Prepare

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

result:
    
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
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
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"registerBank",
    "args":["cebbank","LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDK0oxWjRiTGFPOFNobTFyOXFOOHB6RjJxbwp1c08rSnhoOWpGcDdQTkcwMERGTUR5RUhWNU9JajhhdFlVdzBwRkFTUG95dldHMGhlRzdCVU00bVRpWHNZOFF3ClpYMjZ6L2I2bk05Q3ZtM2xlell1NjgwT2NQQnNIczdLb0RZTkUrUGM2c0EwRGVDVnUxNm5aeGpQbFRwbUxMdmkKb2Z3bS9ReVlKZlVsNElRazFRSURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=","TV1xznZQSfCliJBdTvKNhn5s1M8Ji9tuET68hrMQHpbd5LO83K7yHyMJQQo7lX4iV0qVvYyK5p8zeJwVVyOeNQd1jRpi83XmfLqWHzhddf3zK81xOpzz+2GVuZolyKS2dzXyheJ7holm3PyxUwZtL5+qWfQhAXBKLqN0LrU5pcE="]
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
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryBank",
    "args":["cebbank:71266a9913375add341f1f7a4a606a2ee4e24efaf59bb2a3eb4b71aaaf5fb3fe"]
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

#### registerAccount user01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["user01","LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDcGxmMDlJNTl5SDYzRnJwSlhqMSs0WmhjZgpSOXJVVi82akhWWVAxelQ1aHlUZUwzNWQ4NVpwTzhOTDdWV2MyYVpHeDFjZ2dvRG8rRUkvMVJobDA4UHhSbFdKCi9XVGt1VzBLbDIxeGRGUXFnSnk1Q3JUeUc0SURmS2NTRkg3dG9tZzFqeVlHaHRFZ2g5TUI3RVVZNXd0M0hWZHYKWEdWQWlMbzViU3VOaVRVV2J3SURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=","qWdNBVo7E0kF3snN35CGR51oayyohZ8WOEZ4cPU2TINtohoE0epmtlZUFvXcf5xn/ZiNuapwQWf0xXH9wkaic4uM1DttoJKYjbvW2q9lfMJfa+zlmmEJjY2UauS9u9LxXTwzfgQUt9Aq5UkWfMuwrw9OfosfaLdHoPK3nkwDFFM="]
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

#### queryAccount user01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["user01:e9629822204b9e329a767f5c4bf399b77fa0dfa63c9df7b88b040c5003c2af3a"]
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
        "message": "{\"addr\":\"user01:e9629822204b9e329a767f5c4bf399b77fa0dfa63c9df7b88b040c5003c2af3a\",\"rsaPublicKey\":\"LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDcGxmMDlJNTl5SDYzRnJwSlhqMSs0WmhjZgpSOXJVVi82akhWWVAxelQ1aHlUZUwzNWQ4NVpwTzhOTDdWV2MyYVpHeDFjZ2dvRG8rRUkvMVJobDA4UHhSbFdKCi9XVGt1VzBLbDIxeGRGUXFnSnk1Q3JUeUc0SURmS2NTRkg3dG9tZzFqeVlHaHRFZ2g5TUI3RVVZNXd0M0hWZHYKWEdWQWlMbzViU3VOaVRVV2J3SURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=\"}"
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
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["channel01","LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDbFBibUthUk5FVm05NHUwSmZ1anVwd1ZVSgo5MVc3YmQyODc4NHdDbS9HUi85Wm1jRzZkRlNRUC9yaTQ5RHExR0hLekJEd0lvRzVXOGJmZndLeERiOTIzOWsrCmxwaHV1VjloZ0I3TlliQ25VUUo0WmkvUUlQZDFqR2pSem1DYlcvajJOdDJid1JJL3loNHNMbVhMb2Qzb2FuczcKcjR5ZzhnN2k4blRZUHlLWllRSURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=","MRCrGTmkoGLx+IQgWWNsjPv6u4SH+YbZRfS5JZtBBFFcXXpVEIeO651oynD+JvCkJyRGjFCyx5swBeBw2xXS0EYDMxyQbLASJQKy42o6sd6L9k+ael+EaJ6ypIY8PyBs3QfZg9Nf53b65xyEaSbzCWLxHUWs/QKg0T003sBXtvk="]
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
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["channel01:cc68d3ccb7be906dbf9d63e4ae3e9605c8ddd63fb710603e3cb45976932db9f4"]
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
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["fund01","LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDZ0dDTEdVN1hmbFQ2UjNQdW1NN3dPRjJmdQowbUhmREl6dlA3SkUxUEJvdXR3NzkvT0F6NWJnWXIrQUtwM2VhSFhlemlIb3JTQTZFUFNDUGNoTm01ZzFsOCt4CjQ0ODJ0dm4zWVA3YWdxTmJISHI2MjNwOEprOXEwWFVhS3U0OXVEMDNjLy9EcnM2bjVjWE9QdVF4VVFkQ3VnYTkKT1FVR0RDTEs4RElUVzdXU1J3SURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=","o7JiiwKn2MNSsa9qsXIHY0EcFX0fZ90ewz8d0FrsKbf33hBdllYl2nzDScg4UmVkwuQqbUkyn+viGzTqOsYAaeZ+zGCtaswjpw+Iac01WSfjkbJ67kiduCBDW9EjYbIyHE6dnkm/c3DQdghSz4T8rE+vulH5P1lQgs6CZ0K8Yhw="]
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
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["fund01:da336a2a2d81cdf570b5d3d70926144997e5e277f49897c6885acacba8b1cc13"]
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

#### registerAccount treaty01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["treaty01","LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDeHV2RXZ5OTd4OUpiR3ZVMzlBMTZIZ3FTNgo1WDhGaGFMNzh5Wm1QSE9DZnNjY3paZk4ydE5uaHlqR3dvb1Zwbmd5bFJ5d3Bib29BekFoeVNLcVZROGdpNjBHCnRJa01QWEZ0RHZyVFJOUFZhd2MwdTloVVY5SnhSdkx1c0JjSzBmSGtkUGhuamlTa25nNGk3Mk51OUJFUlFveGMKNTQzOVU2aEpLZE45SUdta2lRSURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=","qsE9futAjI7wJng4CAm9b9gkttEYIXE23PrWNnTAup75tciPSPzaSq3/usZ0fUrCWxHPACpvmqCmk6UrZ+pxn8DJITEyy+ediVYIaY83Rrm2plfB+D78FfGDX9VkzwBnWmfeK3dKQyjau/HsbsBCPTsZTPpOOsdkP2kGw1RalRQ="]
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

#### queryAccount treaty01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["treaty01:e4811d624ebbacbaadc59bf0192c4229a9a4e05e1fa269dde345bf2bfb3f9bdd"]
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

#### registerAccount contract01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"registerAccount",
    "args":["contract01","LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDcDlxUTBINVY4WUkwOWFGMTJvVnhjWTlVUQpuUzNVRFZwT0c2OXo3a0lzK3o3WUJXR0hlVWtKSjk2U2JrVVU1WWtoU0JhbXZoTXBoUW9STS90MlI5YTJpem5FCkNyWmladVpHZmxJdEFtU0lxOXB5Ris5M1UwYVdQWTcxc1VqSCtnQ0orcFBPSFJwSWFoRWtHZ2t1QTN1NU5wZncKUTBJQnJGOHdNWi9NSnhNSFl3SURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=","Qt1K73w/cmz7ueNyNSshFnRJ1lr2kQamgisdu/nB2F+OI70R1GNm833iLMfty4VZtuFI+oi70BdzFO/ZM4JDjHgWNYm9R9V92iRHsK+2F7FwFlR/8VvwFIkBF0METfZQBVF5fcya+Vk0qhfLuNDl8YkWS7SQW4Jkeodq12+XdO8="]
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

#### queryAccount contract01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryAccount",
    "args":["contract01:8017d1bd15326b48b7020ff6fe434ae4999cf050025ec64871516677f1177460"]
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

### treaty register/query
#### registerTreaty treaty01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"registerTreaty",
    "args":["fund01:da336a2a2d81cdf570b5d3d70926144997e5e277f49897c6885acacba8b1cc13","eyJhZGRyIjoidHJlYXR5MDE6ZTQ4MTFkNjI0ZWJiYWNiYWFkYzU5YmYwMTkyYzQyMjlhOWE0ZTA1ZTFmYTI2OWRkZTM0NWJmMmJmYjNmOWJkZCIsIm5hbWUiOiLlroHlpI/opb/pg6jlnLDljLrmr43kurLmsLTnqpbpobnnm64iLCJkZXRhaWwiOiLlroHlpI/opb/pg6jlnLDljLrmr43kurLmsLTnqpbpobnnm64iLCJnb2FsIjoxMDAwMDAwMDAwMDAwLCJwYXJ0eUEiOiLmn5Dln7rph5HkvJoiLCJwYXJ0eUIiOiLmn5DlnLDljLoiLCJjcmVhdGVUaW1lc3RhbXAiOjE0ODYxOTY4NjgsImZvdW5kYXRpb24iOiJmdW5kMDE6ZGEzMzZhMmEyZDgxY2RmNTcwYjVkM2Q3MDkyNjE0NDk5N2U1ZTI3N2Y0OTg5N2M2ODg1YWNhY2JhOGIxY2MxMyJ9","JNU66rJPa2qPmy0NCsvyZKLYG51bVkUu1hPqsIogbax5owoRAmz+q8lLzLsyHv7MfHufl5YfLjR4U5NxDXML2h8+++Q6k7vKKx1d/e6cgaDbHiuuP8OAKPvnnuskpjKjCKEFnOEcApxGqOTr6/YvxcRkpoBvoIyecEMPBA3euYo="]
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

#### queryTreaty treaty01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryTreaty",
    "args":["treaty01:e4811d624ebbacbaadc59bf0192c4229a9a4e05e1fa269dde345bf2bfb3f9bdd"]
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
        "message": "{\"addr\":\"treaty01:e4811d624ebbacbaadc59bf0192c4229a9a4e05e1fa269dde345bf2bfb3f9bdd\",\"name\":\"宁夏西部地区母亲水窖项目\",\"detail\":\"宁夏西部地区母亲水窖项目\",\"goal\":1000000000000,\"partyA\":\"某基金会\",\"partyB\":\"某地区\",\"createTimestamp\":1486196868,\"foundation\":\"fund01:da336a2a2d81cdf570b5d3d70926144997e5e277f49897c6885acacba8b1cc13\"}"
      },
      "id": 5
    }      

#### queryTreaties treaty01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryTreaties",
    "args":["treaty01:e4811d624ebbacbaadc59bf0192c4229a9a4e05e1fa269dde345bf2bfb3f9bdd"]
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
        "message": "[{\"addr\":\"treaty01:e4811d624ebbacbaadc59bf0192c4229a9a4e05e1fa269dde345bf2bfb3f9bdd\",\"name\":\"宁夏西部地区母亲水窖项目\",\"detail\":\"宁夏西部地区母亲水窖项目\",\"goal\":1000000000000,\"partyA\":\"某基金会\",\"partyB\":\"某地区\",\"createTimestamp\":1486196868,\"foundation\":\"fund01:da336a2a2d81cdf570b5d3d70926144997e5e277f49897c6885acacba8b1cc13\"}]"
      },
      "id": 5
    }         

### contract register/query
#### registerContract contract01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"registerContract",
    "args":["fund01:da336a2a2d81cdf570b5d3d70926144997e5e277f49897c6885acacba8b1cc13","eyJhZGRyIjoiY29udHJhY3QwMTo4MDE3ZDFiZDE1MzI2YjQ4YjcwMjBmZjZmZTQzNGFlNDk5OWNmMDUwMDI1ZWM2NDg3MTUxNjY3N2YxMTc3NDYwIiwibmFtZSI6IuWugeWkj+ilv+mDqOWcsOWMuuavjeS6suawtOeqlumhueebrlhY5Y6/WFjmnZHmsLTnqpYiLCJkZXRhaWwiOiLlroHlpI/opb/pg6jlnLDljLrmr43kurLmsLTnqpbpobnnm65YWOWOv1hY5p2R5rC056qWIiwic3RhcnRUaW1lIjoiMjAxNy0wMS0wMSIsImVuZFRpbWUiOiIyMDE4LTAxLTAxIiwicGFydHlBIjoi5p+Q5Z+66YeR5LyaIiwicGFydHlCIjoi5p+Q5Zyw5Yy6WFjljr9YWOadkeaWveW3pemYnyIsImRlcG9zaXRCYW5rIjoi5YWJ5aSn6ZO26KGMIiwiYmFua0FjY291bnQiOiIxMjk4aGZhazA5a2tsamFkc2tmIn0=","H4EEzd7UShFKmLwfq66yIXC4db3JMQh5ysdlafADd+DRHjzp7TyNnuhIHNgl5f3djS5rV6Mlyro3k9HyXv2VJcABtpGeEnphO6v8Fh8slMYtqXJdcVrmSCWUD9d+IzhrpNe+Wgv6bMFhy9iYECqkLTHUJnRxWopl7Pi6LHMyXr0="]
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

#### queryContract contract01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryContract",
    "args":["contract01:8017d1bd15326b48b7020ff6fe434ae4999cf050025ec64871516677f1177460"]
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
        "message": "{\"addr\":\"contract01:8017d1bd15326b48b7020ff6fe434ae4999cf050025ec64871516677f1177460\",\"name\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"detail\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"startTime\":\"2017-01-01\",\"endTime\":\"2018-01-01\",\"partyA\":\"某基金会\",\"partyB\":\"某地区XX县XX村施工队\",\"depositBank\":\"光大银行\",\"bankAccount\":\"1298hfak09kkljadskf\"}"
      },
      "id": 5
    }    

#### queryContracts contract01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryContracts",
    "args":["contract01:8017d1bd15326b48b7020ff6fe434ae4999cf050025ec64871516677f1177460"]
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
        "message": "[{\"addr\":\"contract01:8017d1bd15326b48b7020ff6fe434ae4999cf050025ec64871516677f1177460\",\"name\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"detail\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"startTime\":\"2017-01-01\",\"endTime\":\"2018-01-01\",\"partyA\":\"某基金会\",\"partyB\":\"某地区XX县XX村施工队\",\"depositBank\":\"光大银行\",\"bankAccount\":\"1298hfak09kkljadskf\"}]"
      },
      "id": 5
    }       

### donor register/query addContribution addTrack
#### registerDonor user01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"registerDonor",
    "args":["channel01:cc68d3ccb7be906dbf9d63e4ae3e9605c8ddd63fb710603e3cb45976932db9f4","eyJhZGRyIjoidXNlcjAxOmU5NjI5ODIyMjA0YjllMzI5YTc2N2Y1YzRiZjM5OWI3N2ZhMGRmYTYzYzlkZjdiODhiMDQwYzUwMDNjMmFmM2EiLCJuYW1lIjoiZG9ub3JVc2VyMDEiLCJ0b3RhbCI6MTAwMDAwMDAwfQ==","V7MTYzD/SOjprKj0uHqeuhk0iIXfgsu1cQW2UQgFbvoa1kJINTkEjXO9ZQWEDxHPPpfld0CBQQs2VLyLRIUr7ZSlQ2K31otyyO6Smj5h8l9b4yX80vIVliC5Hs2XCpnhpkbwb7MxU91utRwaSja3sEsGn0YXjk3La54L8kCPGS4="]
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

#### queryDonor user01

    {
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"queryDonor",
    "args":["user01:e9629822204b9e329a767f5c4bf399b77fa0dfa63c9df7b88b040c5003c2af3a"]
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
        "message": "{\"addr\":\"user01:e9629822204b9e329a767f5c4bf399b77fa0dfa63c9df7b88b040c5003c2af3a\",\"name\":\"donorUser01\",\"total\":100000000,\"contributions\":[{\"name\":\"donorUser01\",\"treatyname\":\"宁夏西部地区母亲水窖项目\",\"treatyaddr\":\"treaty01:e4811d624ebbacbaadc59bf0192c4229a9a4e05e1fa269dde345bf2bfb3f9bdd\",\"amount\":100000000,\"timestamp\":1486364970}],\"trackings\":[{\"name\":\"donorUser01\",\"contractname\":\"宁夏西部地区母亲水窖项目XX县XX村水窖\",\"contractaddr\":\"contract01:8017d1bd15326b48b7020ff6fe434ae4999cf050025ec64871516677f1177460\",\"amount\":100000000,\"timestamp\":1486364970}]}"
      },
      "id": 5
    }

#### addContribution user01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"addContribution",
    "args":["channel01:cc68d3ccb7be906dbf9d63e4ae3e9605c8ddd63fb710603e3cb45976932db9f4","user01:e9629822204b9e329a767f5c4bf399b77fa0dfa63c9df7b88b040c5003c2af3a","eyJuYW1lIjoiZG9ub3JVc2VyMDEiLCJ0cmVhdHluYW1lIjoi5a6B5aSP6KW/6YOo5Zyw5Yy65q+N5Lqy5rC056qW6aG555uuIiwidHJlYXR5YWRkciI6InRyZWF0eTAxOmU0ODExZDYyNGViYmFjYmFhZGM1OWJmMDE5MmM0MjI5YTlhNGUwNWUxZmEyNjlkZGUzNDViZjJiZmIzZjliZGQiLCJhbW91bnQiOjEwMDAwMDAwMCwidGltZXN0YW1wIjoxNDg2MzY0OTcwfQ==","hWrKIlKgygzHXegQ1lMgSpOMAjJ4985Qt8cJPc0TPfZZDhLXLExpPyCNjCV79yI3UjwWFVdLvSz1lWjcBKJUQU4aDKKf/0z04gWoYB6oecpsvp9X89E1eFRs9eUpS/tEH5S9YJ5kAtnJX0fDOKovCNPFFg6AXanUn8vXjZ6WZ7Y="]
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

#### addTrack user01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"addTrack",
    "args":["channel01:cc68d3ccb7be906dbf9d63e4ae3e9605c8ddd63fb710603e3cb45976932db9f4","user01:e9629822204b9e329a767f5c4bf399b77fa0dfa63c9df7b88b040c5003c2af3a","eyJuYW1lIjoiZG9ub3JVc2VyMDEiLCJjb250cmFjdG5hbWUiOiLlroHlpI/opb/pg6jlnLDljLrmr43kurLmsLTnqpbpobnnm65YWOWOv1hY5p2R5rC056qWIiwiY29udHJhY3RhZGRyIjoiY29udHJhY3QwMTo4MDE3ZDFiZDE1MzI2YjQ4YjcwMjBmZjZmZTQzNGFlNDk5OWNmMDUwMDI1ZWM2NDg3MTUxNjY3N2YxMTc3NDYwIiwiYW1vdW50IjoxMDAwMDAwMDAsInRpbWVzdGFtcCI6MTQ4NjM2NDk3MH0=","L+3QF8+7qjMSgtWnQIwXh9cvqZNt/UBbBg5RUQ1eXFNgHy0VEWvHnK0112ojfJ5HrptF7jvw2JeFvWQgRotyHVz4gnlETbj6eESNpMamJIkpiyYnxM1VmnzkJz19GA3mka3Iq7qluiuNziXJUBqh1/Stl8Aw58X3MqmvzVEwq50="]
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
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"coinbase",
    "args":["cebbank:71266a9913375add341f1f7a4a606a2ee4e24efaf59bb2a3eb4b71aaaf5fb3fe","eyJ2ZXJzaW9uIjoxNzAxMDAxLCJ0aW1lc3RhbXAiOjE0ODUxNTcyOTIsInR4b3V0IjpbeyJ0eEhhc2giOiJ0eElESGFzaCIsInZhbHVlIjoxMDAwMDAwMDAwMDAwLCJhZGRyIjoiY2ViYmFuazo3MTI2NmE5OTEzMzc1YWRkMzQxZjFmN2E0YTYwNmEyZWU0ZTI0ZWZhZjU5YmIyYTNlYjRiNzFhYWFmNWZiM2ZlIiwic2lnbiI6ImIyOXVmUFlJU3MxQzdYWmMxbjhXcFdmaktUMUE0S1lJZ3JhRjdsbDBsM081NitZQThHdnk3WjdFbzZiSkUyOUZLRVZxUVBDaGY1RmVJOWg2QjFGdG9nbktvMFBVaDdMaUhaZnRsOVdlVjJCMVJMcmxwMGRlTlVvMU5TRUw3Qk03bVdLYi85ZjVQVlJ2WnVYY2lQaEdQVFpZaCtYWkdiT3JNMnJ4N0ZHYVorUT0ifV0sImZvdW5kZXIiOiJlcmljIn0=","jT9HO1vO3blGUSZ1tE1usKmzgdJhHl6Ia9moGpfYEvhSx8gMkclzU19AAv1/6k1g2BhLJw4MDqVMj8MBzm8NgHjOcD5oD0/+hOrAy+nIdOw9CZeSAnu5lOBKgb+XRqqq/WBJRFhpgwaJQAkYwNXQjtz56D7E2oedX3rtYxTPWwY="]
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

### changeCoin cebbank -> user01

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
    "type": 1,
    "chaincodeID":{
    "name":"3e0ff3192ad6b96324567aac390244cb585bc794fc38f6eaf921a2e4dc1c26fa03b4d87152208619fa2fe8223423a0a353a65704e4e4379d371b79275dbb5db5"
    },
    "ctorMsg": {
    "function":"changeCoin",
    "args":["cebbank:71266a9913375add341f1f7a4a606a2ee4e24efaf59bb2a3eb4b71aaaf5fb3fe","eyJ2ZXJzaW9uIjoxNzAxMDAxLCJ0aW1lc3RhbXAiOjE0ODUyNDgwMTQsInR4aW4iOlt7InNvdXJjZVR4SGFzaCI6ImViMTg3Y2JkZmQyZWQzMmRkOTAyYWJmN2VlOTk0MGY1ODFmZTFlYTg2ZTdlMmNkY2IwZTE2M2E5ZDU4NGE5MWMiLCJhZGRyIjoiY2ViYmFuazo3MTI2NmE5OTEzMzc1YWRkMzQxZjFmN2E0YTYwNmEyZWU0ZTI0ZWZhZjU5YmIyYTNlYjRiNzFhYWFmNWZiM2ZlIn1dLCJ0eG91dCI6W3sidmFsdWUiOjk5OTkwMDAwMDAwMCwiYWRkciI6ImNlYmJhbms6NzEyNjZhOTkxMzM3NWFkZDM0MWYxZjdhNGE2MDZhMmVlNGUyNGVmYWY1OWJiMmEzZWI0YjcxYWFhZjVmYjNmZSIsInNpZ24iOiJicVBLbFhpUlh3MVJobWQrU281Ti9DU002V3JEbzV6a2VkTUMwYTFwcDhzZSthbVRRMDUreTljd3d3NmJMYzdPSG0xRGRSdWZLcUViY0dtSGpBUkpzWi9Vakx6SlhNaUQvd0VYK1U1bHdjQnMrNU9GVXFXc0lEZ2NKZHNiWklLRkJDc1pCYmVDejhyaTFidE5pWDIrbXhMUkI5bkFscVJreWVBZG0zL1cwU3M9In0seyJ2YWx1ZSI6MTAwMDAwMDAwLCJhZGRyIjoidXNlcjAxOmU5NjI5ODIyMjA0YjllMzI5YTc2N2Y1YzRiZjM5OWI3N2ZhMGRmYTYzYzlkZjdiODhiMDQwYzUwMDNjMmFmM2EiLCJzaWduIjoiTUdYbjlDRFBKbnpWODAva1hYVkVRUG5zcDJUQmxySHFzRnZtN0JTWlRPbHU1SnRuUHE0Uno0RS9GbkNHR3JwK2prZ2QwemZBMWVtU3lnWTQvYmplNlRBQjM1UFd4R2N5REJOQW9lK1RKVzMxTUZJdU5wd216VGRmMXM0cXhHUUhDdFU3RHFMN0ZabU14Z3JKMy9ENStoeFpWWElwMlBVY0VyTW1leG8wczBzPSJ9XSwiZm91bmRlciI6ImVyaWMifQ==","Gp5oTUlNVE14nHlKm0Sbto1HmZdyWZVljSMI3urNLXPatkpXjQKk5RxiNl8TSPpiS/LAecB4OBQegFeQ3HeqhBZsfxPUm3AZJm7AbdDlTYLarv2qiKOyKjSouU2ckZ1bKuEc7ETZfhPCnVHtqEq03OWEtDhSzIn4OmnE1nlBd+Y="]
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
        "message": "{\"addr\":\"cebbank:71266a9913375add341f1f7a4a606a2ee4e24efaf59bb2a3eb4b71aaaf5fb3fe\",\"balance\":999900000000,\"rsaPublicKey\":\"LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDK0oxWjRiTGFPOFNobTFyOXFOOHB6RjJxbwp1c08rSnhoOWpGcDdQTkcwMERGTUR5RUhWNU9JajhhdFlVdzBwRkFTUG95dldHMGhlRzdCVU00bVRpWHNZOFF3ClpYMjZ6L2I2bk05Q3ZtM2xlell1NjgwT2NQQnNIczdLb0RZTkUrUGM2c0EwRGVDVnUxNm5aeGpQbFRwbUxMdmkKb2Z3bS9ReVlKZlVsNElRazFRSURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=\",\"txouts\":{\"ccffdf9eff747968ea7b87fca70ed18735abb746a4b825524e03c8c57d55cb8c:0\":{\"value\":999900000000,\"addr\":\"cebbank:71266a9913375add341f1f7a4a606a2ee4e24efaf59bb2a3eb4b71aaaf5fb3fe\",\"sign\":\"bqPKlXiRXw1Rhmd+So5N/CSM6WrDo5zkedMC0a1pp8se+amTQ05+y9cwww6bLc7OHm1DdRufKqEbcGmHjARJsZ/UjLzJXMiD/wEX+U5lwcBs+5OFUqWsIDgcJdsbZIKFBCsZBbeCz8ri1btNiX2+mxLRB9nAlqRkyeAdm3/W0Ss=\"}}}"
      },
      "id": 5
    }  

query user01:

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"addr\":\"user01:e9629822204b9e329a767f5c4bf399b77fa0dfa63c9df7b88b040c5003c2af3a\",\"balance\":100000000,\"rsaPublicKey\":\"LS0tLS1CRUdJTiBwdWJsaWMga2V5LS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDcGxmMDlJNTl5SDYzRnJwSlhqMSs0WmhjZgpSOXJVVi82akhWWVAxelQ1aHlUZUwzNWQ4NVpwTzhOTDdWV2MyYVpHeDFjZ2dvRG8rRUkvMVJobDA4UHhSbFdKCi9XVGt1VzBLbDIxeGRGUXFnSnk1Q3JUeUc0SURmS2NTRkg3dG9tZzFqeVlHaHRFZ2g5TUI3RVVZNXd0M0hWZHYKWEdWQWlMbzViU3VOaVRVV2J3SURBUUFCCi0tLS0tRU5EIHB1YmxpYyBrZXktLS0tLQo=\",\"txouts\":{\"ccffdf9eff747968ea7b87fca70ed18735abb746a4b825524e03c8c57d55cb8c:1\":{\"value\":100000000,\"addr\":\"user01:e9629822204b9e329a767f5c4bf399b77fa0dfa63c9df7b88b040c5003c2af3a\",\"sign\":\"MGXn9CDPJnzV80/kXXVEQPnsp2TBlrHqsFvm7BSZTOlu5JtnPq4Rz4E/FnCGGrp+jkgd0zfA1emSygY4/bje6TAB35PWxGcyDBNAoe+TJW31MFIuNpwmzTdf1s4qxGQHCtU7DqL7FZmMxgrJ3/D5+hxZVXIp2PUcErMmexo0s0s=\"}}}"
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







           
