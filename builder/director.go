package builder

import "fmt"

// Director constructs an object using the Builder interface
type Director struct{}

// Run build product through Builder
func (d *Director) Run() {
	b := NewBuilder()

	fmt.Println("Creating product...")

	p := b.SetName(
		"Product XXX",
	).SetSize(
		90,
	).BuildProduct()

	fmt.Printf("Product build: [%s]\n", p.Description())
}
