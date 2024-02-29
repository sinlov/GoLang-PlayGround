package go_strings

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLastIndex(t *testing.T) {
	// mock LastIndex
	type args struct {
		target string
		sub    string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "go gopher",
			args: args{
				target: "go gopher",
				sub:    "go",
			},
			wantRes: 0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do LastIndex
			gotResult := strings.LastIndex(tc.args.target, tc.args.sub)

			// verify LastIndex
			assert.Equal(t, tc.wantRes, gotResult)
		})
	}
}
