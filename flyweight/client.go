package flyweight

// Client which need to use the same flywegith object through the application
type Client struct{}

// Run flyweight operation twice
func (Client) Run() {
	f1 := GetFlyweight("1")
	f1.Operation()

	f2 := GetFlyweight("1")
	f2.Operation()
}
