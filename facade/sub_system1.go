package facade

import "fmt"

// SubSystem1 a complicated subsystem which need to apply facade pattern
// to simplify the client's usage
type SubSystem1 struct{}

// Operation1 operation provided by sub system 1
func (SubSystem1) Operation1() {
	fmt.Println("this is an operation from subsystem 1")
}
