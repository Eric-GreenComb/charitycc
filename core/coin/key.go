package coin

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ecloudtime/charitycc/errors"
)

// Key represents the key for a transaction in storage. It has both a
// hash and index
type Key struct {
	TxHashAsHex string
	TxIndex     uint32
}

func (k *Key) String() string {
	return fmt.Sprintf("%s:%d", k.TxHashAsHex, k.TxIndex)
}

// ParseKey parse key string into Key object, return error if something invalid happened
func ParseKey(keyStr string) (*Key, error) {
	if !strings.Contains(keyStr, ":") {
		return nil, errors.InvalidTxKey
	}

	subKeys := strings.Split(keyStr, ":")
	if len(subKeys) != 2 {
		return nil, errors.InvalidTxKey
	}

	txIdx, err := strconv.ParseUint(subKeys[1], 10, 32)
	if err != nil {
		return nil, err
	}

	return &Key{TxHashAsHex: subKeys[0], TxIndex: uint32(txIdx)}, nil
}
