package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GenAddr gen addr by public key
func GenAddr(ID, publicKey string) string {
	fHash := sha256.Sum256([]byte(publicKey))
	lHash := sha256.Sum256(fHash[:])
	strHash := hex.EncodeToString(lHash[:])
	return fmt.Sprintf("%s:%s", ID, strHash)
}
