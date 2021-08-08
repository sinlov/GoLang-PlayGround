package hex_code

import (
	"bytes"
	"strconv"
)

func ByteArr2HexStr(bArr []byte) string {
	buf := new(bytes.Buffer)
	for _, b := range bArr {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buf.WriteString("0")
		}
		buf.WriteString(s)
	}
	return buf.String()
}

func HexStr2ByteArr(hexString string) ([]byte, error) {
	length := len(hexString) / 2
	slice := make([]byte, length)
	rs := []rune(hexString)
	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, err := strconv.ParseInt(s, 16, 10)
		if err != nil {
			return nil, err
		}
		slice[i] = byte(value & 0xFF)
	}
	return slice, nil
}