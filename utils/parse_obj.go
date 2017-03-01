package utils

import (
	"encoding/json"
	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/golang/protobuf/proto"
)

// ParseTestArrayByBytes unmarshal txData into TestArray object
func ParseTestArrayByBytes(data []byte) (*protos.TestArray, error) {
	array := new(protos.TestArray)
	err := proto.Unmarshal(data, array)
	if err != nil {
		return nil, err
	}

	return array, nil
}

// ParseTestMapByBytes unmarshal txData into TestMap object
func ParseTestMapByBytes(data []byte) (*protos.TestMap, error) {
	_map := new(protos.TestMap)
	err := proto.Unmarshal(data, _map)
	if err != nil {
		return nil, err
	}

	return _map, nil
}

// ParseTxByBytes unmarshal txData into TX object
func ParseTxByBytes(txData []byte) (*protos.TX, error) {
	tx := new(protos.TX)
	err := proto.Unmarshal(txData, tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// ParseTxByJsonBytes unmarshal txData into TX object
func ParseTxByJsonBytes(txData []byte) (*protos.TX, error) {

	tx := new(protos.TX)
	err := json.Unmarshal(txData, tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// ParseSmartContractByJsonBytes unmarshal smartContractData into SmartContract object
func ParseSmartContractByJsonBytes(smartContractData []byte) (*protos.SmartContract, error) {

	smartContract := new(protos.SmartContract)
	err := json.Unmarshal(smartContractData, smartContract)
	if err != nil {
		return nil, err
	}

	return smartContract, nil
}

// ParseBargainByJsonBytes unmarshal bargainData into bargain object
func ParseBargainByJsonBytes(bargainData []byte) (*protos.Bargain, error) {

	bargain := new(protos.Bargain)
	err := json.Unmarshal(bargainData, bargain)
	if err != nil {
		return nil, err
	}

	return bargain, nil
}

// ParseDonorByJsonBytes unmarshal donorData into Donor object
func ParseDonorByJsonBytes(donorData []byte) (*protos.Donor, error) {

	donor := new(protos.Donor)
	err := json.Unmarshal(donorData, donor)
	if err != nil {
		return nil, err
	}

	return donor, nil
}

// ParseDonorContributionByJsonBytes unmarshal contributionData into Contribution object
func ParseDonorContributionByJsonBytes(contributionData []byte) (*protos.DonorContribution, error) {

	contribution := new(protos.DonorContribution)
	err := json.Unmarshal(contributionData, contribution)
	if err != nil {
		return nil, err
	}

	return contribution, nil
}

// ParseDonorTrackByJsonBytes unmarshal trackData into Track object
func ParseDonorTrackByJsonBytes(trackData []byte) (*protos.DonorTrack, error) {

	tracking := new(protos.DonorTrack)
	err := json.Unmarshal(trackData, tracking)
	if err != nil {
		return nil, err
	}

	return tracking, nil
}
