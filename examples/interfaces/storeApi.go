package main

//START OMIT
type Datastorer interface {
	Put(key string, value interface{}) error
	Get(key string) (interface{}, bool, error)
	Remove(key string) error
}

//END OMIT
