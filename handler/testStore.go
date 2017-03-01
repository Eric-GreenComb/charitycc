package handler

import (
	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/service"
)

// TestStore
func TestStore(store store.Store, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.IncorrectNumberArguments
	}

	return service.TestStore(store, args)
}

func QueryStore(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryStore(store, args)
}

// TestArray
func TestArray(store store.Store, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.IncorrectNumberArguments
	}

	return service.TestArray(store, args)
}

func QueryArray(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryArray(store, args)
}

// TestMap
func TestMap(store store.Store, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.IncorrectNumberArguments
	}

	return service.TestMap(store, args)
}

func QueryMap(store store.Store, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.IncorrectNumberArguments
	}

	if args[0] == "" {
		return nil, errors.InvalidArgs
	}

	return service.QueryMap(store, args)
}
