package utils

import (
	"encoding/json"
	"github.com/CebEcloudTime/charitycc/protos"
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

// ParseTxByJsonBytes unmarshal txData into TX object
func ParseTxByJsonBytes(txData []byte) (*protos.TX, error) {

	tx := new(protos.TX)
	err := json.Unmarshal(txData, tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// ParseTreatyByJsonBytes unmarshal treatyData into Treaty object
func ParseTreatyByJsonBytes(treatyData []byte) (*protos.Treaty, error) {

	treaty := new(protos.Treaty)
	err := json.Unmarshal(treatyData, treaty)
	if err != nil {
		return nil, err
	}

	return treaty, nil
}

// ParseContractByJsonBytes unmarshal contractData into contract object
func ParseContractByJsonBytes(contractData []byte) (*protos.Contract, error) {

	contract := new(protos.Contract)
	err := json.Unmarshal(contractData, contract)
	if err != nil {
		return nil, err
	}

	return contract, nil
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
