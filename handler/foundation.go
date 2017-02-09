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

// RegisterSmartContract register foundation smartContract
func RegisterSmartContract(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_base64SmartContractData := args[1]
	_base64SourcSign := args[2]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_base64SmartContractData))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.RegisterSmartContract(store, args)
}

// ChangeSmartContractStatus change foundation smartContract status
func ChangeSmartContractStatus(store store.Store, args []string) ([]byte, error) {

	if len(args) != 4 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_smartContractAddr := args[1]
	_smartContractStatus := args[2]
	_base64SourcSign := args[3]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_smartContractAddr + _smartContractStatus))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.ChangeSmartContractStatus(store, args)
}

func QuerySmartContract(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QuerySmartContract(store, args)
}

func QuerySmartContracts(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QuerySmartContracts(store, args)
}

// RegisterBargain register foundation bargain
func RegisterBargain(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_base64BargainData := args[1]
	_base64SourcSign := args[2]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_base64BargainData))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.RegisterBargain(store, args)
}

func QueryBargain(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryBargain(store, args)
}

func QueryBargains(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryBargains(store, args)
}

// Drawed foundation drawing
func Drawed(store store.Store, args []string) ([]byte, error) {

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

	return service.Drawed(store, args)
}

// DoDrawing foundation drawing and addTrack
func DoDrawing(store store.Store, args []string) ([]byte, error) {

	// drawed

	// addTrack

	return nil, nil
}
