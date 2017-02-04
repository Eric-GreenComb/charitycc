package utils

import (
	"crypto/sha256"
	"encoding/hex"
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

// GenTransactionHash generates the Hash for the transaction.
func GenTransactionHash(tx *protos.TX) string {
	txBytes, err := proto.Marshal(tx)
	if err != nil {
		return ""
	}

	fHash := sha256.Sum256(txBytes)
	lHash := sha256.Sum256(fHash[:])
	return hex.EncodeToString(lHash[:])
}
