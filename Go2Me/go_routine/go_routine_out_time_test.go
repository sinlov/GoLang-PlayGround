package go_routine

import "testing"

func Test_try_count_sum_out(t *testing.T) {
	type args struct {
		is_time_out bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name:"not time out",
			args:args{
				is_time_out:false,
			},
		},
		{
			name:"time out",
			args:args{
				is_time_out:true,
			},
		},
	}
	for _, tt := range tests {
		try_count_sum_out(tt.args.is_time_out)
	}
}
