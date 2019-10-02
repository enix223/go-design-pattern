package facade

// Example decorator example
type Example struct{}

// Run decorator example
func (Example) Run() {
	c := &Client{}
	c.Run()
}
