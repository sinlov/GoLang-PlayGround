package go_strings

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHasSuffix(t *testing.T) {
	// mock HasSuffix
	type args struct {
		target string
		suffix string
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		{
			name: "sample",
			args: args{
				target: "Amigo",
				suffix: "go",
			},
			wantRes: true,
		},
		{
			name: "two",
			args: args{
				target: "Ami.go",
				suffix: ".go",
			},
			wantRes: true,
		},
		{
			name: "three",
			args: args{
				target: "foo/one/Ami.go",
				suffix: "one/Ami.go",
			},
			wantRes: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do HasSuffix
			gotResult := strings.HasSuffix(tc.args.target, tc.args.suffix)

			// verify HasSuffix
			assert.Equal(t, tc.wantRes, gotResult)
		})
	}
}
