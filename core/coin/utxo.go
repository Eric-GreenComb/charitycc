package coin

import (
	"github.com/ecloudtime/charitycc/core/store"
	"github.com/ecloudtime/charitycc/errors"
	"github.com/ecloudtime/charitycc/protos"
	"github.com/ecloudtime/charitycc/utils"
)

// UTXO includes the storage for the chaincode API
type UTXO struct {
	store store.Store
}

// MakeUTXO constructs a new UTXO with the given store
func MakeUTXO(store store.Store) *UTXO {
	utxo := &UTXO{}
	utxo.store = store
	return utxo
}

// ExecResult is the result of processing a transaction
type ExecResult struct {
	SumCurrentOutputs uint64
	SumPriorOutputs   uint64
	TxHash            string
}

// Coinbase processes the given coinbase transaction and outputs a result
func (u *UTXO) Coinbase(newTX *protos.TX) (*ExecResult, error) {

	txhash := utils.GenTransactionHash(newTX)

	execResult := &ExecResult{}
	execResult.TxHash = txhash

	// Loop through outputs first
	for index, output := range newTX.Txout {

		outerAccount, err := u.store.GetAccount(output.Addr)
		if err != nil {
			return nil, errors.InvalidAccount
		}

		currKey := &Key{TxHashAsHex: txhash, TxIndex: uint32(index)}

		outerAccount.CoinKey = append(outerAccount.CoinKey, currKey.String())

		outerAccount.Balance += output.Value
		if err := u.store.PutAccount(outerAccount); err != nil {
			return nil, err
		}

		if err := u.store.PutCoin(currKey.String(), output); err != nil {
			return nil, err
		}

		execResult.SumCurrentOutputs += output.Value
	}

	// if err := u.store.PutTran(newTX); err != nil {
	// 	return nil, err
	// }

	return execResult, nil
}

// Execute processes the given transaction and outputs a result
func (u *UTXO) Execute(newTX *protos.TX) (*ExecResult, error) {

	txhash := utils.GenTransactionHash(newTX)

	execResult := &ExecResult{}
	execResult.TxHash = txhash

	// Now loop over inputs,
	for _, ti := range newTX.Txin {
		prevTxHash := ti.SourceTxHash
		prevOutputIx := ti.Idx
		ownerAddr := ti.Addr
		keyToPrevOutput := &Key{TxHashAsHex: prevTxHash, TxIndex: prevOutputIx}

		ownerAccount, err := u.store.GetAccount(ownerAddr)
		if err != nil {
			return nil, err
		}

		txout, err := u.store.GetCoin(keyToPrevOutput.String())
		if err != nil {
			return nil, err
		}

		ownerAccount.Balance -= txout.Value
		ownerAccount.CoinKey = DeleteArray(ownerAccount.CoinKey, keyToPrevOutput.String())

		// save owner account
		if err := u.store.PutAccount(ownerAccount); err != nil {
			return nil, err
		}

		if err := u.store.DeleteCoin(keyToPrevOutput.String()); err != nil {
			return nil, err
		}

		execResult.SumPriorOutputs += txout.Value
	}

	// Loop through outputs
	for index, output := range newTX.Txout {

		outerAccount, err := u.store.GetAccount(output.Addr)
		if err != nil {
			return nil, errors.InvalidAccount
		}

		currKey := &Key{TxHashAsHex: txhash, TxIndex: uint32(index)}

		outerAccount.Balance += output.Value
		outerAccount.CoinKey = append(outerAccount.CoinKey, currKey.String())

		if err := u.store.PutAccount(outerAccount); err != nil {
			return nil, err
		}

		if err := u.store.PutCoin(currKey.String(), output); err != nil {
			return nil, err
		}

		execResult.SumCurrentOutputs += output.Value
	}

	// one of transfer main point is in == out, no coin mined, no coin lose
	if execResult.SumCurrentOutputs != execResult.SumPriorOutputs {
		return nil, errors.TxInOutNotBalance
	}

	if execResult.SumCurrentOutputs > execResult.SumPriorOutputs {
		return nil, errors.TxOutMoreThanTxIn
	}

	// if err := u.store.PutTran(newTX); err != nil {
	// 	return nil, err
	// }

	return execResult, nil
}

func DeleteArray(array []string, key string) []string {
	index := 0
	endIndex := len(array) - 1

	var result = make([]string, 0)
	for k, v := range array {
		if v == key {
			result = append(result, array[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, array[index:endIndex+1]...)
		}
	}

	return result
}
