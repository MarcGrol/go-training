package main

import "fmt"

type SimplisticDatastore struct {
	data map[string]interface{}
}

func NewSimpplisticDatastore() Datastorer {
	return &SimplisticDatastore{
		data: map[string]interface{}{},
	}
}

func (ds *SimplisticDatastore) Put(key string, value interface{}) error {
	ds.data[key] = value
	fmt.Printf("Put %+v\n", value)
	return nil
}

func (ds *SimplisticDatastore) Get(key string) (interface{}, bool, error) {
	value, found := ds.data[key]
	if !found {
		fmt.Printf("Key %s not found\n", key)
	} else {
		fmt.Printf("Got %+v\n", value)
	}
	return value, found, nil
}

func (ds *SimplisticDatastore) Remove(key string) error {
	delete(ds.data, key)
	fmt.Printf("Remove key %s\n", key)
	return nil
}
