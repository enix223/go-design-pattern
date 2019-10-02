package decorator

// WrapperFn wrapper function
type WrapperFn func()

// wrapper
type wrapper func()

// Operation implement the Component interface
func (w wrapper) Operation() {
	w()
}

// PrecedingWrap create a decorator to the Component
// fn is call before c Component's Operation
func PrecedingWrap(c Component, fn WrapperFn) Component {
	return wrapper(func() {
		fn()
		c.Operation()
	})
}

// SucceedingWrap create a decorator to the Component
// fn is call after c Component's Operation
func SucceedingWrap(c Component, fn WrapperFn) Component {
	return wrapper(func() {
		c.Operation()
		fn()
	})
}
