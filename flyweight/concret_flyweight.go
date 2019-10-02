package flyweight

import "fmt"

// ConcretFlyweight an instance which can be reuse through the system
type ConcretFlyweight struct{}

// Operation business logic
func (c *ConcretFlyweight) Operation() {
	fmt.Printf("flyweight object %p\n", c)
}
