package utils

import (
	"encoding/json"

	"github.com/ecloudtime/charitycc/protos"
	"github.com/golang/protobuf/proto"
)

// ParseTxByBytes unmarshal txData into TX object
func ParseTxByBytes(txData []byte) (*protos.TX, error) {
	tx := new(protos.TX)
	err := proto.Unmarshal(txData, tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// ParseTxByJSONBytes unmarshal txData json into TX object
func ParseTxByJSONBytes(txData []byte) (*protos.TX, error) {

	tx := new(protos.TX)
	err := json.Unmarshal(txData, tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// ParseSmartContractByJSONBytes unmarshal smartContractData into SmartContract object
func ParseSmartContractByJSONBytes(smartContractData []byte) (*protos.SmartContract, error) {

	smartContract := new(protos.SmartContract)
	err := json.Unmarshal(smartContractData, smartContract)
	if err != nil {
		return nil, err
	}

	return smartContract, nil
}

// ParseBargainByJSONBytes unmarshal bargainData into bargain object
func ParseBargainByJSONBytes(bargainData []byte) (*protos.Bargain, error) {

	bargain := new(protos.Bargain)
	err := json.Unmarshal(bargainData, bargain)
	if err != nil {
		return nil, err
	}

	return bargain, nil
}

// ParseDonorByJSONBytes unmarshal donorData into Donor object
func ParseDonorByJSONBytes(donorData []byte) (*protos.Donor, error) {

	donor := new(protos.Donor)
	err := json.Unmarshal(donorData, donor)
	if err != nil {
		return nil, err
	}

	return donor, nil
}

// ParseDonorContributionByJSONBytes unmarshal contributionData into Contribution object
func ParseDonorContributionByJSONBytes(contributionData []byte) (*protos.DonorContribution, error) {

	contribution := new(protos.DonorContribution)
	err := json.Unmarshal(contributionData, contribution)
	if err != nil {
		return nil, err
	}

	return contribution, nil
}

// ParseDonorTrackByJSONBytes unmarshal trackData into Track object
func ParseDonorTrackByJSONBytes(trackData []byte) (*protos.DonorTrack, error) {

	tracking := new(protos.DonorTrack)
	err := json.Unmarshal(trackData, tracking)
	if err != nil {
		return nil, err
	}

	return tracking, nil
}
