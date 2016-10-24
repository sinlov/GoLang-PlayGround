package go_routine

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		a []int
		c chan int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		sum(tt.args.a, tt.args.c)
	}
}

func Test_go_sum(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name:"Base",
		},
	}
	for range tests {
		go_sum()
	}
}
