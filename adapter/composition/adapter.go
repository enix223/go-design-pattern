package composition

import (
	"github.com/enix223/go-design-pattern/adapter/adaptee"
)

// Adapter using compoisition style to implement adapter pattern
type Adapter struct {
	adaptee adaptee.Adaptee
}

// NewAdapter create an adapter
func NewAdapter() *Adapter {
	return &Adapter{
		adaptee: &ConcretAdaptee{},
	}
}

// Request implement the target interface
func (a *Adapter) Request() {
	a.adaptee.SpecificRequest()
}
