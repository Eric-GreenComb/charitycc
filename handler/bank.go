package handler

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"

	"github.com/ecloudtime/charitycc/core/store"
	"github.com/ecloudtime/charitycc/errors"
	"github.com/ecloudtime/charitycc/service"
	"github.com/ecloudtime/charitycc/utils"
)

// RegisterBank register bank ,with systemadmin pub/priv key
func RegisterBank(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	addr := args[0]
	publicKey := args[1]

	base64Sign := args[2]
	sysadminSign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(addr + publicKey))

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

// DestoryCoinbase destory addr coinbase
func DestoryCoinbase(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	_targetAddr := args[0]
	_bankAddr := args[1]
	_base64Sign := args[2]

	bankRsaPublicKey, err := service.QueryAccountRsaPublicKey(store, _bankAddr)
	if err != nil {
		return nil, err
	}

	bankSign, err := base64.StdEncoding.DecodeString(_base64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_targetAddr))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], bankRsaPublicKey, bankSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.DestoryCoinbase(store, args)
}

// ChangeCoin user change coin with CNY 1yuan = 1000000 coin
func ChangeCoin(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_base64TxData := args[1]
	_base64SourcSign := args[2]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_base64TxData))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.ChangeCoin(store, args)
}

// BuyCoin user change coin with CNY 1yuan = 1000000 coin
func BuyCoin(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_base64TxData := args[1]
	_base64SourcSign := args[2]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_base64TxData))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.BuyCoin(store, args)
}
