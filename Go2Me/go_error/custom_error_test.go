package go_error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiyEqual(t *testing.T) {
	t.Logf("~> mock DiyEqual")
	// mock DiyEqual

	var testNum = 10
	t.Logf("~> do DiyEqual")
	// do DiyEqual
	result, err := Equal(testNum)
	// verify DiyEqual
	assert.Empty(t, err)
	assert.Equal(t, testNum, result)
}
