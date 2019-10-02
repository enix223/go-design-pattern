package facade

import (
	"fmt"
)

// Facade a high level wrapper of the subsystems
// to provide simple interface to the client
type Facade struct {
	sub1 *SubSystem1
	sub2 *SubSystem2
}

// NewFacade create a new facade
func NewFacade() *Facade {
	return &Facade{
		sub1: &SubSystem1{},
		sub2: &SubSystem2{},
	}
}

// Execute a simple interface which provide to client
// to hide the detail usage of the subsystem interfaces
func (f *Facade) Execute(typ int) {
	if typ == 1 {
		f.sub1.Operation1()
	} else if typ == 2 {
		f.sub2.Operation2()
	} else {
		panic(fmt.Sprintf("invalid type: %d", typ))
	}
}
