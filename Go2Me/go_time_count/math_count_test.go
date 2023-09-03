package go_time_count

import "testing"

func TestMutilAdd(t *testing.T) {
	type args struct {
		start int
		count int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base test 0 - 10",
			args: args{
				start: 0,
				count: 10,
			},
		},
		{
			name: "base test 0 - 100",
			args: args{
				start: 0,
				count: 100,
			},
		},
	}
	for _, tt := range tests {
		MultiAdd(tt.args.start, tt.args.count)
	}
}
