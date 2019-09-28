package abstractfactory

// ConcretProduct11 concret product 1
type ConcretProduct11 struct {
}

// Name return the name of the product
func (p *ConcretProduct11) Name() string {
	return "concret product 1-1"
}
