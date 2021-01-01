package _0007_Reverse_Integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	// mock reverse
	qs := []question{

		{
			param{321},
			answer{123},
		},

		{
			param{-123},
			answer{-321},
		},

		{
			param{120},
			answer{21},
		},

		{
			param{2147483648},
			answer{0},
		},
		{
			param{1534236469},
			answer{0},
		},
		{
			param{-2147483648},
			answer{0},
		},
	}
	t.Logf("~> LeetCode reverse start")
	// do reverse
	for _, q := range qs {
		a, p := q.answer, q.param
		//t.Logf("in [ %v ], out [%v]", p.one, reverse(p.one))
		// verify reverse
		assert.Equal(t, a.one, reverse(p.one), "in [ %v ], out [%v]", p.one, reverse(p.one))
	}
	t.Logf("~> LeetCode reverse end")
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
	one int
}
