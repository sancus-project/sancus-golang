package conf

import (
	"sync"
)

var (
	registry map[interface{}]Mapper
	mu       sync.Mutex
)

func GetMapper(set interface{}) Mapper {
	mu.Lock()
	defer mu.Unlock()

	if m, ok := registry[set]; ok {
		return m
	}

	return nil
}

func registerMapper(set interface{}, m Mapper) Mapper {
	if m != nil && set != nil {
		mu.Lock()
		defer mu.Unlock()

		registry[set] = m
		return m
	}
	return nil
}

func init() {
	registry = make(map[interface{}]Mapper)
}
