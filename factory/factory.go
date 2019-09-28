package factory

// ProductFactory factory to make a product
type ProductFactory interface {
	MakeProduct() Product
}
