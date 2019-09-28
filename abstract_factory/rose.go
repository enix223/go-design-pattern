package abstractfactory

// Rose rose flower
type Rose struct {
}

// Color returns the color of the flower
func (f *Rose) Color() string {
	return "red"
}
