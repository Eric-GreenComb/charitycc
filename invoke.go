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
		// args: addr,txData(coinbase tx Base64),bankSign(base64)
		return handler.Coinbase(storeCC, args)
	case "changeCoin":
		// args: bankID,publicKey(bank base64),amount(change amount), txData(coinbase tx Base64),bankSign(base64)
		return handler.ChangeCoin(storeCC, args)

	case "registerChannelAccount":
		// args: channelID,publicKey(channel base64),channelSign(base64)
		return handler.RegisterChannelAccount(storeCC, args)

	case "registerFoundationAccount":
		// args: foundationID,publicKey(foundation base64),foundationSign(base64)
		return handler.RegisterFoundationAccount(storeCC, args)
	case "registerTreatyAccount":
		// args: treatyID,publicKey(treaty base64),treatySign(base64)
		return handler.RegisterTreatyAccount(storeCC, args)

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
