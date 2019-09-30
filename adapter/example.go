package adapter

import (
	"fmt"

	"github.com/enix223/go-design-pattern/adapter/composition"
	"github.com/enix223/go-design-pattern/adapter/inheritance"
)

// Example adapter example
type Example struct{}

// Run dapter example
func (e *Example) Run() {
	fmt.Println("======================================")
	fmt.Println("Adapter implemented with inheritance")
	fmt.Println("======================================")
	fmt.Println("Creating client with origin target...")
	c1 := &Client{
		target: &ConcretTarget{},
	}
	c1.Run()

	fmt.Println("Creating client with new adaper...")
	c2 := &Client{
		target: &inheritance.Adapter{},
	}
	c2.Run()

	fmt.Println()

	fmt.Println("======================================")
	fmt.Println("Adapter implemented with composition")
	fmt.Println("======================================")
	fmt.Println("Creating client with origin target...")
	c3 := &Client{
		target: &ConcretTarget{},
	}
	c3.Run()

	fmt.Println("Creating client with new adaper...")
	c4 := &Client{
		target: composition.NewAdapter(),
	}
	c4.Run()
}
