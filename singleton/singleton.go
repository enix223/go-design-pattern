package singleton

import (
	"fmt"
	"sync"
)

// Singleton singleton
type Singleton struct {
}

// SayHello hello world
func (s *Singleton) SayHello() string {
	return "hello world"
}

// mu make sure instance should only be instance once
var mu sync.Mutex
var instance *Singleton = nil

// GetSingletonInstance get singleton instance
// if created before, then return it directly, otherwise create it
func GetSingletonInstance() *Singleton {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		fmt.Println("[GetSingletonInstance] create singleton")
		instance = &Singleton{}
	}

	return instance
}
