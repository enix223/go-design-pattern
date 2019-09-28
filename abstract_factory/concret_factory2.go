package abstractfactory

// ConcretFactory2 factory to create concret product 11 & 20
type ConcretFactory2 struct {
}

// CreateProduct1 create product 1
func (f *ConcretFactory2) CreateProduct1() Product1 {
	return &ConcretProduct11{}
}

// CreateProduct2 create product 2
func (f *ConcretFactory2) CreateProduct2() Product2 {
	return &ConcretProduct20{}
}
