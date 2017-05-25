package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ecloudtime/charitycc/core/coin"
	"github.com/ecloudtime/charitycc/core/store"
	"github.com/ecloudtime/charitycc/errors"
	"github.com/ecloudtime/charitycc/protos"
	"github.com/ecloudtime/charitycc/utils"
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

	newSmartContract, err := utils.ParseSmartContractByJSONBytes(smartContractData)
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

	smartContractAccount, err := store.GetAccount(addr)
	if err != nil {
		return nil, err
	}
	smartContractExt.Balance = smartContractAccount.Balance

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

		smartContractAccount, err := store.GetAccount(_addr)
		if err != nil {
			continue
		}
		_smartContractExt.Balance = smartContractAccount.Balance

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

	_, err = InitBargainTrack(store, _bargainAddr)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func InitBargainTrack(store store.Store, bargainAddr string) ([]byte, error) {
	var bargainTrack protos.BargainTrack
	bargainTrack.Addr = bargainAddr

	if err := store.PutBargainTrack(&bargainTrack); err != nil {
		return nil, err
	}

	return nil, nil
}

func InitBargain(store store.Store, base64BargainData string) ([]byte, error) {

	bargainData, err := base64.StdEncoding.DecodeString(base64BargainData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newBargain, err := utils.ParseBargainByJSONBytes(bargainData)
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

	_drawUUID := args[1]
	_base64TxData := args[2]

	txData, err := base64.StdEncoding.DecodeString(_base64TxData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	sourceTX, err := utils.ParseTxByJSONBytes(txData)
	if err != nil {
		return nil, err
	}

	if sourceTX.Founder == "" {
		return nil, errors.TxNoFounder
	}

	smartContractAddr, bargainAddr, amount := GetDrawedInfo(*sourceTX)

	drawedTX, err := GenDrawedTxData(store, *sourceTX)
	if err != nil {
		return nil, err
	}

	// drawed
	utxo := coin.MakeUTXO(store)

	_, err = utxo.Execute(drawedTX)
	if err != nil {
		return nil, err
	}

	bargain, err := store.GetBargain(bargainAddr)
	if err != nil {
		return nil, err
	}

	SaveDonorTrack(store, _drawUUID, drawedTX, bargain)

	err = SaveProcessDrawed(store, _drawUUID, smartContractAddr, drawedTX, bargain, amount)
	if err != nil {
		return nil, err
	}

	err = SaveSmartContractTrack(store, smartContractAddr, bargain, amount, sourceTX.Timestamp)
	if err != nil {
		return nil, err
	}

	err = SaveBargainTrack(store, bargain, amount, sourceTX.Timestamp)
	if err != nil {
		return nil, err
	}

	err = SaveDrawedFundFee(store, smartContractAddr, amount)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func SaveDrawedFundFee(store store.Store, smartContractAddr string, amount uint64) error {

	_smartContract, err := store.GetSmartContract(smartContractAddr)
	if err != nil {
		return err
	}

	tmpFund, err := store.GetFund(_smartContract.FundAddr)
	if err != nil {
		return err
	}

	tmpFund.Balance -= amount

	if err := store.PutFund(tmpFund); err != nil {
		return err
	}

	return nil
}

func GetDrawedInfo(sourceTX protos.TX) (string, string, uint64) {
	smartContractAddr := sourceTX.Txin[0].Addr
	bargainAddr := sourceTX.Txout[0].Addr
	amount := sourceTX.Txout[0].Value

	return smartContractAddr, bargainAddr, amount
}

func GenDrawedTxData(store store.Store, sourceTX protos.TX) (*protos.TX, error) {
	smartContractAddr := sourceTX.Txin[0].Addr
	bargainAddr := sourceTX.Txout[0].Addr
	amount := sourceTX.Txout[0].Value

	smartContractAccount, err := store.GetAccount(smartContractAddr)
	if err != nil {
		return nil, err
	}

	if smartContractAccount.Balance < amount {
		return nil, fmt.Errorf("CC account not enough:  %d;%d", smartContractAccount.Balance, amount)
	}

	var tx protos.TX

	tx.Version = sourceTX.Version
	tx.Timestamp = sourceTX.Timestamp

	txins, txouts, err := GenDrawedTxInOutData(store, smartContractAccount, smartContractAddr, bargainAddr, amount)
	if err != nil {
		return nil, err
	}

	tx.Txin = txins

	tx.Txout = txouts

	tx.InputData = sourceTX.InputData

	tx.Founder = sourceTX.Founder

	return &tx, nil
}

func GenDrawedTxInOutData(store store.Store, account *protos.Account, smartContractAddr, bargainAddr string, amount uint64) ([]*protos.TX_TXIN, []*protos.TX_TXOUT, error) {
	var txins []*protos.TX_TXIN
	var txouts []*protos.TX_TXOUT

	amountTemp := amount

	for _, v := range account.CoinKey {
		if amountTemp == 0 {
			break
		}

		var txin protos.TX_TXIN
		txin.Addr = smartContractAddr
		_keySplit := strings.Split(v, ":")
		txin.SourceTxHash = _keySplit[0]
		_nIdx, _ := strconv.ParseUint(_keySplit[1], 10, 64)
		txin.Idx = uint32(_nIdx)
		txins = append(txins, &txin)

		sourceTxout, err := store.GetCoin(v)
		if err != nil {
			continue
		}

		if sourceTxout.Value <= amountTemp {

			var txout protos.TX_TXOUT

			txout.Value = sourceTxout.Value
			txout.Addr = bargainAddr
			txout.Attr = sourceTxout.Attr

			txouts = append(txouts, &txout)

			amountTemp -= sourceTxout.Value

		} else {

			var txout protos.TX_TXOUT

			txout.Value = amountTemp
			txout.Addr = bargainAddr
			txout.Attr = sourceTxout.Attr
			txouts = append(txouts, &txout)

			var txoutRe protos.TX_TXOUT

			txoutRe.Value = sourceTxout.Value - amountTemp
			txoutRe.Addr = smartContractAddr
			txoutRe.Attr = sourceTxout.Attr
			txouts = append(txouts, &txoutRe)

			amountTemp = 0
		}
	}

	return txins, txouts, nil
}

func SaveBargainTrack(store store.Store, bargain *protos.Bargain, amount uint64, timestamp int64) error {
	tmpBargainTrack, err := store.GetBargainTrack(bargain.Addr)
	if err != nil {
		return err
	}

	var bargainHistory protos.BargainHistory

	bargainHistory.Amount = amount
	bargainHistory.Timestamp = timestamp

	tmpBargainTrack.Addr = bargain.Addr
	tmpBargainTrack.Trans = append(tmpBargainTrack.Trans, &bargainHistory)

	if err := store.PutBargainTrack(tmpBargainTrack); err != nil {
		return err
	}

	return nil
}

func SaveSmartContractTrack(store store.Store, smartContractAddr string, bargain *protos.Bargain, amount uint64, timestamp int64) error {

	tmpSmartContractTrack, err := store.GetSmartContractTrack(smartContractAddr)
	if err != nil {
		return err
	}

	var smartContractHistory protos.SmartContractHistory

	smartContractHistory.BargainAddr = bargain.Addr
	smartContractHistory.BargainName = bargain.Name
	smartContractHistory.Type = "draw"
	smartContractHistory.Amount = amount
	smartContractHistory.Timestamp = timestamp

	tmpSmartContractTrack.Addr = smartContractAddr
	tmpSmartContractTrack.Trans = append(tmpSmartContractTrack.Trans, &smartContractHistory)

	if err := store.PutSmartContractTrack(tmpSmartContractTrack); err != nil {
		return err
	}

	return nil
}

func SaveDonorTrack(store store.Store, drawUUID string, sourceTX *protos.TX, bargain *protos.Bargain) error {

	for _, output := range sourceTX.Txout {
		if output.Addr != bargain.Addr {
			continue
		}

		donorTrack, donorAddr, err := GenDonorTrackByTxout(output, drawUUID, bargain, sourceTX.Timestamp)
		if err != nil {
			return err
		}

		tmpDonor, err := store.GetDonor(donorAddr)
		if err != nil {
			return err
		}

		tmpDonor.Trackings = append(tmpDonor.Trackings, donorTrack)

		if err := store.PutDonor(tmpDonor); err != nil {
			return err
		}
	}

	return nil
}

func GenDonorTrackByTxout(txout *protos.TX_TXOUT, drawUUID string, bargain *protos.Bargain, timestamp int64) (*protos.DonorTrack, string, error) {
	var donorTrack protos.DonorTrack

	_attrs := strings.Split(txout.Attr, ",")
	if len(_attrs) != 2 {
		return nil, "", errors.InvalidTxoutAttr
	}

	_donorAddr := _attrs[0]
	_donorUUID := _attrs[1]

	donorTrack.Donorid = _donorUUID
	donorTrack.Drawid = drawUUID
	donorTrack.AccountAddr = txout.Addr
	donorTrack.AccountName = bargain.Name
	donorTrack.Amount = txout.Value
	donorTrack.DonorAmount = txout.Value
	donorTrack.Type = 1
	donorTrack.Timestamp = timestamp

	return &donorTrack, _donorAddr, nil
}

func SaveProcessDrawed(store store.Store, drawUUID, smartContractAddr string, sourceTX *protos.TX, bargain *protos.Bargain, amount uint64) error {
	var processDrawed protos.ProcessDrawed

	processDrawed.DrawUUID = drawUUID
	processDrawed.SmartContractAddr = smartContractAddr
	smartContract, err := store.GetSmartContract(smartContractAddr)
	if err != nil {
		return err
	}
	processDrawed.DonorName = smartContract.Name
	processDrawed.BargainAddr = bargain.Addr
	processDrawed.BargainName = bargain.Name
	processDrawed.Amount = amount
	processDrawed.AcceptName = bargain.PartyB
	processDrawed.Timestamp = sourceTX.Timestamp

	processDrawed.Remark = sourceTX.InputData

	return store.PutProcessDrawed(&processDrawed)
}

// Drawed foundation drawing
func DrawedTx(store store.Store, args []string) (*protos.TX, error) {

	_base64TxData := args[1]

	txData, err := base64.StdEncoding.DecodeString(_base64TxData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newTX, err := utils.ParseTxByJSONBytes(txData)
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

// QueryProcessDrawed get a ProcessDrawed
func QueryProcessDrawed(store store.Store, args []string) ([]byte, error) {

	drawUUID := args[0]

	processDrawed, err := store.GetProcessDrawed(drawUUID)
	if err != nil {
		return nil, err
	}

	return json.Marshal(processDrawed)
}
