package hex_code

import (
	"bytes"
	"errors"
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

func HexStr2ByteArrUnSafe(hexString string) []byte {
	byteArr, err := HexStr2ByteArr(hexString)
	if err != nil {
		return nil
	}
	return byteArr
}

func BlockCopy(src []byte, srcOffset int, dst []byte, dstOffset, count int) (bool, error) {
	srcLen := len(src)
	if srcOffset > srcLen || count > srcLen || srcOffset+count > srcLen {
		return false, errors.New("the src buffer index is out of rang")
	}
	dstLen := len(dst)
	if dstOffset > dstLen || count > dstLen || dstOffset+count > dstLen {
		return false, errors.New("the dst buffer index is out of rang")
	}
	index := 0
	for i := srcOffset; i < srcOffset+count; i++ {
		dst[dstOffset+index] = src[srcOffset+index]
		index++
	}
	return true, nil
}
