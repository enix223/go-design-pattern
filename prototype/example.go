package prototype

import "fmt"

// Example for prototype pattern
type Example struct{}

// Run prototype pattern example, to clone objects
func (e *Example) Run() {
	fmt.Println("Creating client with ConcretPrototype1...")
	c1 := &Client{
		prototype: NewConcretPrototype1("concret prototype 1"),
	}
	c1.Run()

	fmt.Println()

	fmt.Println("Creating client with ConcretPrototype2...")
	c2 := &Client{
		prototype: NewConcretPrototype2("concret prototype 2"),
	}
	c2.Run()
}
