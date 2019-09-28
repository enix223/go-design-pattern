package builder

import "fmt"

// Product represents the complex object under construction
type Product struct {
	Name         string
	Manufacturer string
	Weight       int
	Size         int
}

// Description return the description of the product
func (p *Product) Description() string {
	return fmt.Sprintf(
		"name: %s, manufacturer: %s, weight: %d, size: %d",
		p.Name,
		p.Manufacturer,
		p.Weight,
		p.Size,
	)
}
