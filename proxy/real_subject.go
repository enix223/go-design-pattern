package proxy

import "fmt"

// RealSubject subject that perform some business related logic
type RealSubject struct{}

// Operation run business related operations
func (RealSubject) Operation() error {
	fmt.Println("hello world")
	return nil
}
