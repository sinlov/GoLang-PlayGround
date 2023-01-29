package security_test

import (
	"fmt"
	"github.com/sinlov/GoLang-PlayGround/util/security"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURLPathEncode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "URLPathEncode: 你好",
			source: `你好 我在`,
			want:   `%E4%BD%A0%E5%A5%BD%20%E6%88%91%E5%9C%A8`,
		},
	}
	// do
	for _, test := range tests {
		result := security.URLPathEncode(test.source)
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)
	}
}

func TestURLPathDecode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "URLPathDecode: 你好",
			source: `%E4%BD%A0%E5%A5%BD%20%E6%88%91%E5%9C%A8`,
			want:   `你好 我在`,
		},
	}
	// do
	for _, test := range tests {
		result, err := security.URLPathDecode(test.source)
		if err != nil {
			t.Fatalf("TestURLPathDecode error: %v", err)
		}
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)
	}
}

func TestURLQueryEncode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "URLQueryEncode: 你好",
			source: `你好 我在`,
			want:   `%E4%BD%A0%E5%A5%BD+%E6%88%91%E5%9C%A8`,
		},
	}
	// do
	for _, test := range tests {
		result := security.URLQueryEncode(test.source)
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)
	}
}

func TestURLQueryDecode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "URLQueryDecode: 你好",
			source: `%E4%BD%A0%E5%A5%BD%20%E6%88%91%E5%9C%A8`,
			want:   `你好 我在`,
		},
	}
	// do
	for _, test := range tests {
		result, err := security.URLQueryDecode(test.source)
		if err != nil {
			t.Fatalf("TestURLPathDecode error: %v", err)
		}
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)

	}
}

func TestUnicodeDecode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
		err    error
	}{
		{
			name:   "UnicodeDecode: 你好",
			source: `\u4f60\u597d\u0020\u6211\u5728`,
			want:   `你好 我在`,
		},

		{
			name:   "UnicodeDecode: error",
			source: `4f60597d002062115728`,
			want:   `你好 我在`,
			err:    fmt.Errorf("input unicode error, string must has %v", "\\u"),
		},
	}
	// do
	for _, test := range tests {
		result, err := security.UnicodeDecode(test.source)
		// verify
		if err != nil {
			t.Logf("TestUnicodeDecode error: %v", err)
			assert.Equal(t, test.err, err)
		} else {
			t.Logf("source: %v, result: %v", test.source, result)
			assert.Equal(t, test.want, result)
		}
	}
}
