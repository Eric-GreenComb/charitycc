package handler

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"

	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/utils"
)

// DoDonate do donate
func DoDonate(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" || args[1] == "" || args[2] == "" {
		return nil, errors.InvalidArgs
	}

	donorID := args[0]
	publicKey := args[1]

	base64Sign := args[2]
	sign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	hash := sha256.Sum256([]byte(donorID + publicKey))

	bVerify := utils.RsaVerify(crypto.SHA256, hash[:], []byte(publicKey), sign)
	if !bVerify {
		return nil, errors.VerifyRsaSign
	}

	// buy cebcoin, make bank buy tx(in: bank; out: donor) and changecoin

	// get treaty, get treaty by treaty id

	// do donate, make donate tx1(in: donor; out: treaty) tx2(in: donor; out: channel) tx3(in: donor; out: treaty)

	// save tx1,tx2,tx3 and save tx1,tx2,tx3 info into donor info'addr

	return nil, nil
}
