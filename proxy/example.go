package proxy

import "fmt"

// Example proxy example
type Example struct{}

// Run proxy example
func (Example) Run() {
	fmt.Println("=================")
	fmt.Println("Run without proxy")
	fmt.Println("=================")
	c1 := NewClientWithSubject()
	c1.Run()

	fmt.Println()

	fmt.Println("=======================")
	fmt.Println("Run with readable proxy")
	fmt.Println("=======================")
	c2 := NewClientWithReadableProxy()
	c2.Run()

	fmt.Println()

	fmt.Println("===========================")
	fmt.Println("Run with non readable proxy")
	fmt.Println("===========================")
	c3 := NewClientWithNonReadableProxy()
	c3.Run()
}
