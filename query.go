package main

import (
	"errors"
	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/handler"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Query chaincode query
func (t *CharityCC) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// construct a new store
	storeCC := store.MakeCCStore(stub)

	switch function {

	case "queryBank":
		// args: addr(%s:%s id, publickey -> sha256 -> sha256 -> hex)
		return handler.QueryBank(storeCC, args)

	case "queryFund":
		// args: addr(%s:%s id, publickey -> sha256 -> sha256 -> hex)
		return handler.QueryFund(storeCC, args)

	case "queryChannel":
		// args: addr(%s:%s id, publickey -> sha256 -> sha256 -> hex)
		return handler.QueryChannel(storeCC, args)

	case "queryAccount":
		// args: addr(%s:%s id, publickey -> sha256 -> sha256 -> hex)
		return handler.QueryAccount(storeCC, args)

	case "querySmartContract":
		// args: addr(client gen addr)
		return handler.QuerySmartContract(storeCC, args)
	case "querySmartContractExt":
		// args: addrs(client gen addrs) return: SmartContractExt List
		return handler.QuerySmartContractExt(storeCC, args)
	case "querySmartContractExts":
		// args: addrs(client gen addrs, split by ,) return: SmartContractExt List
		return handler.QuerySmartContractExts(storeCC, args)
	case "querySmartContractTrack":
		// args: addrs(client gen addrs, split by ,)
		return handler.QuerySmartContractTrack(storeCC, args)

	case "queryBargain":
		// args: addr(client gen addr)
		return handler.QueryBargain(storeCC, args)
	case "queryBargains":
		// args: addrs(client gen addrs, split by ,)
		return handler.QueryBargains(storeCC, args)
	case "queryBargainTrack":
		// args: addrs(client gen addrs, split by ,)
		return handler.QueryBargainTrack(storeCC, args)

	case "queryDonor":
		// args: addr(client gen addr)
		return handler.QueryDonor(storeCC, args)

	case "queryDonorDetail":
		//捐款详情

	case "queryTransDetail":
		//捐款或提款的交易详情

	case "queryProcessDonored":
		// args: donorUUID
		return handler.QueryProcessDonored(storeCC, args)

	default:
		return nil, errors.New("Unsupported operation")
	}

	return nil, nil
}
