package abstractfactory

// ConcretProduct21 concret product 2
type ConcretProduct21 struct {
}

// Manufacturer return the name of manufacturer of the product
func (p *ConcretProduct21) Manufacturer() string {
	return "Manufacturer ZZZ"
}
