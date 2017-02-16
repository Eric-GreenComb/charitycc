package service

import (
	"encoding/base64"

	"github.com/CebEcloudTime/charitycc/core/coin"
	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
)

// RegisterBank register bank account
func RegisterBank(store store.Store, args []string) ([]byte, error) {

	addr := args[0]
	publicKey := args[1]

	return InitAccount(store, addr, publicKey)
}

// GetBank get bank account
func Coinbase(store store.Store, args []string) ([]byte, error) {

	base64TxData := args[1]
	txData, err := base64.StdEncoding.DecodeString(base64TxData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newTX, err := utils.ParseTxByJsonBytes(txData)
	if err != nil {
		return nil, errors.InvalidTX
	}

	if newTX.Founder == "" {
		return nil, errors.TxNoFounder
	}

	utxo := coin.MakeUTXO(store)

	_, err = utxo.Coinbase(newTX)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// DestoryCoinbase destory addr coinbase
func DestoryCoinbase(store store.Store, args []string) ([]byte, error) {

	_targetAddr := args[0]

	_tmpAccount, err := store.GetAccount(_targetAddr)
	if err != nil {
		return nil, errors.InvalidAccount
	}

	account := &protos.Account{
		Addr:         _targetAddr,
		Balance:      0,
		RsaPublicKey: _tmpAccount.RsaPublicKey,
		Txouts:       make(map[string]*protos.TX_TXOUT),
	}
	if err := store.PutAccount(account); err != nil {
		return nil, err
	}

	return nil, nil
}

// ChangeCoin get bank account
func ChangeCoin(store store.Store, args []string) ([]byte, error) {

	base64TxData := args[1]
	txData, err := base64.StdEncoding.DecodeString(base64TxData)
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

	return nil, nil
}
