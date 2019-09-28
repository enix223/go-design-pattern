package prototype

import "fmt"

// ConcretProtype2 concret prototype to implements an operation for cloning itself
type ConcretProtype2 struct {
	Description string
}

// NewConcretPrototype2 create concret prototype 2
func NewConcretPrototype2(d string) Prototype {
	return &ConcretProtype2{
		Description: d,
	}
}

// Clone implements an operation for cloning itself
func (c *ConcretProtype2) Clone() Prototype {
	clone := ConcretProtype2{}
	clone.Description = c.Description
	return &clone
}

func (c *ConcretProtype2) String() string {
	return fmt.Sprintf("[id: %p, description: %s]", c, c.Description)
}
