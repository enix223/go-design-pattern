package bridge

// Example bridge example
type Example struct{}

// Run bridge example
func (Example) Run() {
	c := NewClient()

	c.Run()
}
