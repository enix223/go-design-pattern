package abstractfactory

// ProductFactory factory to create product
type ProductFactory interface {
	CreateProduct1() Product1
	CreateProduct2() Product2
}
