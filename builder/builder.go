package builder

// Builder interface for building product object
type Builder interface {
	SetName(string) Builder
	SetWeight(int) Builder
	SetSize(int) Builder
	SetManufacturer(string) Builder
	BuildProduct() *Product
}
