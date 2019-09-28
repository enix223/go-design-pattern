package builder

import "fmt"

// Example of abstract factory pattern
type Example struct{}

// Run abstract factory example
func (e *Example) Run() {
	fmt.Println("Creating directory...")
	d := Director{}
	d.Run()
}
