package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// CharityCC chaincode struct
type CharityCC struct {
}

func main() {
	err := shim.Start(new(CharityCC))
	if err != nil {
		fmt.Printf("Error starting CharityCC: %s", err)
	}
}
