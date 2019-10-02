package facade

// Client a client which used to call the interface methods
// of different subsystems, from now on, it just need to
// reference to the facade, and use its simple interface
type Client struct{}

// Run call facade's simple interface
func (Client) Run() {
	f := NewFacade()
	f.Execute(1)
	f.Execute(2)
}
