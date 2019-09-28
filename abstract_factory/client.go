package abstractfactory

import "fmt"

// Client abstract factory user
type Client struct {
	// factory to create Product1 & Product2
	factory ProductFactory
}

// Run run client
func (c *Client) Run() {
	fmt.Println("Creating product1...")
	p1 := c.factory.CreateProduct1()
	fmt.Println("Creating product2...")
	p2 := c.factory.CreateProduct2()

	fmt.Printf("Name of product1: %s\n", p1.Name())
	fmt.Printf("Manufacturer of product2: %s\n", p2.Manufacturer())
}
