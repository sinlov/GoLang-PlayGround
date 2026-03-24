package doublestar

import (
	"testing"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/stretchr/testify/assert"
)

func TestDoublestarMatchKeys(t *testing.T) {
	// mock DoublestarMatchKeys
	type args struct {
		Include []string
		Exclude []string
	}
	tests := []struct {
		name            string
		args            args
		target          string
		wantResIncludes bool
		wantResExcludes bool
	}{
		{
			name: "basic",
			args: args{
				Include: []string{
					"main",
				},
				Exclude: []string{},
			},
			target:          "main",
			wantResIncludes: true,
			wantResExcludes: false,
		},
		{
			name: "config",
			args: args{
				Include: []string{
					"main",
					"renovate/*",
					"*-feature-*",
				},
				Exclude: []string{},
			},
			target:          "10-feature-test-version-update",
			wantResIncludes: true,
			wantResExcludes: false,
		},
		{
			name: "config",
			args: args{
				Include: []string{
					"main",
					"renovate/*",
					"*-feature-*",
				},
				Exclude: []string{},
			},
			target:          "feature-test-version-update",
			wantResIncludes: false,
			wantResExcludes: false,
		},
		{
			name: "config",
			args: args{
				Include: []string{
					"main",
					"renovate/*",
					"**",
				},
				Exclude: []string{},
			},
			target:          "feature-test-version-update\n",
			wantResIncludes: true,
			wantResExcludes: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do DoublestarMatchKeys
			var resIncludes bool
			for _, pattern := range tc.args.Include {
				if ok, _ := doublestar.Match(pattern, tc.target); ok {
					resIncludes = true
				}
			}

			var resExcludes bool
			for _, pattern := range tc.args.Exclude {
				if ok, _ := doublestar.Match(pattern, tc.target); ok {
					resExcludes = true
				}
			}

			// verify DoublestarMatchKeys
			assert.Equal(t, tc.wantResIncludes, resIncludes)
			assert.Equal(t, tc.wantResExcludes, resExcludes)
		})
	}
}

// more info see: https://github.com/bmatcuk/doublestar/?tab=readme-ov-file#patterns
func TestDoublestarMatch(t *testing.T) {
	// mock DoublestarMatch
	type args struct {
		pattern string
		target  string
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
		wantErr error
	}{
		{
			name: "sample",
			args: args{
				pattern: "a",
				target:  "a/b/c",
			},
			wantRes: false,
		},
		{
			name: "branch origin/main",
			args: args{
				pattern: "main",
				target:  "origin/main",
			},
			wantRes: false,
		},
		{
			name: "branch main",
			args: args{
				pattern: "main",
				target:  "main",
			},
			wantRes: true,
		},
		{
			name: "branch any sequence of non-path-separators",
			args: args{
				pattern: "*",
				target:  "feature-test-version-update",
			},
			wantRes: true,
		},
		{
			name: "branch feature try",
			args: args{
				pattern: "*-feature-*",
				target:  "10-feature-test-version-update",
			},
			wantRes: true,
		},
		{
			name: "branch feature try",
			args: args{
				pattern: "*-feature-*",
				target:  "feature-test-version-update",
			},
			wantRes: false,
		},
		{
			name: "branch feature try",
			args: args{
				pattern: "renovate/*",
				target:  "feature-test-version-update",
			},
			wantRes: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do DoublestarMatch
			gotResult, gotErr := doublestar.Match(tc.args.pattern, tc.args.target)

			// verify DoublestarMatch
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			assert.Equal(t, tc.wantRes, gotResult)
		})
	}
}
