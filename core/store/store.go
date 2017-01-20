package store

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
)

// Store interface describes the storage used by this chaincode. The interface
// was created so either the state database store can be used or a in memory
// store can be used for unit testing.
type Store interface {
	GetTran(string) (*protos.TX, bool, error)
	PutTran(*protos.TX) error

	GetAccount(string) (*protos.Account, error)
	PutAccount(*protos.Account) error
}

// Store struct uses a chaincode stub for state access
type CCStore struct {
	stub shim.ChaincodeStubInterface
}

// MakeCCStore returns a store for storing keys in the state
func MakeCCStore(stub shim.ChaincodeStubInterface) Store {
	store := &CCStore{}
	store.stub = stub
	return store
}

// GetTran returns a transaction for the given hash
func (s *CCStore) GetTran(key string) (*protos.TX, bool, error) {
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, false, fmt.Errorf("Error getting state from stub:  %s", err)
	}
	if data == nil || len(data) == 0 {
		return nil, false, nil
	}

	tx, err := utils.ParseTxByBytes(data)
	if err != nil {
		return nil, false, err
	}

	return tx, true, nil
}

// PutTran adds a transaction to the state with the hash as a key
func (s *CCStore) PutTran(tx *protos.TX) error {
	txBytes, err := proto.Marshal(tx)
	if err != nil {
		return err
	}

	return s.stub.PutState(utils.GenTransactionHash(tx), txBytes)
}

// GetAccount returns account from world states
func (s *CCStore) GetAccount(addr string) (*protos.Account, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateAccountKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no account found")
	}

	account := new(protos.Account)
	if err := proto.Unmarshal(data, account); err != nil {
		return nil, err
	}

	return account, nil
}

// PutAccount update or insert account into world states
func (s *CCStore) PutAccount(account *protos.Account) error {
	key := utils.GenerateAccountKey(account.Addr)

	aBytes, err := proto.Marshal(account)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}
