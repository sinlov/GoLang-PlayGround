package go_map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortMapByKey(t *testing.T) {
	// mock SortMapByKey
	t.Logf("~> mock SortMapByKey")
	intMap := map[int]string{}
	intMap[5] = "Hello World"
	intMap[2] = "Hello World"
	intMap[1] = "Hello World"
	intMap[3] = "Hello World"
	intMap[4] = "Hello World"

	// do SortMapByKey
	err := SortMapByKey(intMap, func(key int, value string) {
		t.Logf("intMap key:value => [ %v : %v ]", key, value)
	})
	t.Logf("~> do SortMapByKey")
	// verify SortMapByKey
	assert.Nil(t, err)

	floatMap := map[float32]string{}
	floatMap[0.5] = "0.5"
	floatMap[0.1] = "0.1"
	floatMap[0.3] = "0.3"
	floatMap[0.2] = "0.2"
	floatMap[0.4] = "0.4"

	err = SortMapByKey(floatMap, func(k float32, val string) {
		t.Logf("floatMap key:value => [ %v : %v ]", k, val)
	})
	assert.Nil(t, err)

	stringMap := map[string]string{}
	stringMap["xyz"] = "more info"
	stringMap["abc"] = "more info"
	stringMap["500"] = "more info"
	stringMap["cfg"] = "more info"

	err = SortMapByKey(stringMap, func(k string, val string) {
		t.Logf("stringMap key:value => [ %v : %v ]", k, val)
	})
	assert.Nil(t, err)

}

func BenchmarkSortMapByKey(b *testing.B) {
	m := map[int]string{}
	m[5] = "Hello World"
	m[2] = "Hello World"
	m[1] = "Hello World"
	m[3] = "Hello World"
	m[4] = "Hello World"
	for i := 0; i < b.N; i++ {
		err := SortMapByKey(m, func(key int, value string) {
			//b.Logf("key:value => [ %v : %v ]", key, value)
		})
		assert.Nil(b, err)
	}
}

func TestSortMapByKeyErr(t *testing.T) {
	// mock SortMapByKeyErr
	m := map[int]string{}
	m[5] = "Hello World"
	m[2] = "Hello World"
	m[1] = "Hello World"
	n := "error map type as string"
	t.Logf("~> mock SortMapByKeyErr")
	errFirstParam := SortMapByKey(n, func(key int, value string) {
	})
	// do SortMapByKeyErr
	t.Logf("~> do SortMapByKeyErr")
	// verify SortMapByKeyErr
	if errFirstParam != nil {
		t.Logf("errFirstParam %v", errFirstParam)
	}
	assert.NotNil(t, errFirstParam)

	errSecondParamSize := SortMapByKey(m, func(key int) {
	})
	if errSecondParamSize != nil {
		t.Logf("errSecondParamSize %v", errSecondParamSize)
	}
	assert.NotNil(t, errSecondParamSize)

	errSecondParamSizeThree := SortMapByKey(m, func(key int, val string, other int) {
	})
	if errSecondParamSizeThree != nil {
		t.Logf("errSecondParamSizeThree %v", errSecondParamSizeThree)
	}
	assert.NotNil(t, errSecondParamSizeThree)

	errSecondParamKeyName := SortMapByKey(m, func(key string, val string) {
		t.Logf("key:errSecondParamKeyName => [ %v : %v ]", key, val)
	})
	if errSecondParamKeyName != nil {
		t.Logf("errSecondParamKeyName %v", errSecondParamKeyName)
	}
	assert.NotNil(t, errSecondParamKeyName)

	errSecondParamKeyValue := SortMapByKey(m, func(key int, val int64) {
		t.Logf("key:errSecondParamKeyValue => [ %v : %v ]", key, val)
	})
	if errSecondParamKeyValue != nil {
		t.Logf("errSecondParamKeyValue %v", errSecondParamKeyValue)
	}
	assert.NotNil(t, errSecondParamKeyValue)
}
