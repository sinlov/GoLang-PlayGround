package SingletonMode

var instanceStarving *starving

// init
func init() {
	instanceStarving = new(starving)
	instanceStarving.name = "first init"
}

// GetStarving
// Compared with lazy mode, it is more secure,
// but it will slow down the startup of the program.
func GetStarving() *starving {
	return instanceStarving
}
