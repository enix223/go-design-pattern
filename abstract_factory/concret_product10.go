package abstractfactory

// ConcretProduct10 concret product 1
type ConcretProduct10 struct {
}

// Name return the name of the product
func (p *ConcretProduct10) Name() string {
	return "concret product 1-0"
}
