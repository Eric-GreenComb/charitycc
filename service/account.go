package service

import (
	"encoding/base64"
	"encoding/json"

	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/protos"
)

// RegisterAccount register account
func RegisterAccount(store store.Store, args []string) ([]byte, error) {

	addr := args[0]
	publicKey := args[1]

	return InitAccount(store, addr, publicKey)
}

func InitAccount(store store.Store, addr, publicKey string) ([]byte, error) {

	if tmpaccount, err := store.GetAccount(addr); err == nil && tmpaccount != nil && tmpaccount.Addr == addr {
		return nil, errors.AlreadyRegisterd
	}

	account := &protos.Account{
		Addr:         addr,
		Balance:      0,
		RsaPublicKey: publicKey,
	}
	if err := store.PutAccount(account); err != nil {
		return nil, err
	}
	return nil, nil
}

// QueryAccount get user account
func QueryAccount(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	account, err := store.GetAccount(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(account)
}

// QueryAccountObj
func QueryAccountObj(store store.Store, addr string) (*protos.Account, error) {

	account, err := store.GetAccount(addr)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// QueryAccountObj
func QueryAccountRsaPublicKey(store store.Store, addr string) ([]byte, error) {

	account, err := store.GetAccount(addr)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(account.RsaPublicKey)
}

// QueryCoin get
func QueryCoin(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	txout, err := store.GetCoin(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(txout)
}
