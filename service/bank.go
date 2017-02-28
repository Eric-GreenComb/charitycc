package service

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"

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

// BuyCoin user change coin with CNY 1yuan = 1000000 coin
func BuyCoin(store store.Store, args []string) ([]byte, error) {

	base64TxData := args[1]
	txData, err := base64.StdEncoding.DecodeString(base64TxData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	sourceTX, err := utils.ParseTxByJsonBytes(txData)
	if err != nil {
		return nil, err
	}

	if sourceTX.Founder == "" {
		return nil, errors.TxNoFounder
	}

	buyCoinTX, err := GenBuyCoinTxData(store, *sourceTX)
	if err != nil {
		return nil, err
	}

	utxo := coin.MakeUTXO(store)

	_, err = utxo.Execute(buyCoinTX)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func GenBuyCoinTxData(store store.Store, sourceTX protos.TX) (*protos.TX, error) {
	bankAddr := sourceTX.Txin[0].Addr
	donorAddr := sourceTX.Txout[0].Addr
	amount := sourceTX.Txout[0].Value
	attr := sourceTX.Txout[0].Attr

	bankAccount, err := store.GetAccount(bankAddr)
	if err != nil {
		return nil, err
	}

	if bankAccount.Balance < amount {
		return nil, errors.AccountNotEnoughBalance
	}

	var tx protos.TX

	tx.Version = 170101
	tx.Timestamp = time.Now().UTC().Unix()

	txins, txouts, err := GenBuyCoinTxInOutData(bankAccount, bankAddr, donorAddr, attr, amount)
	if err != nil {
		return nil, err
	}

	tx.Txin = txins

	tx.Txout = txouts

	tx.InputData = sourceTX.InputData

	tx.Founder = sourceTX.Founder

	return &tx, nil
}

func GenBuyCoinTxInOutData(account *protos.Account, bankAddr, donorAddr, attr string, amount uint64) ([]*protos.TX_TXIN, []*protos.TX_TXOUT, error) {
	var txins []*protos.TX_TXIN
	var txouts []*protos.TX_TXOUT

	amountTemp := amount

	for k, v := range account.GetTxouts() {
		if amountTemp == 0 {
			break
		}

		var txin protos.TX_TXIN
		txin.Addr = bankAddr
		_keySplit := strings.Split(k, ":")
		txin.SourceTxHash = _keySplit[0]
		_nIdx, _ := strconv.ParseUint(_keySplit[1], 10, 64)
		txin.Idx = uint32(_nIdx)
		txins = append(txins, &txin)

		if v.Value <= amountTemp {

			var txout protos.TX_TXOUT

			txout.Value = v.Value
			txout.Addr = donorAddr
			txout.Attr = attr

			txouts = append(txouts, &txout)

			amountTemp -= v.Value

		} else {

			var txout protos.TX_TXOUT

			txout.Value = amountTemp
			txout.Addr = donorAddr
			txout.Attr = attr
			txouts = append(txouts, &txout)

			var txoutRe protos.TX_TXOUT

			txoutRe.Value = v.Value - amountTemp
			txoutRe.Addr = bankAddr
			txoutRe.Attr = "init,0"
			txouts = append(txouts, &txoutRe)

			amountTemp = 0
		}
	}

	return txins, txouts, nil
}
