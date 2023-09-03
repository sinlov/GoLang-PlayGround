package go_map

import (
	"fmt"
	"reflect"
	"sort"
)

// Sort the map key (int \ float \ string) to traverse the map
//
//	sortMap -> The map to be sorted
//	sortFunc -> Map traversal received, the input parameters should match the key and value of the map. like func(k int, val string)
//	return nil is success, other sort error
func SortMapByKey(sortMap interface{}, sortFunc interface{}) error {
	eachMapValue := reflect.ValueOf(sortMap)
	eachFuncValue := reflect.ValueOf(sortFunc)
	eachMapType := eachMapValue.Type()
	eachFuncType := eachFuncValue.Type()
	if eachMapValue.Kind() != reflect.Map {
		return fmt.Errorf(`SortMap error. parameter "sortMap" first type must is map[...]...{}`)
	}
	if eachFuncValue.Kind() != reflect.Func {
		return fmt.Errorf(`SortMap error. parameter "sortFunc" second type must is func(key ..., value ...)`)
	}
	if eachFuncType.NumIn() != 2 {
		return fmt.Errorf(`SortMap error. "sortFunc" input parameter count must is 2`)
	}
	if eachFuncType.In(0).Kind() != eachMapType.Key().Kind() {
		return fmt.Errorf(`SortMap error. "sortFunc" input parameter 1 type not equal of "sortMap" key, want %v`, eachMapType.Key().Kind())
	}
	if eachFuncType.In(1).Kind() != eachMapType.Elem().Kind() {
		return fmt.Errorf(`SortMap error. "sortFunc" input parameter 2 type not equal of "sortMap" value, want %v`, eachMapType.Elem().Kind())
	}

	// sort key
	// Get the sorted map key and value, and call sortFunc as a parameter
	switch eachMapType.Key().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		keys := make([]int, 0)
		keysMap := map[int]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, int(value.Int()))
			keysMap[int(value.Int())] = value
		}
		sort.Ints(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.Float64, reflect.Float32:
		keys := make([]float64, 0)
		keysMap := map[float64]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, float64(value.Float()))
			keysMap[float64(value.Float())] = value
		}
		sort.Float64s(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.String:
		keys := make([]string, 0)
		keysMap := map[string]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, value.String())
			keysMap[value.String()] = value
		}
		sort.Strings(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	default:
		return fmt.Errorf(`"SortMap" key type must is int or float or string`)
	}
	return nil
}
