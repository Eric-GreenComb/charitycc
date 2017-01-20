package handler

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"

	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/service"
	"github.com/CebEcloudTime/charitycc/utils"
)

// RegisterFoundationAccount register foundation
func RegisterFoundationAccount(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	foundationID := args[0]
	publicKey := args[1]

	base64Sign := args[2]
	foundationSign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(foundationID + publicKey))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], []byte(publicKey), foundationSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.RegisterAccount(store, args)
}

func QueryFoundationAccount(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryAccount(store, args)
}

// RegisterTreatyAccount register foundation
func RegisterTreatyAccount(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	treatyID := args[0]
	publicKey := args[1]

	base64Sign := args[2]
	treatySign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(treatyID + publicKey))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], []byte(publicKey), treatySign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.RegisterAccount(store, args)
}

func QueryTreatyAccount(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryAccount(store, args)
}

// DoDrawing do drawing
func DoDrawing(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	treatyID := args[0]
	publicKey := args[1]

	base64Sign := args[2]
	treatySign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(treatyID + publicKey))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], []byte(publicKey), treatySign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	// get treaty account, check if treaty account < value

	// traverse account's txout, make txs(in: treaty; out: contract)...

	// util , txout > value, make tx(in: treaty; out: contract,treaty)

	// save tx1,tx2,tx3... and save tx1,tx2,tx3... info into donor info'addr

	return nil, nil
}
