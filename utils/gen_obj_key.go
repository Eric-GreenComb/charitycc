package utils

import (
	"fmt"
)

func GenerateAccountKey(addr string) string {
	return fmt.Sprintf("account_addr_%s", addr)
}

func GenerateTreatyKey(addr string) string {
	return fmt.Sprintf("treaty_addr_%s", addr)
}

func GenerateContractKey(addr string) string {
	return fmt.Sprintf("contract_addr_%s", addr)
}
