package utils

import (
	"crypto/sha256"
	"encoding/hex"
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
