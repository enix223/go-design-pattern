package prototype

import "fmt"

// ConcretProtype1 concret prototype to implements an operation for cloning itself
type ConcretProtype1 struct {
	Name string
}

// NewConcretPrototype1 create concret prototype 1
func NewConcretPrototype1(n string) Prototype {
	return &ConcretProtype1{
		Name: n,
	}
}

// Clone implements an operation for cloning itself
func (c *ConcretProtype1) Clone() Prototype {
	clone := ConcretProtype1{}
	clone.Name = c.Name
	return &clone
}

func (c *ConcretProtype1) String() string {
	return fmt.Sprintf("[id: %p, name: %s]", c, c.Name)
}
