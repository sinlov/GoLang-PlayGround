package _0009_palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	// mock isPalindrome
	qs := []question{
		{
			param{9},
			answer{true},
		},
		{
			param{10},
			answer{false},
		},
		{
			param{-10},
			answer{false},
		},
		{
			param{-123},
			answer{false},
		},
		{
			param{313},
			answer{true},
		},
		{
			param{312},
			answer{false},
		},
		{
			param{1534276469},
			answer{false},
		},
	}
	t.Logf("~> LeetCode isPalindrome start")
	// do isPalindrome
	for _, q := range qs {
		a, p := q.answer, q.param
		//t.Logf("in [ %v ], out [%v]", p.one, isPalindrome(p.one))
		// verify isPalindrome
		assert.Equal(t, a.one, isPalindrome(p.one),
			"in [ %v ], out [%v]", p.one, isPalindrome(p.one))
	}
	t.Logf("~> LeetCode isPalindrome end")
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one int
}

// one first  answer
type answer struct {
	one bool
}
