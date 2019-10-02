package composite

import "fmt"

// Composite the collection of the leaves and other composites
type Composite struct {
	value      int
	collection []Component
}

// NewComposite create composite
func NewComposite(v int) *Composite {
	return &Composite{
		value:      v,
		collection: make([]Component, 0),
	}
}

// Operation implements the component interface
func (c *Composite) Operation() {
	fmt.Printf("this is a composite (value = %d)\n", c.value)
	for _, component := range c.collection {
		component.Operation()
	}
}

// Add component to the collection
func (c *Composite) Add(component ...Component) {
	c.collection = append(c.collection, component...)
}

// Remove component from the collection
func (c *Composite) Remove(component Component) {
	idx := -1
	for i, item := range c.collection {
		if item == component {
			idx = i
			break
		}
	}

	if idx == -1 {
		panic("component not in the tree")
	}

	c.collection = append(c.collection[:idx], c.collection[idx+1:]...)
}

// GetChild get child from the tree
func (c *Composite) GetChild(idx int) Component {
	if idx >= len(c.collection) || idx < 0 {
		panic("idx out of range")
	}

	return c.collection[idx]
}
