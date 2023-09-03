package go_routine

import "testing"

func Test_sumStartEnd(t *testing.T) {
	type args struct {
		start int
		end   int
		ch    chan string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		sumStartEnd(tt.args.start, tt.args.end, tt.args.ch)
	}
}

func Test_routinesSumStartEnd_no_buffer(t *testing.T) {
	type args struct {
		isBuff bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "no buffer Test",
			args: args{
				isBuff: false,
			},
		},
	}
	for _, tt := range tests {
		routinesSumStartEnd(tt.args.isBuff)
	}
}

func Test_routinesSumStartEnd(t *testing.T) {
	type args struct {
		isBuff bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "buffer test",
			args: args{
				isBuff: true,
			},
		},
	}
	for _, tt := range tests {
		routinesSumStartEnd(tt.args.isBuff)
	}
}
