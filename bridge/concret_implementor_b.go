package bridge

import "fmt"

// ConcretImplmentorB concre implementor
type ConcretImplmentorB struct{}

// OperationImpl implement implementor
func (c *ConcretImplmentorB) OperationImpl() {
	fmt.Println("operation by concret implementor b")
}
