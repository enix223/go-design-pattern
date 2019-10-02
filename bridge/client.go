package bridge

import "fmt"

// Client that use the abstraction interface
type Client struct {
	abstraction *Abstraction
}

// NewClient create client
func NewClient() *Client {
	return &Client{
		abstraction: &Abstraction{},
	}
}

// Run client example
func (c *Client) Run() {
	// set implementor to ConcretImplementorA
	fmt.Println("======================================")
	fmt.Println("Run client with concret implementor a")
	fmt.Println("======================================")
	c.abstraction.SetImpl(&ConcretImplmentorA{})
	c.abstraction.Operation()

	fmt.Println()

	// switch implementor to ConcretImplementorB
	fmt.Println("======================================")
	fmt.Println("Run client with concret implementor b")
	fmt.Println("======================================")
	c.abstraction.SetImpl(&ConcretImplmentorB{})
	c.abstraction.Operation()
}
