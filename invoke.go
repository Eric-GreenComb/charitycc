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

	case "registerTreaty":
		// args: foundationAddr, treatyData(treaty json base64),foundationSign(base64)
		return handler.RegisterTreaty(storeCC, args)
	case "changeTreatyStatus":
		// args: foundationAddr, addr(treaty addr), status,foundationSign(base64)
		return handler.ChangeTreatyStatus(storeCC, args)
	case "registerContract":
		// args: foundationAddr, contractData(contract json base64),foundationSign(base64)
		return handler.RegisterContract(storeCC, args)

	case "registerDonor":
		// args: channelAddr,donorData(donor json base64),channelSign(base64)
		return handler.RegisterDonor(storeCC, args)
	case "addContribution":
		// args: channelAddr,donorAddr,contributionData(contribution json base64),channelSign(base64)
		return handler.AddContribution(storeCC, args)
	case "addTrack":
		// args: channelAddr,donorAddr,trackData(track json base64),channelSign(base64)
		return handler.AddTrack(storeCC, args)
	case "donated":
		// args: channelAddr,donorAddr,txData(coinbase tx Base64),channelSign(base64)
		return handler.Donated(storeCC, args)
	case "doDonating":
		// args: ...
		return handler.DoDonating(storeCC, args)

	case "drawed":
		// args: foundationAddr,txData(coinbase tx Base64),foundationSign(base64)
		return handler.Drawed(storeCC, args)
	case "doDrawing":
		// args: foundationAddr,txData(coinbase tx Base64),foundationSign(base64)
		return handler.DoDrawing(storeCC, args)

	default:
		return nil, errors.New("Unsupported operation")
	}

	return nil, nil
}
