package main

import (
	"errors"
	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/handler"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *CharityCC) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// construct a new store
	storeCC := store.MakeCCStore(stub)

	switch function {

	case "registerBank":
		// args: bankID,publicKey(bank base64),sysadminSign(base64)
		return handler.RegisterBank(storeCC, args)

	case "coinbase":
		// args: addr, txData(coinbase tx Base64),bankSign(base64)
		return handler.Coinbase(storeCC, args)
	case "destroycoinbase":
		// args: targetAddr, bankAddr, bankSign(base64)
		return handler.DestoryCoinbase(storeCC, args)
	case "changeCoin":
		// args: sourceAddr, txData(coinbase tx Base64),sourcSign(base64)
		return handler.ChangeCoin(storeCC, args)

	case "registerAccount":
		// args: ID,publicKey(ID base64),sysadminSign(base64)
		return handler.RegisterAccount(storeCC, args)

	case "doDonate":
		// args: business(Business id),donorID,publicKey(donor base64),channelID,publicKey(channel base64),treatyid,value,sign(base64)
		return handler.DoDonate(storeCC, args)
	case "doDrawing":
		// args: business(Business id),treatyID,publicKey(treaty base64),contractid,value,sign(base64)
		return handler.DoDrawing(storeCC, args)

	case "openDonorAccount":
		// args: donorID,publicKey(donor base64),userSign(base64)
		return handler.OpenDonorAccount(storeCC, args)

	default:
		return nil, errors.New("Unsupported operation")
	}

	return nil, nil
}
