package go_base

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestStrconvParseInt(t *testing.T) {
	// mock StrconvParseInt
	i, err := strconv.ParseInt("0", 10, 0)
	if err != nil {
		t.Error(err)
	}
	assert.LessOrEqual(t, int64(0), i)
	_, err = strconv.ParseInt("", 10, 0)
	if err == nil {
		t.Error("if string \"\" will not cover")
	}

	var uNum uint
	uNum = 10
	assert.Equal(t, "10", strconv.Itoa(int(uNum)))
}
