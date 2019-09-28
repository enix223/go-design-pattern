package singleton

import (
	"fmt"
	"sync"
)

// Example singleton example
type Example struct{}

// Run the singleton example
func (e *Example) Run() {
	// Singleton pattern implemented by package variables
	fmt.Println("==== Singleton implementation with package variable ====")
	fmt.Println("Creating two goroutines to get singleton instance...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		s := GetSingletonInstance()
		fmt.Printf("[GOROUTINE #1] singleton instance: %p, he say: %s\n", s, s.SayHello())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		s := GetSingletonInstance()
		fmt.Printf("[GOROUTINE #2] singleton instance: %p, he say: %s\n", s, s.SayHello())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done")

	// Singleton pattern implemented by sync.Once
	fmt.Println("==== Singleton implementation with sync.Once ====")
	getInstance := BuildSingletonCreator()

	wg.Add(1)
	go func() {
		s := getInstance()
		fmt.Printf("[GOROUTINE #1] singleton instance: %p, he say: %s\n", s, s.SayHello())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		s := getInstance()
		fmt.Printf("[GOROUTINE #2] singleton instance: %p, he say: %s\n", s, s.SayHello())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done")
}
