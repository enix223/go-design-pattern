package factory

import "fmt"

// ConcretProduct1 a concret product
type ConcretProduct1 struct {
	Name string
}

// Description return the description of the product
func (c *ConcretProduct1) Description() string {
	return fmt.Sprintf("[Manufacture by XXX] product name: %s", c.Name)
}
