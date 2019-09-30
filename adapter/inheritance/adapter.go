package inheritance

import "fmt"

// Adapter implement both old interface and new interface
// forward the old interface method to the new interface method
type Adapter struct{}

// Request implement the old interface
func (a *Adapter) Request() {
	a.SpecificRequest()
}

// SpecificRequest implement new interface
func (a *Adapter) SpecificRequest() {
	fmt.Println("New interface specific request")
}
