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

// Donated donor donated
func Donated(store store.Store, args []string) ([]byte, error) {

	_base64TxData := args[2]

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

	return nil, nil
}

// DoDonating donor donating all proccess
func DoDonating(store store.Store, args []string) ([]byte, error) {

	_donorUUID := args[0]
	_donorAddr := args[1]
	_smartContractAddr := args[2]
	_amount := args[3]
	_bankAddr := args[4]
	_channelAddr := args[5]
	_fundAddr := args[6]

	amount, err := strconv.ParseUint(_amount, 10, 64)
	if err != nil {
		return nil, err
	}

	// querySmartContract
	smartContract, err := QuerySmartContractObj(store, _smartContractAddr)
	if err != nil {
		return nil, errors.InvalidSmartContract
	}

	// genContribution
	contribution := GenDonatingContribution(_donorUUID, *smartContract, amount)

	bankAccount, err := QueryAccountObj(store, _bankAddr)
	if err != nil {
		return nil, errors.InvalidAccount
	}
	if bankAccount.Balance < amount {
		return nil, errors.AccountNotEnoughBalance
	}

	utxo := coin.MakeUTXO(store)

	// changeCoin
	changeCoinTx, err := GenDonatingChangeCoinTxData(_donorUUID, _bankAddr, _donorAddr, amount, *bankAccount)
	if err != nil {
		return nil, err
	}

	execResult, err := utxo.Execute(changeCoinTx)
	if err != nil {
		return nil, err
	}

	// donated
	donatedTx, donatedTrack := GenDonatingDonatedTxData(_donorAddr, _donorUUID, execResult.TxHash, _smartContractAddr, _channelAddr, _fundAddr, amount, *smartContract)

	execResult, err = utxo.Execute(donatedTx)
	if err != nil {
		return nil, err
	}

	// SaveDonorContributionTrack addContribution addTrack
	err = SaveDonorContributionTrack(store, _donorAddr, contribution, donatedTrack)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// GenDonatingContribution gen donate Contribution
func GenDonatingContribution(donorUUID string, smartContract protos.SmartContract, amount uint64) protos.DonorContribution {
	var contribution protos.DonorContribution

	contribution.Donorid = donorUUID
	contribution.SmartContractName = smartContract.Name
	contribution.SmartContractAddr = smartContract.Addr
	contribution.Amount = amount
	contribution.Timestamp = time.Now().UTC().Unix()

	return contribution
}

// GenDonatingChangeCoinTxData gen donate changecoin tx
func GenDonatingChangeCoinTxData(donorUUID, bankAddr, donorAddr string, amount uint64, bankAccount protos.Account) (*protos.TX, error) {
	var tx protos.TX

	tx.Version = TxVersion
	tx.Timestamp = time.Now().UTC().Unix()

	tx.Txin = GenChangeCoinTxin(bankAddr, bankAccount)

	tx.Txout = GenChangeCoinTxout(bankAddr, donorAddr, bankAccount, amount)

	tx.InputData = donorUUID

	tx.Founder = donorAddr

	return &tx, nil
}

func GenChangeCoinTxin(bankAddr string, bankAccount protos.Account) []*protos.TX_TXIN {
	var txins []*protos.TX_TXIN
	var txin protos.TX_TXIN

	txin.Addr = bankAddr
	var _key string
	for k, _ := range bankAccount.GetTxouts() {
		_key = k
		break
	}

	_keySplit := strings.Split(_key, ":")
	txin.SourceTxHash = _keySplit[0]
	_nIdx, _ := strconv.ParseUint(_keySplit[1], 10, 64)
	txin.Idx = uint32(_nIdx)

	txins = append(txins, &txin)

	return txins
}

func GenChangeCoinTxout(bankAddr, donorAddr string, bankAccount protos.Account, amount uint64) []*protos.TX_TXOUT {
	var txouts []*protos.TX_TXOUT

	rechangeTxout := GenChangeCoinRechangeTxout(bankAddr, bankAccount, amount)
	txouts = append(txouts, &rechangeTxout)

	changeCoinTxout := GenChangeCoinChangeCoinTxout(donorAddr, amount)
	txouts = append(txouts, &changeCoinTxout)

	return txouts
}

func GenChangeCoinRechangeTxout(bankAddr string, bankAccount protos.Account, amount uint64) protos.TX_TXOUT {
	var txout protos.TX_TXOUT

	txout.Value = bankAccount.Balance - amount
	txout.Addr = bankAddr

	return txout
}

func GenChangeCoinChangeCoinTxout(donorAddr string, amount uint64) protos.TX_TXOUT {
	var txout protos.TX_TXOUT

	txout.Value = amount
	txout.Addr = donorAddr

	return txout
}

// GenDonatingDonatedTxData gen donate donated tx
func GenDonatingDonatedTxData(donorAddr, donorUUID, sourceTxHash, smartContractAddr, channelAddr, fundAddr string, amount uint64, smartContract protos.SmartContract) (*protos.TX, []*protos.DonorTrack) {
	var tx protos.TX
	var donatedTrack []*protos.DonorTrack

	tx.Version = TxVersion
	tx.Timestamp = time.Now().UTC().Unix()

	tx.Txin = GenDonatingTxin(donorAddr, sourceTxHash)

	tx.Txout, donatedTrack = GenDonatingTxout(donorAddr, donorUUID, smartContractAddr, channelAddr, fundAddr, amount, smartContract)

	tx.InputData = donorUUID

	tx.Founder = channelAddr

	return &tx, donatedTrack
}

func GenDonatingTxin(donorAddr, sourceTxHash string) []*protos.TX_TXIN {
	var txins []*protos.TX_TXIN
	var txin protos.TX_TXIN
	txin.Addr = donorAddr
	txin.SourceTxHash = sourceTxHash
	txin.Idx = 1

	txins = append(txins, &txin)

	return txins
}

func GenDonatingTxout(donorAddr, donorUUID, smartContractAddr, channelAddr, fundAddr string, amount uint64, smartContract protos.SmartContract) ([]*protos.TX_TXOUT, []*protos.DonorTrack) {
	var txouts []*protos.TX_TXOUT
	var tracks []*protos.DonorTrack

	smartContractTxout, smartContractTrack := GenDonatingSmartContractTxout(donorAddr, donorUUID, smartContractAddr, amount, smartContract)
	txouts = append(txouts, &smartContractTxout)
	tracks = append(tracks, &smartContractTrack)

	channelTxout, channelTrack := GenDonatingChannelTxout(donorAddr, donorUUID, channelAddr, amount, smartContract)
	txouts = append(txouts, &channelTxout)
	tracks = append(tracks, &channelTrack)

	fundTxout, fundTrack := GenDonatingFundTxout(donorAddr, donorUUID, fundAddr, amount, smartContract)
	txouts = append(txouts, &fundTxout)
	tracks = append(tracks, &fundTrack)

	return txouts, tracks
}

func GenDonatingSmartContractTxout(donorAddr, donorUUID, smartContractAddr string, amount uint64, smartContract protos.SmartContract) (protos.TX_TXOUT, protos.DonorTrack) {

	var txout protos.TX_TXOUT

	txout.Addr = smartContractAddr
	txout.Value = amount / 1000 * (1000 - smartContract.FundManangerFee - smartContract.ChannelFee)

	txout.Attr = donorAddr + "," + donorUUID

	var track protos.DonorTrack

	track.Donorid = donorUUID
	track.AccountAddr = smartContractAddr
	track.Amount = txout.Value
	track.DonorAmount = amount
	track.Type = 0
	track.Timestamp = time.Now().UTC().Unix()

	return txout, track
}

func GenDonatingChannelTxout(donorAddr, donorUUID, channelAddr string, amount uint64, smartContract protos.SmartContract) (protos.TX_TXOUT, protos.DonorTrack) {

	var txout protos.TX_TXOUT

	txout.Addr = channelAddr
	txout.Value = amount / 1000 * smartContract.ChannelFee

	txout.Attr = donorAddr + "," + donorUUID

	var track protos.DonorTrack

	track.Donorid = donorUUID
	track.AccountAddr = channelAddr
	track.Amount = txout.Value
	track.DonorAmount = amount
	track.Type = 2
	track.Timestamp = time.Now().UTC().Unix()

	return txout, track
}

func GenDonatingFundTxout(donorAddr, donorUUID, fundAddr string, amount uint64, smartContract protos.SmartContract) (protos.TX_TXOUT, protos.DonorTrack) {

	var txout protos.TX_TXOUT

	txout.Addr = fundAddr
	txout.Value = amount / 1000 * smartContract.FundManangerFee

	txout.Attr = donorAddr + "," + donorUUID

	var track protos.DonorTrack

	track.Donorid = donorUUID
	track.AccountAddr = fundAddr
	track.Amount = txout.Value
	track.DonorAmount = amount
	track.Type = 3
	track.Timestamp = time.Now().UTC().Unix()

	return txout, track
}

func SaveDonorContributionTrack(store store.Store, donorAddr string, contribution protos.DonorContribution, donatedTrack []*protos.DonorTrack) error {
	tmpDonor, err := store.GetDonor(donorAddr)
	if err != nil {
		return err
	}

	tmpDonor.Contributions = append(tmpDonor.Contributions, &contribution)

	for _, track := range donatedTrack {
		tmpDonor.Trackings = append(tmpDonor.Trackings, track)
	}

	if err := store.PutDonor(tmpDonor); err != nil {
		return err
	}

	return nil
}
