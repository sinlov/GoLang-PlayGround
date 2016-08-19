package go_parse

import (
	"fmt"
	"strconv"
	"strings"
)

func Utf82Unicode() {
	bStr := "转换前的中文"
	res := utf82Unicode(bStr)
	fmt.Println(bStr, res)
}

func Unicode2utf8() {
	bStr := "转换前的中文"
	enUnicode := utf82Unicode(bStr)
	enUtf8 :=unicode2Utf8(enUnicode)
	fmt.Println(bStr, enUnicode, enUtf8)
}

func utf82Unicode(text string) string  {
	cover := strconv.QuoteToASCII(text);
	res := cover[1 : len(cover) - 1]
	return res
}

func unicode2Utf8(text string) string {
	unicodeTemp := strings.Split(text, "\\u")
	var context string
	for _, v := range unicodeTemp {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	return context
}
