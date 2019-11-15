package go_byte

import "unsafe"

// more fast than IsBigEndian
func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return b == 0x04
}

func IsBigEndian() bool {
	s := int16(0x1234)
	b := int8(s)
	return b != 0x34
}
