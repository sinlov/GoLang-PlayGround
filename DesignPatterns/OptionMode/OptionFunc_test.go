package OptionMode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOption(t *testing.T) {
	// mock
	absOptions := newOptionABS("A", "B", 100)
	optionFast := NewOption()
	changeOption := NewOption(
		WithA("B"),
		WithB("A"),
	)
	// do
	t.Logf("absOptions %v", absOptions)
	t.Logf("optionFast %v", optionFast)
	t.Logf("changeOption %v", changeOption)

	assert.Equal(t, absOptions.A, optionFast.A)
	assert.Equal(t, absOptions.B, optionFast.B)
	assert.Equal(t, absOptions.C, optionFast.C)
	assert.Equal(t, changeOption.A, "B")
	assert.Equal(t, changeOption.B, "A")
	assert.Equal(t, changeOption.C, 100)
}
