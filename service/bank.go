package service

import (
	"encoding/base64"
	"encoding/json"

	"github.com/CebEcloudTime/charitycc/core/coin"
	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
)

// RegisterBank register bank account
func RegisterBank(store store.Store, args []string) ([]byte, error) {

	bankID := args[0]
	publicKey := args[1]

	addr := utils.GenAddr(bankID, publicKey)

	if tmpaccount, err := store.GetAccount(addr); err == nil && tmpaccount != nil && tmpaccount.Addr == addr {
		return nil, errors.AlreadyRegisterd
	}

	account := &protos.Account{
		Addr:         addr,
		Balance:      0,
		RsaPublicKey: publicKey,
		Txouts:       make(map[string]*protos.TX_TXOUT),
	}
	if err := store.PutAccount(account); err != nil {
		return nil, err
	}

	return nil, nil
}

// GetBank get bank account
func QueryBank(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	account, err := store.GetAccount(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(account)
}

// GetBank get bank account
func Coinbase(store store.Store, args []string) ([]byte, error) {

	base64TxData := args[2]
	txData, err := base64.StdEncoding.DecodeString(base64TxData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newTX, err := utils.ParseTxByBytes(txData)
	if err != nil {
		return nil, err
	}

	if newTX.Founder == "" {
		return nil, errors.TxNoFounder
	}

	utxo := coin.MakeUTXO(store)

	execResult, err := utxo.Execute(true, newTX)
	if err != nil {
		return nil, err
	}

	return execResult.ObjResult, nil
}

// ChangeCoin get bank account
func ChangeCoin(store store.Store, args []string) ([]byte, error) {

	base64TxData := args[2]
	txData, err := base64.StdEncoding.DecodeString(base64TxData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newTX, err := utils.ParseTxByBytes(txData)
	if err != nil {
		return nil, err
	}

	if newTX.Founder == "" {
		return nil, errors.TxNoFounder
	}

	utxo := coin.MakeUTXO(store)

	execResult, err := utxo.Execute(false, newTX)
	if err != nil {
		return nil, err
	}

	return execResult.ObjResult, nil

	return nil, nil
}
