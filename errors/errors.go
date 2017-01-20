package errors

import "errors"

var (
	// ErrBase64Decoding verify rsa signture error
	Base64Decoding = errors.New("base64 decoding error")

	// ErrInvalidTX
	InvalidTX = errors.New("transaction invalid")

	// ErrInvalidTX
	InvalidAccount = errors.New("account invalid")

	// ErrInvalidTxKey returned if given key is invalid
	InvalidTxKey = errors.New("invalid tx key")

	// ErrMustCoinbase
	MustCoinbase = errors.New("tx must be coinbase")

	// ErrCantCoinbase
	CantCoinbase = errors.New("tx must not be coinbase")

	// ErrTxInOutNotBalance returned when txouts + fee != txins
	TxInOutNotBalance = errors.New("tx in & out not balance")

	// ErrTxOutMoreThanTxIn
	TxOutMoreThanTxIn = errors.New("tx out more than tx in")

	// ErrCollisionTxOut
	CollisionTxOut = errors.New("account has collision tx out")

	// ErrTxNoFounder
	TxNoFounder = errors.New("tx has no founder")

	// ErrAccountNoTxOut
	AccountNoTxOut = errors.New("account has no such tx out")

	// ErrAccountNotEnoughBalance
	AccountNotEnoughBalance = errors.New("account has not enough balance")

	// ErrTxOutLock
	TxOutLock = errors.New("tx out can be spend only after until time")

	// ErrAlreadyRegisterd
	AlreadyRegisterd = errors.New("the addr has been registerd into coin")

	// ErrIncorrectNumberArguments incorrect number of arguments
	IncorrectNumberArguments = errors.New("Incorrect number of arguments")

	// ErrInvalidArgs is returned if there are some unused args or not enough args in params
	InvalidArgs = errors.New("invalid args")

	// ErrVerifyRsaSign verify rsa signture error
	VerifyRsaSign = errors.New("verify rsa signture error")
)
