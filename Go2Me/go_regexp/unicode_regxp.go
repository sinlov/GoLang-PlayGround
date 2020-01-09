package go_regexp

import (
	"regexp"
	"strconv"
)

func tranWord2ASCII(word string) string {
	return strconv.QuoteToASCII(word)
}

func MatchStringTargetChinese(word, regexpKey string) bool {
	quoteToASCII := strconv.QuoteToASCII(word)
	mustCompile := regexp.MustCompile(regexpKey)
	return mustCompile.MatchString(quoteToASCII)
}
