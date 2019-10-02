package bridge

// Abstraction abstraction interface for client to use
type Abstraction struct {
	impl Implementor
}

// Operation encapsulate the implementor operation
func (a *Abstraction) Operation() {
	a.impl.OperationImpl()
}

// SetImpl user can switch implementor at runtime
func (a *Abstraction) SetImpl(impl Implementor) {
	a.impl = impl
}
