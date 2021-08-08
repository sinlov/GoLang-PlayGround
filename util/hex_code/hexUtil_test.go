package hex_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteArr2HexStr(t *testing.T) {
	t.Logf("\nTestByteArr2HexStr")
	str4Hex := "12345qwert"
	t.Logf("str4Hex %s", str4Hex)
	hexStr := []byte(str4Hex)
	t.Logf("hexStr: %s", hexStr)
	hexString := ByteArr2HexStr(hexStr)
	t.Logf("ByteArr2HexString %s", hexString)
	assert.Equal(t, "31323334357177657274", hexString)
}

func TestHexStr2ByteArr(t *testing.T) {
	t.Logf("\nTestHexStr2ByteArr")
	str4Hex := "12345qwert"
	t.Logf("str4Hex: %s", str4Hex)
	hexStr := []byte(str4Hex)
	t.Logf("hexStr: %s", hexStr)
	hexString := ByteArr2HexStr(hexStr)
	t.Logf("ByteArr2HexStr: %s", hexString)
	byteArr, err := HexStr2ByteArr(hexString)
	t.Logf("byteArr: %s", byteArr)
	assert.Equal(t, hexStr, byteArr, err)
}
