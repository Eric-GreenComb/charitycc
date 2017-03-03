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

// RegisterChannel register Channel
func RegisterChannel(store store.Store, args []string) ([]byte, error) {
	_channelAddr := args[0]
	_channelPublicKey := args[1]

	_, err := InitAccount(store, _channelAddr, _channelPublicKey)
	if err != nil {
		return nil, err
	}

	_, err = InitChannel(store, _channelAddr)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func InitChannel(store store.Store, addr string) ([]byte, error) {
	var newChannel protos.Channel
	newChannel.Addr = addr

	if tmpChannel, err := store.GetChannel(newChannel.Addr); err == nil && tmpChannel != nil && tmpChannel.Addr == newChannel.Addr {
		return nil, errors.AlreadyRegisterd
	}

	if err := store.PutChannel(&newChannel); err != nil {
		return nil, err
	}

	return nil, nil

}

// QueryChannel get a Channel
func QueryChannel(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	channel, err := store.GetChannel(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(channel)
}

// Donated donor donated
func Donated(store store.Store, args []string) ([]byte, error) {

	_donorAddr := args[1]
	_donorUUID := args[2]
	_smartContractAddr := args[3]
	_base64TxData := args[4]

	txData, err := base64.StdEncoding.DecodeString(_base64TxData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	// querySmartContract
	smartContract, err := QuerySmartContractObj(store, _smartContractAddr)
	if err != nil {
		return nil, errors.InvalidSmartContract
	}

	sourceTX, err := utils.ParseTxByJsonBytes(txData)
	if err != nil {
		return nil, err
	}

	if sourceTX.Founder == "" {
		return nil, errors.TxNoFounder
	}

	amount := GetAmountByTxouts(sourceTX)

	// genContribution
	contribution := GenDonatingContribution(_donorUUID, *smartContract, amount, sourceTX.Timestamp)

	// donated
	donatedTx, donatedTrack, processDonored := GenDonatedTxData(_donorAddr, _donorUUID, *smartContract, *sourceTX, amount)

	utxo := coin.MakeUTXO(store)

	_, err = utxo.Execute(donatedTx)
	if err != nil {
		return nil, err
	}

	// SaveDonorContributionTrack addContribution addTrack
	err = SaveDonorContributionTrack(store, _donorAddr, contribution, donatedTrack)
	if err != nil {
		return nil, err
	}

	err = SaveSmartContractExt(store, _smartContractAddr, amount, processDonored.SmartContractAmount)
	if err != nil {
		return nil, err
	}

	processDonored.Remark = sourceTX.InputData

	store.PutProcessDonored(processDonored)

	SaveFundFee(store, smartContract.FundAddr, amount, processDonored.SmartContractAmount)

	SaveChannelFee(store, smartContract.ChannelAddr, processDonored.ChannelAmount)

	return nil, nil
}

func GetAmountByTxouts(tx *protos.TX) uint64 {
	var amount uint64
	for _, output := range tx.Txout {
		amount += output.Value
	}
	return amount
}

// GenDonatedTxData gen donate donated tx
func GenDonatedTxData(donorAddr, donorUUID string, smartContract protos.SmartContract, sourceTX protos.TX, amount uint64) (*protos.TX, []*protos.DonorTrack, *protos.ProcessDonored) {
	var tx protos.TX
	var donatedTrack []*protos.DonorTrack

	smartContractAddr := smartContract.Addr
	channelAddr := smartContract.ChannelAddr
	fundAddr := smartContract.FundAddr

	tx.Version = sourceTX.Version
	tx.Timestamp = sourceTX.Timestamp

	tx.Txin = sourceTX.Txin

	var processDonored *protos.ProcessDonored
	tx.Txout, donatedTrack, processDonored = GenDonatingTxout(donorAddr, donorUUID, smartContractAddr, channelAddr, fundAddr, amount, smartContract, sourceTX.Timestamp)

	tx.InputData = donorUUID

	tx.Founder = channelAddr

	return &tx, donatedTrack, processDonored
}

func SaveSmartContractExt(store store.Store, smartContractAddr string, amount, realAmount uint64) error {
	tmpSmartContractExt, err := store.GetSmartContractExt(smartContractAddr)
	if err != nil {
		return err
	}

	tmpSmartContractExt.Total += amount
	tmpSmartContractExt.ValidTotal += realAmount
	tmpSmartContractExt.DonateNumber += 1

	if err := store.PutSmartContractExt(tmpSmartContractExt); err != nil {
		return err
	}

	return nil
}

// GenDonatingContribution gen donate Contribution
func GenDonatingContribution(donorUUID string, smartContract protos.SmartContract, amount uint64, timestamp int64) protos.DonorContribution {
	var contribution protos.DonorContribution

	contribution.Donorid = donorUUID
	contribution.SmartContractName = smartContract.Name
	contribution.SmartContractAddr = smartContract.Addr
	contribution.Amount = amount
	contribution.Timestamp = timestamp

	return contribution
}

func GenChangeCoinTxin(bankAddr string, bankAccount protos.Account) []*protos.TX_TXIN {
	var txins []*protos.TX_TXIN
	var txin protos.TX_TXIN

	txin.Addr = bankAddr
	_key := bankAccount.CoinKey[0]

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

func GenDonatingTxin(donorAddr, sourceTxHash string) []*protos.TX_TXIN {
	var txins []*protos.TX_TXIN
	var txin protos.TX_TXIN
	txin.Addr = donorAddr
	txin.SourceTxHash = sourceTxHash
	txin.Idx = 1

	txins = append(txins, &txin)

	return txins
}

func GenDonatingTxout(donorAddr, donorUUID, smartContractAddr, channelAddr, fundAddr string, amount uint64, smartContract protos.SmartContract, timestamp int64) ([]*protos.TX_TXOUT, []*protos.DonorTrack, *protos.ProcessDonored) {
	var txouts []*protos.TX_TXOUT
	var tracks []*protos.DonorTrack

	smartContractTxout, smartContractTrack := GenDonatingSmartContractTxout(donorAddr, donorUUID, smartContractAddr, amount, smartContract, timestamp)
	txouts = append(txouts, &smartContractTxout)
	tracks = append(tracks, &smartContractTrack)

	channelTxout, channelTrack := GenDonatingChannelTxout(donorAddr, donorUUID, channelAddr, amount, smartContract, timestamp)
	txouts = append(txouts, &channelTxout)
	tracks = append(tracks, &channelTrack)

	fundTxout, fundTrack := GenDonatingFundTxout(donorAddr, donorUUID, fundAddr, amount, smartContract, timestamp)
	txouts = append(txouts, &fundTxout)
	tracks = append(tracks, &fundTrack)

	var processDonored protos.ProcessDonored

	processDonored.DonorUUID = donorUUID
	processDonored.DonorAddr = donorAddr
	processDonored.SmartContractAddr = smartContract.Addr
	processDonored.SmartContractName = smartContract.Name
	processDonored.FundName = smartContract.FundName
	processDonored.ChannelName = smartContract.ChannelName
	processDonored.Amount = amount
	processDonored.SmartContractAmount = smartContractTrack.Amount
	processDonored.ChannelAmount = channelTrack.Amount
	processDonored.FundAmount = fundTrack.Amount
	processDonored.Timestamp = timestamp

	return txouts, tracks, &processDonored
}

func GenDonatingSmartContractTxout(donorAddr, donorUUID, smartContractAddr string, amount uint64, smartContract protos.SmartContract, timestamp int64) (protos.TX_TXOUT, protos.DonorTrack) {

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
	track.Timestamp = timestamp

	return txout, track
}

func GenDonatingChannelTxout(donorAddr, donorUUID, channelAddr string, amount uint64, smartContract protos.SmartContract, timestamp int64) (protos.TX_TXOUT, protos.DonorTrack) {

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
	track.Timestamp = timestamp

	return txout, track
}

func GenDonatingFundTxout(donorAddr, donorUUID, fundAddr string, amount uint64, smartContract protos.SmartContract, timestamp int64) (protos.TX_TXOUT, protos.DonorTrack) {

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
	track.Timestamp = timestamp

	return txout, track
}

func SaveDonorContributionTrack(store store.Store, donorAddr string, contribution protos.DonorContribution, donatedTrack []*protos.DonorTrack) error {
	tmpDonor, err := store.GetDonor(donorAddr)
	if err != nil {
		return err
	}

	tmpDonor.Total += contribution.Amount
	tmpDonor.Contributions = append(tmpDonor.Contributions, &contribution)

	for _, track := range donatedTrack {
		tmpDonor.Trackings = append(tmpDonor.Trackings, track)
	}

	if err := store.PutDonor(tmpDonor); err != nil {
		return err
	}

	return nil
}

// QueryProcessDonored get a ProcessDonored
func QueryProcessDonored(store store.Store, args []string) ([]byte, error) {

	donorUUID := args[0]

	processDonored, err := store.GetProcessDonored(donorUUID)
	if err != nil {
		return nil, err
	}

	return json.Marshal(processDonored)
}

func SaveFundFee(store store.Store, fundAddr string, amount, realAmount uint64) error {
	tmpFund, err := store.GetFund(fundAddr)
	if err != nil {
		return err
	}

	tmpFund.Total += amount
	tmpFund.ValidTotal += realAmount
	tmpFund.Balance += realAmount

	if err := store.PutFund(tmpFund); err != nil {
		return err
	}

	return nil
}

func SaveChannelFee(store store.Store, channelAddr string, amount uint64) error {
	tmpChannel, err := store.GetChannel(channelAddr)
	if err != nil {
		return err
	}

	tmpChannel.Total += amount

	if err := store.PutChannel(tmpChannel); err != nil {
		return err
	}

	return nil
}
