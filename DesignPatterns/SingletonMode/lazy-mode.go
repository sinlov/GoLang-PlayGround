package SingletonMode

// Set a private variable as a singleton to be returned each time
var instanceLazy *lazy

// There are thread safety issues, it is possible to create multiple objects at high concurrency
func GetLazy() *lazy {
	if instanceLazy == nil {
		instanceLazy = new(lazy)
	}
	return instanceLazy
}
