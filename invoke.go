package main

import (
	"errors"

	"github.com/ecloudtime/charitycc/core/store"
	"github.com/ecloudtime/charitycc/handler"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Invoke chaincode invoke
func (t *CharityCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	// construct a new store
	storeCC := store.MakeCCStore(stub)

	var Avalbytes []byte
	var err error

	switch function {

	case "registerAccount":
		// args: addr,publicKey(ID base64),sysadminSign(base64)
		Avalbytes, err = handler.RegisterAccount(storeCC, args)

	case "coinbase":
		// args: addr, txData(coinbase tx Base64),bankSign(base64)
		Avalbytes, err = handler.Coinbase(storeCC, args)
	case "destroycoinbase":
		// args: targetAddr, bankAddr, bankSign(base64)
		Avalbytes, err = handler.DestoryCoinbase(storeCC, args)
	case "changeCoin":
		// args: sourceAddr, txData(coinbase tx Base64),sourcSign(base64)
		Avalbytes, err = handler.ChangeCoin(storeCC, args)
	case "buyCoin":
		// args: sourceAddr, txData(coinbase tx Base64),sourcSign(base64)
		Avalbytes, err = handler.BuyCoin(storeCC, args)

	case "registerBank":
		// args: addr,bankPublicKey(base64),sysadminSign(base64)
		Avalbytes, err = handler.RegisterBank(storeCC, args)

	case "registerFund":
		// args: fundAddr,fundPublicKey(base64),sysadminSign(base64)
		Avalbytes, err = handler.RegisterFund(storeCC, args)

	case "registerChannel":
		// args: channelAddr,channelPublicKey(base64),sysadminSign(base64)
		Avalbytes, err = handler.RegisterChannel(storeCC, args)

	case "registerDonor":
		// args: channelAddr,donorAddr,donorPublicKey(base64),channelSign(base64)
		Avalbytes, err = handler.RegisterDonor(storeCC, args)
	// case "addContribution":
	// 	// args: channelAddr,donorAddr,contributionData(contribution json base64),channelSign(base64)
	// 	return handler.AddContribution(storeCC, args)
	// case "addTrack":
	// 	// args: channelAddr,donorAddr,trackData(track json base64),channelSign(base64)
	// 	return handler.AddTrack(storeCC, args)
	case "donated":
		// args: donorAddr,donorAddr,donorUUID,smartContractAddr,txData(coinbase tx Base64),donorSign(base64),extInfo
		Avalbytes, err = handler.Donated(storeCC, args)
	// case "doDonating":
	// 	// args: donorUUID, donorAddr,smartContractAddr,amount,bankAddr,channelAddr, fundAddr,donorSign(base64)
	// 	return handler.DoDonating(storeCC, args)

	case "registerSmartContract":
		// args: foundationAddr, smartContractAddr, smartContractPublickKey(base64), smartContractData(treaty json base64),foundationSign(base64)
		Avalbytes, err = handler.RegisterSmartContract(storeCC, args)
	case "changeSmartContractStatus":
		// args: foundationAddr, addr(smartContract addr), status,foundationSign(base64)
		Avalbytes, err = handler.ChangeSmartContractStatus(storeCC, args)

	case "registerBargain":
		// args: foundationAddr,bargainAddr, bargainPublickKey(base64), bargainData(contract json base64),foundationSign(base64)
		Avalbytes, err = handler.RegisterBargain(storeCC, args)

	case "drawed":
		// args: foundationAddr,drawUUID,txData(coinbase tx Base64),foundationSign(base64),extInfo
		Avalbytes, err = handler.Drawed(storeCC, args)
	// case "doDrawing":
	// 	// args: foundationAddr,txData(coinbase tx Base64),foundationSign(base64)
	// 	return handler.DoDrawing(storeCC, args)

	default:
		Avalbytes, err = nil, errors.New("Unsupported operation")
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(Avalbytes)
}
