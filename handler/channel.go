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

// RegisterDonor register donor
func RegisterDonor(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_base64DonorData := args[1]
	_base64SourcSign := args[2]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_base64DonorData))

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

	if len(args) != 4 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" {
		return nil, errors.InvalidArgs
	}

	_sourceAddr := args[0]
	_donorAddr := args[1]
	_base64TxData := args[2]
	_base64SourcSign := args[3]

	_sourcePublicKey, err := service.QueryAccountRsaPublicKey(store, _sourceAddr)
	if err != nil {
		return nil, err
	}

	_sourcSign, err := base64.StdEncoding.DecodeString(_base64SourcSign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_donorAddr + _base64TxData))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _sourcePublicKey, _sourcSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.Donated(store, args)
}

// DoDonating donor donating all proccess
func DoDonating(store store.Store, args []string) ([]byte, error) {

	if len(args) != 8 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" || args[4] == "" || args[5] == "" || args[6] == "" || args[7] == "" {
		return nil, errors.InvalidArgs
	}

	_donorUUID := args[0]
	_donorAddr := args[1]
	_smartContractAddr := args[2]
	_amount := args[3]
	_bankAddr := args[4]
	_channelAddr := args[5]
	_fundAddr := args[6]
	_donorBase64Sign := args[7]

	_donorPublicKey, err := service.QueryAccountRsaPublicKey(store, _donorAddr)
	if err != nil {
		return nil, err
	}

	_donorSign, err := base64.StdEncoding.DecodeString(_donorBase64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(_donorUUID + _donorAddr + _smartContractAddr + _amount + _bankAddr + _channelAddr + _fundAddr))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], _donorPublicKey, _donorSign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	return service.DoDonating(store, args)
}
