package go_reflect

import (
	"fmt"
	"reflect"
	"unsafe"
)

func GetStructPtrUnExportedField(source interface{}, fieldName string) reflect.Value {
	// 获取非导出字段反射对象
	v := reflect.ValueOf(source).Elem().FieldByName(fieldName)
	// 构建指向该字段的可寻址（addressable）反射对象
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func GetStructUnExportedField(source interface{}, fieldName string) (accessableField, addressableSourceCopy reflect.Value) {
	v := reflect.ValueOf(source)
	// since source is not a ptr, get an addressable copy of source to modify it later
	addressableSourceCopy = reflect.New(v.Type()).Elem()
	// make a copy of source
	addressableSourceCopy.Set(v)
	accessableField = addressableSourceCopy.FieldByName(fieldName)
	accessableField = reflect.NewAt(accessableField.Type(), unsafe.Pointer(accessableField.UnsafeAddr())).Elem()
	return
}

func SetStructPtrUnExportedStrField(source interface{}, fieldName string, fieldVal interface{}) (err error) {
	v := GetStructPtrUnExportedField(source, fieldName)
	rv := reflect.ValueOf(fieldVal)
	if v.Kind() != rv.Kind() {
		return fmt.Errorf("invalid kind: expected kind %v, got kind: %v", v.Kind(), rv.Kind())
	}
	v.Set(rv)
	return nil
}
