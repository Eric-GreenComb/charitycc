package coin

import (
	baseerrors "errors"
	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
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

		if outerAccount.Txouts == nil || len(outerAccount.Txouts) == 0 {
			outerAccount.Txouts = make(map[string]*protos.TX_TXOUT)
		}

		currKey := &Key{TxHashAsHex: txhash, TxIndex: uint32(index)}
		if _, ok := outerAccount.Txouts[currKey.String()]; ok {
			return nil, errors.CollisionTxOut
		}

		outerAccount.Txouts[currKey.String()] = output
		outerAccount.Balance += output.Value
		if err := u.store.PutAccount(outerAccount); err != nil {
			return nil, err
		}

		execResult.SumCurrentOutputs += output.Value
	}

	if err := u.store.PutTran(newTX); err != nil {
		return nil, err
	}

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
		txout, ok := ownerAccount.Txouts[keyToPrevOutput.String()]
		if !ok {
			return nil, baseerrors.New(keyToPrevOutput.String())
		}

		if ownerAccount.Balance < txout.Value {
			return nil, errors.AccountNotEnoughBalance
		}

		ownerAccount.Balance -= txout.Value
		delete(ownerAccount.Txouts, keyToPrevOutput.String())
		// save owner account
		if err := u.store.PutAccount(ownerAccount); err != nil {
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

		if outerAccount.Txouts == nil || len(outerAccount.Txouts) == 0 {
			outerAccount.Txouts = make(map[string]*protos.TX_TXOUT)
		}

		currKey := &Key{TxHashAsHex: txhash, TxIndex: uint32(index)}
		if _, ok := outerAccount.Txouts[currKey.String()]; ok {
			return nil, errors.CollisionTxOut
		}

		outerAccount.Txouts[currKey.String()] = output
		outerAccount.Balance += output.Value
		if err := u.store.PutAccount(outerAccount); err != nil {
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

	if err := u.store.PutTran(newTX); err != nil {
		return nil, err
	}

	return execResult, nil
}
