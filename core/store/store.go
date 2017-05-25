package store

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"github.com/ecloudtime/charitycc/protos"
	"github.com/ecloudtime/charitycc/utils"
)

// Store interface describes the storage used by this chaincode. The interface
// was created so either the state database store can be used or a in memory
// store can be used for unit testing.
type Store interface {
	GetTran(string) (*protos.TX, bool, error)
	PutTran(*protos.TX) error

	GetAccount(string) (*protos.Account, error)
	PutAccount(*protos.Account) error

	GetCoin(string) (*protos.TX_TXOUT, error)
	PutCoin(string, *protos.TX_TXOUT) error
	DeleteCoin(string) error

	GetSmartContract(string) (*protos.SmartContract, error)
	PutSmartContract(*protos.SmartContract) error

	GetSmartContractExt(string) (*protos.SmartContractExt, error)
	PutSmartContractExt(*protos.SmartContractExt) error

	GetSmartContractTrack(string) (*protos.SmartContractTrack, error)
	PutSmartContractTrack(*protos.SmartContractTrack) error

	GetBargain(string) (*protos.Bargain, error)
	PutBargain(*protos.Bargain) error

	GetBargainTrack(string) (*protos.BargainTrack, error)
	PutBargainTrack(*protos.BargainTrack) error

	GetDonor(string) (*protos.Donor, error)
	PutDonor(*protos.Donor) error

	GetFund(string) (*protos.Foundation, error)
	PutFund(*protos.Foundation) error

	GetChannel(string) (*protos.Channel, error)
	PutChannel(*protos.Channel) error

	GetProcessDonored(string) (*protos.ProcessDonored, error)
	PutProcessDonored(*protos.ProcessDonored) error

	GetProcessDrawed(string) (*protos.ProcessDrawed, error)
	PutProcessDrawed(*protos.ProcessDrawed) error
}

// CCStore Store struct uses a chaincode stub for state access
type CCStore struct {
	stub shim.ChaincodeStubInterface
	// stub *shim.ChaincodeStub
}

// MakeCCStore returns a store for storing keys in the state
func MakeCCStore(stub shim.ChaincodeStubInterface) Store {
	// func MakeCCStore(stub *shim.ChaincodeStub) Store {

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

// GetCoin returns coin from world states
func (s *CCStore) GetCoin(addr string) (*protos.TX_TXOUT, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateCoinKey(addr)

	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no coin found")
	}

	txout := new(protos.TX_TXOUT)
	if err := proto.Unmarshal(data, txout); err != nil {
		return nil, err
	}

	return txout, nil
}

// PutCoin update or insert account into world states
func (s *CCStore) PutCoin(addr string, txout *protos.TX_TXOUT) error {
	key := utils.GenerateCoinKey(addr)

	aBytes, err := proto.Marshal(txout)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}

// DeleteCoin delete coin
func (s *CCStore) DeleteCoin(addr string) error {
	key := utils.GenerateCoinKey(addr)
	s.stub.DelState(key)
	return nil
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
		return nil, fmt.Errorf("no Bargain found")
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

// GetBargainTrack returns BargainTrack from world states
func (s *CCStore) GetBargainTrack(addr string) (*protos.BargainTrack, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateBargainTrackKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no BargainTrack found")
	}

	bargainTrack := new(protos.BargainTrack)
	if err := proto.Unmarshal(data, bargainTrack); err != nil {
		return nil, err
	}

	return bargainTrack, nil
}

// PutBargainTrack update or insert BargainTrack into world states
func (s *CCStore) PutBargainTrack(bargainTrack *protos.BargainTrack) error {
	key := utils.GenerateBargainTrackKey(bargainTrack.Addr)

	aBytes, err := proto.Marshal(bargainTrack)
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

// GetSmartContractExt returns SmartContractExt from world states
func (s *CCStore) GetSmartContractExt(addr string) (*protos.SmartContractExt, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateSmartContractExtKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no smartContract Track found")
	}

	smartContractExt := new(protos.SmartContractExt)
	if err := proto.Unmarshal(data, smartContractExt); err != nil {
		return nil, err
	}

	return smartContractExt, nil
}

// PutSmartContractExt update or insert SmartContractExt into world states
func (s *CCStore) PutSmartContractExt(smartContractExt *protos.SmartContractExt) error {
	key := utils.GenerateSmartContractExtKey(smartContractExt.Addr)

	aBytes, err := proto.Marshal(smartContractExt)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}

// GetSmartContractTrack returns smartContract Track from world states
func (s *CCStore) GetSmartContractTrack(addr string) (*protos.SmartContractTrack, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateSmartContractTrackKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no smartContract Track found")
	}

	smartContractTrack := new(protos.SmartContractTrack)
	if err := proto.Unmarshal(data, smartContractTrack); err != nil {
		return nil, err
	}

	return smartContractTrack, nil
}

// PutSmartContractTrack update or insert smartContract Track into world states
func (s *CCStore) PutSmartContractTrack(smartContractTrack *protos.SmartContractTrack) error {
	key := utils.GenerateSmartContractTrackKey(smartContractTrack.Addr)

	aBytes, err := proto.Marshal(smartContractTrack)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}

// GetFund returns fund from world states
func (s *CCStore) GetFund(addr string) (*protos.Foundation, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateFundKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no fund found : %s", addr)
	}

	foundation := new(protos.Foundation)
	if err := proto.Unmarshal(data, foundation); err != nil {
		return nil, err
	}

	return foundation, nil
}

// PutFund update or insert fund into world states
func (s *CCStore) PutFund(foundation *protos.Foundation) error {
	key := utils.GenerateFundKey(foundation.Addr)

	aBytes, err := proto.Marshal(foundation)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}

// GetChannel returns Channel from world states
func (s *CCStore) GetChannel(addr string) (*protos.Channel, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateChannelKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no Channel found")
	}

	channel := new(protos.Channel)
	if err := proto.Unmarshal(data, channel); err != nil {
		return nil, err
	}

	return channel, nil
}

// PutChannel update or insert Channel into world states
func (s *CCStore) PutChannel(channel *protos.Channel) error {
	key := utils.GenerateChannelKey(channel.Addr)

	aBytes, err := proto.Marshal(channel)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}

// GetProcessDonored returns ProcessDonored from world states
func (s *CCStore) GetProcessDonored(addr string) (*protos.ProcessDonored, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateProcessDonoredKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no ProcessDonored found")
	}

	processDonored := new(protos.ProcessDonored)
	if err := proto.Unmarshal(data, processDonored); err != nil {
		return nil, err
	}

	return processDonored, nil
}

// PutProcessDonored update or insert ProcessDonored into world states
func (s *CCStore) PutProcessDonored(processDonored *protos.ProcessDonored) error {
	key := utils.GenerateProcessDonoredKey(processDonored.DonorUUID)

	aBytes, err := proto.Marshal(processDonored)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}

// GetProcessDrawed returns ProcessDrawed from world states
func (s *CCStore) GetProcessDrawed(addr string) (*protos.ProcessDrawed, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}
	key := utils.GenerateProcessDrawedKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no ProcessDrawed found")
	}

	processDrawed := new(protos.ProcessDrawed)
	if err := proto.Unmarshal(data, processDrawed); err != nil {
		return nil, err
	}

	return processDrawed, nil
}

// PutProcessDrawed update or insert ProcessDrawed into world states
func (s *CCStore) PutProcessDrawed(processDrawed *protos.ProcessDrawed) error {
	key := utils.GenerateProcessDrawedKey(processDrawed.DrawUUID)

	aBytes, err := proto.Marshal(processDrawed)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, aBytes)
}
