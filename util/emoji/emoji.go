package emoji

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//  unicode to emoji
func UnicodeToEmoji(unicode string) string {
	//emoji
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//gen emoji
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(unicode, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			unicode = strings.Replace(unicode, src[i], string(rune(p)), -1)
		}
	}
	return unicode
}

// emoji to unicode
func EmojiToUnicode(emoji string) string {
	ret := ""
	rs := []rune(emoji)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u

		} else {
			ret += string(rs[i])
		}
	}
	return ret
}

func EmojiToCodePoint(emoji string) (uint64, error) {
	ret := ""
	rs := []rune(emoji)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := strconv.FormatInt(int64(rs[i]), 16)
			ret += u

		} else {
			ret += string(rs[i])
		}
	}
	return strconv.ParseUint(ret, 16, 32)
}

func CodePoint2Emoji(codePoint uint64) string {
	unicode := fmt.Sprintf("[\\u%x]", codePoint)
	return UnicodeToEmoji(unicode)
}
