package singleton

import (
	"fmt"
	"sync"
)

// GetInstance function to return a singleton instance
type GetInstance func() *Singleton

// BuildSingletonCreator create a function that return singleton object
func BuildSingletonCreator() GetInstance {
	var once sync.Once
	var obj *Singleton
	return func() *Singleton {
		once.Do(func() {
			fmt.Println("[BuildSingletonCreator] create singleton")
			obj = &Singleton{}
		})
		return obj
	}
}
