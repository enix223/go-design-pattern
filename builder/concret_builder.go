package builder

// ConcretBuilder concret builder to build product object
type ConcretBuilder struct {
	// should not export
	product *Product
}

// NewBuilder create a builder
func NewBuilder() Builder {
	c := &ConcretBuilder{
		product: &Product{},
	}
	return c
}

// BuildProduct return the final product instance
func (c *ConcretBuilder) BuildProduct() *Product {
	// apply default value if user not set
	if c.product.Name == "" {
		c.product.Name = "default name"
	}
	if c.product.Weight == 0 {
		c.product.Weight = 15
	}
	if c.product.Manufacturer == "" {
		c.product.Manufacturer = "default manufacturer"
	}
	if c.product.Size == 0 {
		c.product.Size = 20
	}
	return c.product
}

// SetName set the name of the product
func (c *ConcretBuilder) SetName(name string) Builder {
	c.product.Name = name
	return c
}

// SetWeight set the weight of the product
func (c *ConcretBuilder) SetWeight(weight int) Builder {
	c.product.Weight = weight
	return c
}

// SetSize set the size of the product
func (c *ConcretBuilder) SetSize(size int) Builder {
	c.product.Size = size
	return c
}

// SetManufacturer set the manufacturer of the product
func (c *ConcretBuilder) SetManufacturer(m string) Builder {
	c.product.Manufacturer = m
	return c
}
