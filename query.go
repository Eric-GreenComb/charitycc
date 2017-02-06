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

	case "queryTreaty":
		// args: addr(client gen addr)
		return handler.QueryTreaty(storeCC, args)
	case "queryTreaties":
		// args: addrs(client gen addrs, split by ,)
		return handler.QueryTreaties(storeCC, args)
	case "queryContract":
		// args: addr(client gen addr)
		return handler.QueryContract(storeCC, args)
	case "queryContracts":
		// args: addrs(client gen addrs, split by ,)
		return handler.QueryContracts(storeCC, args)

	case "queryDonor":
		// args: addr(client gen addr)
		return handler.QueryDonor(storeCC, args)

	default:
		return nil, errors.New("Unsupported operation")
	}

	return nil, nil
}
