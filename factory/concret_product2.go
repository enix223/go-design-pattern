package factory

import "fmt"

// ConcretProduct2 a concret product
type ConcretProduct2 struct {
	Name string
}

// Description return the description of the product
func (c *ConcretProduct2) Description() string {
	return fmt.Sprintf("[Manufacture by ZZZ] Product name: %s", c.Name)
}
