package go_base

// go support int and uint, these length can be same with difference code format
// go has define number format
// rune int8 int16 int32 int64 -- rune is short of int32
// byte uint8 uint16 uint32 uint64 -- byte is short of unit8

// Warning! You can not cast different number format such as below

func err_number_format() {
	var oneCode int8 = 10
	var twoCode int16 = 100
	// this will be error
	threeCode := oneCode + twoCode
	print(threeCode)
}
