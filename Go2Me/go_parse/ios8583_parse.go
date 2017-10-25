package go_parse

import (
	"github.com/ideazxy/iso8583"
)

type Data struct {
	No   *iso8583.Numeric      `field:"3" length:"6" encode:"bcd"`    // bcd value encoding
	Oper *iso8583.Numeric      `field:"26" length:"2" encode:"ascii"` // ascii value encoding
	Ret  *iso8583.Alphanumeric `field:"39" length:"2"`
	Sn   *iso8583.Llvar        `field:"45" length:"23" encode:"bcd,ascii"` // bcd length encoding, ascii value encoding
	Info *iso8583.Lllvar       `field:"46" length:"42" encode:"bcd,ascii"`
	Mac  *iso8583.Binary       `field:"64" length:"8"`
}

func DataMessage8583BCD(mitStr string, data interface{}) ([]byte, error) {
	msg := iso8583.NewMessage(mitStr, data)
	msg.MtiEncode = iso8583.BCD
	b, err := msg.Bytes()
	return b, err
}

func DataMessage8583ASCII(mitStr string, data interface{}) ([]byte, error) {
	msg := iso8583.NewMessage(mitStr, data)
	msg.MtiEncode = iso8583.ASCII
	b, err := msg.Bytes()
	return b, err
}
