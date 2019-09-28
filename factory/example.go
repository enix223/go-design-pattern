package factory

import "fmt"

// Example factory method pattern example
type Example struct{}

// ProductFactoryMethod method to create a product
type ProductFactoryMethod func() Product

// MakeProduct implement product factory
func (f ProductFactoryMethod) MakeProduct() Product {
	return f()
}

// Run the factory pattern example
func (e *Example) Run() {
	fmt.Println("Creating client 1...")
	c1 := &Client{
		ProductFactory: ProductFactoryMethod(func() Product {
			// factory methond to create product1
			return &ConcretProduct1{
				Name: "model TERMINATOR",
			}
		}),
	}
	c1.Run()
	fmt.Println("Creating client 2...")
	c2 := &Client{
		ProductFactory: ProductFactoryMethod(func() Product {
			// factory methond to create product2
			return &ConcretProduct2{
				Name: "model X",
			}
		}),
	}
	c2.Run()
}
