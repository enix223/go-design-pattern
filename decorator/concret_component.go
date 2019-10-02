package decorator

import "fmt"

// ConcretComponent origin business related module
type ConcretComponent struct{}

// Operation implements the business operations
func (ConcretComponent) Operation() {
	fmt.Println("hello world")
}
