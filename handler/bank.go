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

// RegisterBank register bank ,with systemadmin pub/priv key
func RegisterBank(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	bankID := args[0]
	publicKey := args[1]

	base64Sign := args[2]
	sysadminSign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(bankID + publicKey))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], []byte(SyetemAdminPublickKey), sysadminSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.RegisterBank(store, args)
}

// Coinbase bank coinbase
func Coinbase(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	bankRsaPublicKey, err := service.QueryAccountRsaPublicKey(store, args[0])
	if err != nil {
		return nil, err
	}

	base64TxData := args[1]

	base64Sign := args[2]
	bankSign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(base64TxData))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], bankRsaPublicKey, bankSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.Coinbase(store, args)
}

// ChangeCoin user change coin with CNY 1yuan = 1000000 coin
func ChangeCoin(store store.Store, args []string) ([]byte, error) {

	if len(args) != 5 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" || args[4] == "" {
		return nil, errors.InvalidArgs
	}

	bankID := args[0]
	base64PublicKey := args[1]
	publicKey, err := base64.StdEncoding.DecodeString(base64PublicKey)
	if err != nil {
		return nil, errors.Base64Decoding
	}
	amount := args[2]
	base64TxData := args[3]

	base64Sign := args[4]
	bankSign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(bankID + base64PublicKey + amount + base64TxData))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], publicKey, bankSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.ChangeCoin(store, args)
}

func QueryBank(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryBank(store, args)
}
