package security

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func URLPathEncode(input string) string {
	return url.PathEscape(input)
}

func URLPathDecode(input string) (string, error) {
	return url.PathUnescape(input)
}

func URLQueryEncode(input string) string {
	return url.QueryEscape(input)
}

func URLQueryDecode(input string) (string, error) {
	return url.QueryUnescape(input)
}

func UnicodeDecode(input string) (out string, err error) {
	var res = []string{""}
	sUnicode := strings.Split(input, "\\u")

	if len(sUnicode) < 2 {
		return "", fmt.Errorf("input unicode error, string must has %v", "\\u")
	}

	var context = ""
	for _, v := range sUnicode {
		var additional = ""
		if len(v) < 1 {
			continue
		}
		if len(v) > 4 {
			rs := []rune(v)
			v = string(rs[:4])
			additional = string(rs[4:])
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			context += v
		}
		context += fmt.Sprintf("%c", temp)
		context += additional
	}
	res = append(res, context)
	return strings.Join(res, ""), nil
}
