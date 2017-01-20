package utils

import (
	"fmt"
)

func GenerateAccountKey(addr string) string {
	return fmt.Sprintf("account_addr_%s", addr)
}
