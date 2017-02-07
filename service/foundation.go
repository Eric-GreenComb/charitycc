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

// RegisterTreaty register treaty
func RegisterTreaty(store store.Store, args []string) ([]byte, error) {

	base64TreatyData := args[1]
	treatyData, err := base64.StdEncoding.DecodeString(base64TreatyData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newTreaty, err := utils.ParseTreatyByJsonBytes(treatyData)
	if err != nil {
		return nil, err
	}

	if tmpTreaty, err := store.GetTreaty(newTreaty.Addr); err == nil && tmpTreaty != nil && tmpTreaty.Addr == newTreaty.Addr {
		return nil, errors.AlreadyRegisterd
	}

	if err := store.PutTreaty(newTreaty); err != nil {
		return nil, err
	}

	return nil, nil
}

// ChangeTreatyStatus change foundation treaty status
func ChangeTreatyStatus(store store.Store, args []string) ([]byte, error) {

	_treatyAddr := args[1]
	_treatyStatus := args[2]

	treaty, err := store.GetTreaty(_treatyAddr)
	if err != nil {
		return nil, err
	}

	_status, err := strconv.ParseUint(_treatyStatus, 10, 64)
	if err != nil {
		return nil, err
	}

	treaty.Status = _status

	if err := store.PutTreaty(treaty); err != nil {
		return nil, err
	}

	return nil, nil
}

// QueryTreaty get a treaty
func QueryTreaty(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	treaty, err := store.GetTreaty(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(treaty)
}

// QueryTreatyObj
func QueryTreatyObj(store store.Store, addr string) (*protos.Treaty, error) {

	treaty, err := store.GetTreaty(addr)
	if err != nil {
		return nil, err
	}

	return treaty, nil
}

// QueryTreaties get treaties by addrs
func QueryTreaties(store store.Store, args []string) ([]byte, error) {

	addrs := args[0]

	var treaties []protos.Treaty

	_addrs := strings.Split(addrs, ",")
	for _, _addr := range _addrs {
		treaty, err := store.GetTreaty(_addr)
		if err != nil {
			continue
		}
		treaties = append(treaties, *treaty)
	}

	return json.Marshal(treaties)
}

// RegisterContract register contract
func RegisterContract(store store.Store, args []string) ([]byte, error) {

	base64ContractData := args[1]
	contractData, err := base64.StdEncoding.DecodeString(base64ContractData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newContract, err := utils.ParseContractByJsonBytes(contractData)
	if err != nil {
		return nil, err
	}

	if tmpContract, err := store.GetContract(newContract.Addr); err == nil && tmpContract != nil && tmpContract.Addr == newContract.Addr {
		return nil, errors.AlreadyRegisterd
	}

	if err := store.PutContract(newContract); err != nil {
		return nil, err
	}

	return nil, nil
}

// QueryContract get a contract
func QueryContract(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	contract, err := store.GetContract(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(contract)
}

// QueryContractObj
func QueryContractObj(store store.Store, addr string) (*protos.Contract, error) {

	contract, err := store.GetContract(addr)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

// QueryContracts get contracts by addrs
func QueryContracts(store store.Store, args []string) ([]byte, error) {

	addrs := args[0]

	var contracts []protos.Contract

	_addrs := strings.Split(addrs, ",")
	for _, _addr := range _addrs {
		_contract, err := store.GetContract(_addr)
		if err != nil {
			continue
		}
		contracts = append(contracts, *_contract)
	}

	return json.Marshal(contracts)
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
