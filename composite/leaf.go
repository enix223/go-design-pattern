package composite

import "fmt"

// Leaf the leaf of the tree
type Leaf struct {
	value int
}

// NewLeaf create a new leaf
func NewLeaf(v int) *Leaf {
	return &Leaf{
		value: v,
	}
}

// Operation implements the component interface
func (l *Leaf) Operation() {
	fmt.Printf("this is a leaf (value = %d)\n", l.value)
}
