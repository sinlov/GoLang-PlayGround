package hex_code

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
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

func ByteArr2HexStrArr(bArr []byte) string {
	buf := new(bytes.Buffer)
	for _, b := range bArr {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buf.WriteString("0")
		}
		buf.WriteString(s)
		buf.WriteString(" ")
	}
	return strings.Trim(buf.String(), "")
}

func ByteArr2HexASCII(bArr []byte) string {
	buf := new(bytes.Buffer)
	count := 0
	line := 0
	buf.WriteString(fmt.Sprintf("%08d ", line))
	for _, b := range bArr {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buf.WriteString("0")
		}
		buf.WriteString(s)
		buf.WriteString(" ")
		count++
		if count%8 == 0 {
			buf.WriteString(" ")
		}
		if count%16 == 0 {
			buf.WriteString("\n")
			line++
			buf.WriteString(fmt.Sprintf("%08d ", line))
		}
	}
	return strings.Trim(buf.String(), "")
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
