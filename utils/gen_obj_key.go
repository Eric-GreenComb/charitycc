package utils

import (
	"fmt"
)

func GenerateAccountKey(addr string) string {
	return fmt.Sprintf("account_addr_%s", addr)
}

func GenerateSmartContractKey(addr string) string {
	return fmt.Sprintf("smartcontract_addr_%s", addr)
}

func GenerateBargainKey(addr string) string {
	return fmt.Sprintf("bargain_addr_%s", addr)
}

func GenerateDonorKey(addr string) string {
	return fmt.Sprintf("donor_addr_%s", addr)
}
