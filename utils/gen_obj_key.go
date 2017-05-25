package utils

import (
	"fmt"
)

// GenerateAccountKey gen account key
func GenerateAccountKey(addr string) string {
	return fmt.Sprintf("account_addr_%s", addr)
}

// GenerateCoinKey gen coin key
func GenerateCoinKey(addr string) string {
	return fmt.Sprintf("coin_addr_%s", addr)
}

// GenerateSmartContractKey gen smartcontract key
func GenerateSmartContractKey(addr string) string {
	return fmt.Sprintf("smartcontract_addr_%s", addr)
}

// GenerateBargainKey gen bargain key
func GenerateBargainKey(addr string) string {
	return fmt.Sprintf("bargain_addr_%s", addr)
}

// GenerateBargainTrackKey gen bargain track key
func GenerateBargainTrackKey(addr string) string {
	return fmt.Sprintf("bargain_track_addr_%s", addr)
}

// GenerateDonorKey gen donor key
func GenerateDonorKey(addr string) string {
	return fmt.Sprintf("donor_addr_%s", addr)
}

// GenerateSmartContractExtKey gen smartcontract ext key
func GenerateSmartContractExtKey(addr string) string {
	return fmt.Sprintf("smartcontract_ext_addr_%s", addr)
}

// GenerateSmartContractTrackKey gen smartcontract track key
func GenerateSmartContractTrackKey(addr string) string {
	return fmt.Sprintf("smartcontract_track_addr_%s", addr)
}

// GenerateFundKey gen fund key
func GenerateFundKey(addr string) string {
	return fmt.Sprintf("fund_addr_%s", addr)
}

// GenerateChannelKey gen channel key
func GenerateChannelKey(addr string) string {
	return fmt.Sprintf("channel_addr_%s", addr)
}

// GenerateProcessDonoredKey gen process donored key
func GenerateProcessDonoredKey(addr string) string {
	return fmt.Sprintf("process_donored_addr_%s", addr)
}

// GenerateProcessDrawedKey gen process drawed key
func GenerateProcessDrawedKey(addr string) string {
	return fmt.Sprintf("process_drawed_addr_%s", addr)
}
