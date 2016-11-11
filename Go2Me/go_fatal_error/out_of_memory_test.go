package go_fatal_error

import "testing"

func Test_listOutOfMem(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name:"OOM test",
		},
	}
	for range tests {
		listOutOfMem()
	}
}
