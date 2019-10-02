package facade

import "fmt"

// SubSystem2 a complicated subsystem which need to apply facade pattern
// to simplify the client's usage
type SubSystem2 struct{}

// Operation2 operation provided by sub system 2
func (SubSystem2) Operation2() {
	fmt.Println("this is an operation from subsystem 2")
}
