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

// RegisterAccount register account ,with systemadmin pub/priv key
func RegisterAccount(store store.Store, args []string) ([]byte, error) {

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

	return service.RegisterAccount(store, args)
}

func QueryAccount(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryAccount(store, args)
}

func QueryCoin(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryCoin(store, args)
}
