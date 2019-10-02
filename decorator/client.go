package decorator

import "fmt"

// Client invoker that need to access Component interface
type Client struct{}

// Run decorate the existing ConcretComponent with preceding call function
// and succeeding call function
func (Client) Run() {
	c := &ConcretComponent{}
	d1 := PrecedingWrap(c, func() {
		fmt.Println("this is a preceding decorator function")
	})
	d1.Operation()

	fmt.Println()

	d2 := SucceedingWrap(c, func() {
		fmt.Println("this is a suceeding decorator function")
	})
	d2.Operation()
}
