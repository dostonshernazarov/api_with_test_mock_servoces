package kv

import (
	"errors"
	"github.com/spf13/cast"
	"sync"
)

type InMemory struct {
	inst sync.Map
}

func NewInMemory() *InMemory {
	return &InMemory{}
}

func (im *InMemory) Set(key string, value string) error {
	im.inst.Store(key, value)
	return nil
}

func (im *InMemory) Get(key string) (string, error) {
	value, exists := im.inst.Load(key)
	if !exists {
		return "", errors.New("not found error")
	}

	return cast.ToString(value), nil
}

func (im *InMemory) Delete(key string) error {
	im.inst.Delete(key)
	return nil
}

func (im *InMemory) GetAll() (map[string]string, error) {
	pairs := make(map[string]string)

	im.inst.Range(func(key, value interface{}) bool {
		pairs[cast.ToString(key)] = cast.ToString(value)
		return true
	})

	return pairs, nil
}
