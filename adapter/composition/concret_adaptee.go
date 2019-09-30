package composition

import "fmt"

// ConcretAdaptee struct that implemented the new interface
type ConcretAdaptee struct{}

// SpecificRequest implement the new interface
func (c *ConcretAdaptee) SpecificRequest() {
	fmt.Println("New interface specific request")
}
