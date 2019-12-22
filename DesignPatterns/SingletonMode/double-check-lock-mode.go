package SingletonMode

import "sync"

var instanceDoubleCheckLock *doubleCheckLock
var once sync.Once

// The callback is run only once during the running of the program
func GetInstanceDoubleCheckLock() *doubleCheckLock {
	once.Do(func() {
		instanceDoubleCheckLock = new(doubleCheckLock)
		instanceDoubleCheckLock.name = "first init"
	})
	return instanceDoubleCheckLock
}
