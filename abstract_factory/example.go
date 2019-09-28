package abstractfactory

import "fmt"

// Example of abstract factory pattern
type Example struct{}

// Run abstract factory example
func (e *Example) Run() {
	fmt.Println("Creating client with ConcretFactory1...")
	c1 := Client{
		factory: &ConcretFactory1{},
	}
	c1.Run()

	fmt.Println()

	fmt.Println("Creating client with ConcretFactory2...")
	c2 := Client{
		factory: &ConcretFactory2{},
	}
	c2.Run()
}
