package main

import (
	"errors"
	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/handler"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *CharityCC) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// construct a new store
	storeCC := store.MakeCCStore(stub)

	switch function {

	case "queryBank":
		// args: addr(%s:%s id, publickey -> sha256 -> sha256 -> hex)
		return handler.QueryBank(storeCC, args)

	case "queryAccount":
		// args: addr(%s:%s id, publickey -> sha256 -> sha256 -> hex)
		return handler.QueryAccount(storeCC, args)

	case "querySmartContract":
		// args: addr(client gen addr)
		return handler.QuerySmartContract(storeCC, args)
	case "querySmartContracts":
		// args: addrs(client gen addrs, split by ,)
		return handler.QuerySmartContracts(storeCC, args)
	case "queryBargain":
		// args: addr(client gen addr)
		return handler.QueryBargain(storeCC, args)
	case "queryBargains":
		// args: addrs(client gen addrs, split by ,)
		return handler.QueryBargains(storeCC, args)

	case "queryDonor":
		// args: addr(client gen addr)
		return handler.QueryDonor(storeCC, args)

	case "queryDonorDetail":
		// big

	case "queryTransDetail":
		// donor,draw small

	case "querySmartContractTrack":
		//

	case "queryFund":
		//

	default:
		return nil, errors.New("Unsupported operation")
	}

	return nil, nil
}
