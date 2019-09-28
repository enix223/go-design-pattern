package prototype

import "fmt"

// Client to use prototype to clone objects
type Client struct {
	prototype Prototype
}

// Run clone objects
func (c *Client) Run() {
	copy := c.prototype.Clone()
	fmt.Printf("old object: %s, cloned object: %s\n", c.prototype, copy)
}
