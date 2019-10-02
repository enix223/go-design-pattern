package bridge

import "fmt"

// ConcretImplmentorA concre implementor
type ConcretImplmentorA struct{}

// OperationImpl implement implementor
func (c *ConcretImplmentorA) OperationImpl() {
	fmt.Println("operation by concret implementor a")
}
