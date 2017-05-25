package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Init chaincode init
func (t *CharityCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// func (t *CharityCC) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	return shim.Success(nil)
}
