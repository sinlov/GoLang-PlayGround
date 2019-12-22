package SingletonMode

import "testing"

func TestGetInstanceDoubleCheckLock(t *testing.T) {
	lock1 := GetInstanceDoubleCheckLock()
	t.Logf("GetInstanceDoubleCheckLock first name: %v", lock1.name)
	lock2 := GetInstanceDoubleCheckLock()
	t.Logf("GetInstanceDoubleCheckLock second name: %v", lock2.name)
}
