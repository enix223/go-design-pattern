package adapter

import "fmt"

// ConcretTarget implement target interface
type ConcretTarget struct{}

// Request implement target interface
func (c *ConcretTarget) Request() {
	fmt.Println("Origin request method")
}
