package main

import (
	"errors"

	"github.com/ecloudtime/charitycc/core/store"
	"github.com/ecloudtime/charitycc/handler"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Query chaincode query
func (t *CharityCC) Query(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	// construct a new store
	storeCC := store.MakeCCStore(stub)

	var Avalbytes []byte
	var err error

	switch function {

	case "queryFund":
		// args: addr(%s:%s id, publickey -> sha256 -> sha256 -> hex)
		Avalbytes, err = handler.QueryFund(storeCC, args)

	case "queryChannel":
		// args: addr(%s:%s id, publickey -> sha256 -> sha256 -> hex)
		Avalbytes, err = handler.QueryChannel(storeCC, args)

	case "queryAccount":
		// args: addr(%s:%s id, publickey -> sha256 -> sha256 -> hex)
		Avalbytes, err = handler.QueryAccount(storeCC, args)
	case "queryCoin":
		// args: addr
		Avalbytes, err = handler.QueryCoin(storeCC, args)

	case "querySmartContract":
		// args: addr(client gen addr)
		Avalbytes, err = handler.QuerySmartContract(storeCC, args)
	case "querySmartContractExt":
		// args: addrs(client gen addrs) return: SmartContractExt List
		Avalbytes, err = handler.QuerySmartContractExt(storeCC, args)
	case "querySmartContractExts":
		// args: addrs(client gen addrs, split by ",") return: SmartContractExt List
		Avalbytes, err = handler.QuerySmartContractExts(storeCC, args)
	case "querySmartContractTrack":
		// args: addrs(client gen addrs)
		Avalbytes, err = handler.QuerySmartContractTrack(storeCC, args)

	case "queryBargain":
		// args: addr(client gen addr)
		Avalbytes, err = handler.QueryBargain(storeCC, args)
	case "queryBargains":
		// args: addrs(client gen addrs, split by ,)
		Avalbytes, err = handler.QueryBargains(storeCC, args)
	case "queryBargainTrack":
		// args: addrs(client gen addrs)
		Avalbytes, err = handler.QueryBargainTrack(storeCC, args)

	case "queryDonor":
		// args: addr(client gen addr)
		Avalbytes, err = handler.QueryDonor(storeCC, args)

	case "queryProcessDonored":
		// args: donorUUID
		Avalbytes, err = handler.QueryProcessDonored(storeCC, args)

	case "queryProcessDrawed":
		// args: drawUUID
		Avalbytes, err = handler.QueryProcessDrawed(storeCC, args)

	default:
		Avalbytes, err = nil, errors.New("Unsupported operation")
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(Avalbytes)
}
