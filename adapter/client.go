package adapter

import "fmt"

// Client adapter client
type Client struct {
	target Target
}

// Run client with old interface and new interface
func (c *Client) Run() {
	fmt.Println("calling client's target Request method...")
	c.target.Request()
}
