package go_safe

import (
	"github.com/sinlov/fastEncryptDecode"
	"github.com/xxtea/xxtea-go/xxtea"
)

func XByteEncode(data []byte) []byte {
	KEY := []byte("1234567890")
	return xxtea.Encrypt(data, KEY)
}

func XByteDecode(data []byte) []byte {
	KEY := []byte("1234567890")
	return xxtea.Decrypt(data, KEY)
}

func XHexStringEncode(data string) string {
	dataEn := []byte(data)
	byteDataEn := XByteEncode(dataEn)
	return fastEncryptDecode.ByteArr2HexStr(byteDataEn)
}

func XHexStringDecode(data string) string {
	dataDe, err := fastEncryptDecode.HexStr2ByteArr(data)
	if err != nil {
		return err.Error()

	}
	byteDataDe := XByteDecode(dataDe)
	outStr := string(byteDataDe)
	return outStr
}
