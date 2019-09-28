package abstractfactory

// ConcretProduct20 concret product 2
type ConcretProduct20 struct {
}

// Manufacturer return the name of manufacturer of the product
func (p *ConcretProduct20) Manufacturer() string {
	return "Manufacturer ABC"
}
