package go_reflect

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCanAddr(t *testing.T) {

	foo := FooStruct{}
	t.Logf("foo before %v", foo)

	fVal := reflect.ValueOf(foo)
	assert.False(t, fVal.CanAddr())
	assert.False(t, fVal.FieldByName("Pub").CanAddr())
	assert.False(t, fVal.FieldByName("Pub").CanSet())
	assert.False(t, fVal.FieldByName("prv").CanAddr())
	assert.False(t, fVal.FieldByName("prv").CanSet())

	pFVal := reflect.ValueOf(&foo)
	assert.False(t, pFVal.CanAddr())
	assert.True(t, pFVal.Elem().FieldByName("Pub").CanAddr())
	assert.True(t, pFVal.Elem().FieldByName("Pub").CanSet())
	pFVal.Elem().FieldByName("Pub").SetString("new foo name")
	assert.Equal(t, "new foo name", foo.Pub)
	assert.True(t, pFVal.Elem().FieldByName("prv").CanAddr())
	assert.False(t, pFVal.Elem().FieldByName("prv").CanSet())
	t.Logf("foo after %v", foo)

	bar := BarStruct{}
	t.Logf("bar before %v", bar)

	bVal := reflect.ValueOf(bar)
	assert.False(t, bVal.CanAddr())
	assert.False(t, bVal.FieldByName("Pub").CanAddr())
	assert.False(t, bVal.FieldByName("prv").CanAddr())
	assert.False(t, bVal.FieldByName("FooStruct").CanAddr())
	assert.False(t, bVal.FieldByName("bazStruct").CanAddr())
	assert.False(t, bVal.FieldByName("PFooStruct").CanAddr())
	assert.False(t, bVal.FieldByName("pBazStruct").CanAddr())

	pBVal := reflect.ValueOf(&bar)
	assert.False(t, pBVal.CanAddr())
	assert.True(t, pBVal.Elem().FieldByName("Pub").CanAddr())
	assert.True(t, pBVal.Elem().FieldByName("Pub").CanSet())
	pBVal.Elem().FieldByName("Pub").SetString("new bar name")
	assert.Equal(t, "new bar name", bar.Pub)
	assert.True(t, pBVal.Elem().FieldByName("prv").CanAddr())
	assert.False(t, pBVal.Elem().FieldByName("prv").CanSet())
	assert.True(t, pBVal.Elem().FieldByName("FooStruct").CanAddr())

	assert.True(t, pBVal.Elem().FieldByName("bazStruct").CanAddr())
	assert.True(t, pBVal.Elem().FieldByName("PFooStruct").CanAddr())
	assert.True(t, pBVal.Elem().FieldByName("pBazStruct").CanAddr())
}
