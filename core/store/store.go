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

	GetSmartContract(string) (*protos.SmartContract, error)
	PutSmartContract(*protos.SmartContract) error

	GetBargain(string) (*protos.Bargain, error)
	PutBargain(*protos.Bargain) error

	GetDonor(string) (*protos.Donor, error)
	PutDonor(*protos.Donor) error
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

// GetSmartContract returns SmartContract from world states
func (s *CCStore) GetSmartContract(addr string) (*protos.SmartContract, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateSmartContractKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no smartContract found")
	}

	smartContract := new(protos.SmartContract)
	if err := proto.Unmarshal(data, smartContract); err != nil {
		return nil, err
	}

	return smartContract, nil
}

// PutSmartContract update or insert SmartContract into world states
func (s *CCStore) PutSmartContract(smartContract *protos.SmartContract) error {
	key := utils.GenerateSmartContractKey(smartContract.Addr)

	aBytes, err := proto.Marshal(smartContract)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}

// GetBargain returns Bargain from world states
func (s *CCStore) GetBargain(addr string) (*protos.Bargain, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateBargainKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no contract found")
	}

	bargain := new(protos.Bargain)
	if err := proto.Unmarshal(data, bargain); err != nil {
		return nil, err
	}

	return bargain, nil
}

// PutBargain update or insert Bargain into world states
func (s *CCStore) PutBargain(bargain *protos.Bargain) error {
	key := utils.GenerateBargainKey(bargain.Addr)

	aBytes, err := proto.Marshal(bargain)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}

// GetDonor returns donor from world states
func (s *CCStore) GetDonor(addr string) (*protos.Donor, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateDonorKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no donor found")
	}

	donor := new(protos.Donor)
	if err := proto.Unmarshal(data, donor); err != nil {
		return nil, err
	}

	return donor, nil
}

// PutDonor update or insert donor into world states
func (s *CCStore) PutDonor(donor *protos.Donor) error {
	key := utils.GenerateDonorKey(donor.Addr)

	aBytes, err := proto.Marshal(donor)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}
