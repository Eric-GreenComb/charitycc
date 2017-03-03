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

// RegisterChannel register channel
func RegisterChannel(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	_channelAddr := args[0]
	_channelPublicKey := args[1]
	_base64SysadminSign := args[2]

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SysadminSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_channelAddr + _channelPublicKey))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], []byte(SyetemAdminPublickKey), _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.RegisterChannel(store, args)
}

func QueryChannel(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryChannel(store, args)
}

// RegisterDonor register donor
func RegisterDonor(store store.Store, args []string) ([]byte, error) {

	if len(args) != 4 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_donorAddr := args[1]
	_donorPublicKey := args[2]
	_base64SourcSign := args[3]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_donorAddr + _donorPublicKey))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.RegisterDonor(store, args)
}

// AddContribution add donor contribution
func AddContribution(store store.Store, args []string) ([]byte, error) {

	if len(args) != 4 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_donorAddr := args[1]
	_base64Contribution := args[2]
	_base64SourcSign := args[3]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_donorAddr + _base64Contribution))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.AddContribution(store, args)
}

// AddTrack add donor track
func AddTrack(store store.Store, args []string) ([]byte, error) {

	if len(args) != 4 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_donorAddr := args[1]
	_base64TrackData := args[2]
	_base64SourcSign := args[3]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_donorAddr + _base64TrackData))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.AddTrack(store, args)
}

func QueryDonor(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryDonor(store, args)
}

// Donated donor donated
func Donated(store store.Store, args []string) ([]byte, error) {

	if len(args) != 7 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" || args[4] == "" || args[5] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_donorAddr := args[1]
	_donorUUID := args[2]
	_smartContractAddr := args[3]
	_base64TxData := args[4]
	_base64SourcSign := args[5]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_donorAddr + _donorUUID + _smartContractAddr + _base64TxData))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.Donated(store, args)
}

func QueryProcessDonored(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryProcessDonored(store, args)
}
