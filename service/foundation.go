package service

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/CebEcloudTime/charitycc/core/coin"
	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
)

// RegisterFund register fund
func RegisterFund(store store.Store, args []string) ([]byte, error) {
	_fundAddr := args[0]
	_fundPublicKey := args[1]

	_, err := InitAccount(store, _fundAddr, _fundPublicKey)
	if err != nil {
		return nil, err
	}

	_, err = InitFund(store, _fundAddr)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func InitFund(store store.Store, addr string) ([]byte, error) {
	var newFoundation protos.Foundation
	newFoundation.Addr = addr

	if tmpFoundation, err := store.GetFund(newFoundation.Addr); err == nil && tmpFoundation != nil && tmpFoundation.Addr == newFoundation.Addr {
		return nil, errors.AlreadyRegisterd
	}

	if err := store.PutFund(&newFoundation); err != nil {
		return nil, err
	}

	return nil, nil

}

// QueryFund get a fund
func QueryFund(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	fund, err := store.GetFund(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(fund)
}

// RegisterSmartContract register smartContract
func RegisterSmartContract(store store.Store, args []string) ([]byte, error) {

	_smartContractAddr := args[1]
	_smartContractPublickKey := args[2]
	_base64SmartContractData := args[3]

	_, err := InitAccount(store, _smartContractAddr, _smartContractPublickKey)
	if err != nil {
		return nil, err
	}

	_, newSmartContract, err := InitSmartContract(store, _base64SmartContractData)
	if err != nil {
		return nil, err
	}

	_, err = InitSmartContractExt(store, _smartContractAddr, newSmartContract)
	if err != nil {
		return nil, err
	}

	_, err = InitSmartContractTrack(store, _smartContractAddr)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func InitSmartContract(store store.Store, base64SmartContractData string) ([]byte, *protos.SmartContract, error) {

	smartContractData, err := base64.StdEncoding.DecodeString(base64SmartContractData)
	if err != nil {
		return nil, nil, errors.Base64Decoding
	}

	newSmartContract, err := utils.ParseSmartContractByJsonBytes(smartContractData)
	if err != nil {
		return nil, nil, err
	}

	if tmpSmartContract, err := store.GetSmartContract(newSmartContract.Addr); err == nil && tmpSmartContract != nil && tmpSmartContract.Addr == newSmartContract.Addr {
		return nil, nil, errors.AlreadyRegisterd
	}

	if err := store.PutSmartContract(newSmartContract); err != nil {
		return nil, nil, err
	}

	return nil, newSmartContract, nil
}

// InitSmartContractExt
func InitSmartContractExt(store store.Store, addr string, smartContract *protos.SmartContract) ([]byte, error) {
	var newSmartContractExt protos.SmartContractExt
	newSmartContractExt.Addr = addr
	newSmartContractExt.SmartContract = smartContract

	if err := store.PutSmartContractExt(&newSmartContractExt); err != nil {
		return nil, err
	}

	return nil, nil
}

// InitSmartContractTrack
func InitSmartContractTrack(store store.Store, addr string) ([]byte, error) {
	var newSmartContractTrack protos.SmartContractTrack
	newSmartContractTrack.Addr = addr

	if tmpSmartContractTrack, err := store.GetSmartContractTrack(newSmartContractTrack.Addr); err == nil && tmpSmartContractTrack != nil && tmpSmartContractTrack.Addr == newSmartContractTrack.Addr {
		return nil, errors.AlreadyRegisterd
	}

	if err := store.PutSmartContractTrack(&newSmartContractTrack); err != nil {
		return nil, err
	}

	return nil, nil
}

// ChangeSmartContractStatus change foundation SmartContract status
func ChangeSmartContractStatus(store store.Store, args []string) ([]byte, error) {

	_smartContractAddr := args[1]
	_smartContractStatus := args[2]

	smartContract, err := store.GetSmartContract(_smartContractAddr)
	if err != nil {
		return nil, err
	}

	_status, err := strconv.ParseUint(_smartContractStatus, 10, 64)
	if err != nil {
		return nil, err
	}

	smartContract.Status = _status

	if err := store.PutSmartContract(smartContract); err != nil {
		return nil, err
	}

	return nil, nil
}

// QuerySmartContract get a smartContract
func QuerySmartContract(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	smartContract, err := store.GetSmartContract(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(smartContract)
}

// QuerySmartContractObj
func QuerySmartContractObj(store store.Store, addr string) (*protos.SmartContract, error) {

	smartContract, err := store.GetSmartContract(addr)
	if err != nil {
		return nil, err
	}

	return smartContract, nil
}

// QuerySmartContracts get smartContracts by addrs
func QuerySmartContracts(store store.Store, args []string) ([]byte, error) {

	addrs := args[0]

	var smartContracts []protos.SmartContract

	_addrs := strings.Split(addrs, ",")
	for _, _addr := range _addrs {
		_smartContract, err := store.GetSmartContract(_addr)
		if err != nil {
			continue
		}
		smartContracts = append(smartContracts, *_smartContract)
	}

	return json.Marshal(smartContracts)
}

// QuerySmartContractExt
func QuerySmartContractExt(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	smartContractExt, err := store.GetSmartContractExt(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(smartContractExt)
}

// QuerySmartContractExts
func QuerySmartContractExts(store store.Store, args []string) ([]byte, error) {

	addrs := args[0]

	var smartContractExts []protos.SmartContractExt

	_addrs := strings.Split(addrs, ",")
	for _, _addr := range _addrs {
		_smartContractExt, err := store.GetSmartContractExt(_addr)
		if err != nil {
			continue
		}
		smartContractExts = append(smartContractExts, *_smartContractExt)
	}

	return json.Marshal(smartContractExts)
}

// QuerySmartContractTrack
func QuerySmartContractTrack(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	smartContract, err := store.GetSmartContractTrack(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(smartContract)
}

// RegisterBargain register bargain
func RegisterBargain(store store.Store, args []string) ([]byte, error) {

	_bargainAddr := args[1]
	_bargainPublickKey := args[2]
	_base64BargainData := args[3]

	_, err := InitAccount(store, _bargainAddr, _bargainPublickKey)
	if err != nil {
		return nil, err
	}

	_, err = InitBargain(store, _base64BargainData)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func InitBargain(store store.Store, base64BargainData string) ([]byte, error) {

	bargainData, err := base64.StdEncoding.DecodeString(base64BargainData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newBargain, err := utils.ParseBargainByJsonBytes(bargainData)
	if err != nil {
		return nil, err
	}

	if tmpBargain, err := store.GetBargain(newBargain.Addr); err == nil && tmpBargain != nil && tmpBargain.Addr == newBargain.Addr {
		return nil, errors.AlreadyRegisterd
	}

	if err := store.PutBargain(newBargain); err != nil {
		return nil, err
	}

	return nil, nil
}

// QueryBargain get a bargain
func QueryBargain(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	bargain, err := store.GetBargain(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(bargain)
}

// QueryBargains get bargains by addrs
func QueryBargains(store store.Store, args []string) ([]byte, error) {

	addrs := args[0]

	var bargains []protos.Bargain

	_addrs := strings.Split(addrs, ",")
	for _, _addr := range _addrs {
		_bargain, err := store.GetBargain(_addr)
		if err != nil {
			continue
		}
		bargains = append(bargains, *_bargain)
	}

	return json.Marshal(bargains)
}

// QueryBargainTrack get a BargainTrack
func QueryBargainTrack(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	bargainTrack, err := store.GetBargainTrack(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(bargainTrack)
}

// Drawed foundation drawing
func Drawed(store store.Store, args []string) ([]byte, error) {

	newTX, err := DrawedTx(store, args)
	if err != nil {
		return nil, err
	}
	return json.Marshal(newTX)
}

// Drawed foundation drawing
func DrawedTx(store store.Store, args []string) (*protos.TX, error) {

	_base64TxData := args[1]

	txData, err := base64.StdEncoding.DecodeString(_base64TxData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newTX, err := utils.ParseTxByJsonBytes(txData)
	if err != nil {
		return nil, err
	}

	if newTX.Founder == "" {
		return nil, errors.TxNoFounder
	}

	utxo := coin.MakeUTXO(store)

	_, err = utxo.Execute(newTX)
	if err != nil {
		return nil, err
	}

	return newTX, nil
}

// GenDrawTxData gen draw tx
func GenDrawTxData(treatyAddr, contractAddr string, amount uint64) (*protos.TX, error) {

	// get treaty account, check if treaty account < amount

	// traverse account's txout, make txs(in: treaty; out: contract)...

	// util , txout > amount, make tx(in: treaty; out: contract,treaty)

	// save tx1,tx2,tx3... and save tx1,tx2,tx3... info into donor info'addr

	return nil, nil
}

// GenTrackByTxData
func GenTracksByTxData(tx protos.TX) ([]protos.DonorTrack, error) {

	return nil, nil
}
