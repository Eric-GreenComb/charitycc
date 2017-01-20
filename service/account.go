package service

import (
	"encoding/json"
	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
)

// RegisterAccount register account
func RegisterAccount(store store.Store, args []string) ([]byte, error) {

	userID := args[0]
	publicKey := args[1]

	addr := utils.GenAddr(userID, publicKey)

	if tmpaccount, err := store.GetAccount(addr); err == nil && tmpaccount != nil && tmpaccount.Addr == addr {
		return nil, errors.AlreadyRegisterd
	}

	account := &protos.Account{
		Addr:         addr,
		Balance:      0,
		RsaPublicKey: publicKey,
		Txouts:       make(map[string]*protos.TX_TXOUT),
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
