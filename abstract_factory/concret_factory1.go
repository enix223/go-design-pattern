package abstractfactory

// ConcretFactory1 factory to create concret product 10 & 21
type ConcretFactory1 struct {
}

// CreateProduct1 create product 1
func (f *ConcretFactory1) CreateProduct1() Product1 {
	return &ConcretProduct10{}
}

// CreateProduct2 create product 2
func (f *ConcretFactory1) CreateProduct2() Product2 {
	return &ConcretProduct21{}
}
