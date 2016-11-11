package go_fatal_error

import "testing"

func Test_recursiveOverFlows(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name:"recursive",
		},
	}
	for range tests {
		recursiveOverFlows()
	}
}
