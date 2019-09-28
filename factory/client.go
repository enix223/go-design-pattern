package factory

import (
	"fmt"
)

// Client factory user
type Client struct {
	// since golang do not support inheritance, we need to inject
	// a concret factory method when initalizing the object
	ProductFactory
}

// Run client example
func (c *Client) Run() {
	fmt.Println("making product...")
	p := c.MakeProduct()
	fmt.Printf("Make product success: %s\n", p.Description())
}
