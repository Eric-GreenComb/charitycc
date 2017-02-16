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

func GenerateSmartContractTrackKey(addr string) string {
	return fmt.Sprintf("smartcontract_track_addr_%s", addr)
}

func GenerateFundKey(addr string) string {
	return fmt.Sprintf("fund_addr_%s", addr)
}

func GenerateChannelKey(addr string) string {
	return fmt.Sprintf("channel_addr_%s", addr)
}

func GenerateProcessDonoredKey(addr string) string {
	return fmt.Sprintf("process_donored_addr_%s", addr)
}

func GenerateProcessDrawedKey(addr string) string {
	return fmt.Sprintf("process_drawed_addr_%s", addr)
}
