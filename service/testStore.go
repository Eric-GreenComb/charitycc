package service

import (
	"encoding/json"
	"github.com/CebEcloudTime/charitycc/core/store"
)

// TestStore
func TestStore(store store.Store, args []string) ([]byte, error) {

	key := args[0]
	value := args[1]

	return nil, store.PutTest(key, value)
}

// QueryStore get a donor
func QueryStore(store store.Store, args []string) ([]byte, error) {

	key := args[0]

	value, err := store.GetTest(key)
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

// TestArray
func TestArray(store store.Store, args []string) ([]byte, error) {

	key := args[0]
	value := args[1]

	return nil, store.PutTestArray(key, value)
}

// QueryArray
func QueryArray(store store.Store, args []string) ([]byte, error) {

	key := args[0]

	array, err := store.GetTestArray(key)
	if err != nil {
		return nil, err
	}

	return json.Marshal(array)
}

// TestMap
func TestMap(store store.Store, args []string) ([]byte, error) {

	rootKey := args[0]
	key := args[1]
	value := args[2]

	return nil, store.PutTestMap(rootKey, key, value)
}

// QueryMap
func QueryMap(store store.Store, args []string) ([]byte, error) {

	key := args[0]

	_map, err := store.GetTestMap(key)
	if err != nil {
		return nil, err
	}

	return json.Marshal(_map)
}
