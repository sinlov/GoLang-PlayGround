package emoji

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

const (
	replacementChar = '\uFFFD'     // Unicode replacement character
	maxRune         = '\U0010FFFF' // Maximum valid Unicode code point.
)

const (
	// 0xd800-0xdc00 encodes the high 10 bits of a pair.
	// 0xdc00-0xe000 encodes the low 10 bits of a pair.
	// the value is those 20 bits plus 0x10000.
	surr1 = 0xd800
	surr2 = 0xdc00
	surr3 = 0xe000

	surrSelf = 0x10000
)

var (
	//emojiRegexp    = regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+]")
	emojiRegexp = regexp.MustCompile(`\[\\u[0-9a-zA-Z]+]`)
	//emojiGenRegexp = regexp.MustCompile("\\[\\\\u|]")
	emojiGenRegexp = regexp.MustCompile(`\[\\u|]`)
)

// unicode to emoji
func UnicodeToEmoji(unicode string) string {
	//emoji
	re := emojiRegexp
	//gen emoji
	reg := emojiGenRegexp
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

// cause Java Javascript code design, emoji use Modified UTF-8
// https://en.wikipedia.org/wiki/UTF-8#Modified_UTF-8
// https://stackoverflow.com/questions/32205446/getting-true-utf-8-characters-in-java-jni
// Java's Modified UTF-8 is not exactly the same as https://en.wikipedia.org/wiki/CESU-8
func ModifiedEmojiStringToUnicode(emoji string) (string, error) {
	unicode, err := ModifiedEmojiToUnicode([]byte(emoji))
	return string(unicode), err
}

// cause Java Javascript code design, emoji use Modified UTF-8
// https://en.wikipedia.org/wiki/UTF-8#Modified_UTF-8
// https://stackoverflow.com/questions/32205446/getting-true-utf-8-characters-in-java-jni
func ModifiedEmojiToUnicode(emoji []byte) ([]rune, error) {
	if len(emoji) <= 0 {
		return nil, fmt.Errorf("len(b) <= 0")
	}

	rr := []rune{}
	index := 0

	for index < len(emoji) {
		switch {
		case emoji[index] == 0b11101101:
			_ = emoji[index]
			v := emoji[index+1]
			w := emoji[index+2]
			_ = emoji[index+3]
			y := emoji[index+4]
			z := emoji[index+5]

			rr = append(rr, rune(0x10000+(int(v)&0x0f)<<16+(int(w)&0x3f)<<10+(int(y)&0x0f)<<6+int(z)&0x3f))

			index += 6

		case 0b11100000 <= emoji[index] && emoji[index] <= 0b11101111:
			x := emoji[index]
			y := emoji[index+1]
			z := emoji[index+2]

			rr = append(rr, rune(int(x)&0xf<<12+(int(y)&0x3f)<<6+int(z)&0x3f))

			index += 3

		case 0b11000000 <= emoji[index] && emoji[index] <= 0b11011111:
			x := emoji[index]
			y := emoji[index+1]

			rr = append(rr, rune((int(x)&0x1f)<<6+int(y)&0x3f))

			index += 2

		case emoji[index] <= 0b01111111:

			rr = append(rr, rune(emoji[index]))

			index++

		default:
			return rr, fmt.Errorf("unexpected byte: %b, index: %v", emoji[0], index)
		}
	}

	return rr, nil
}

func EncodeStringToUTF16Int(s string) []uint16 {
	n := 0
	for _, v := range s {
		n++
		if v >= 0x10000 {
			n++
		}
	}

	a := make([]uint16, n)
	n = 0
	for _, v := range s {
		switch {
		case 0 <= v && v < surr1, surr3 <= v && v < surrSelf:
			// normal rune
			a[n] = uint16(v)
			n++
		case surrSelf <= v && v <= maxRune:
			// needs surrogate sequence
			r1, r2 := utf16.EncodeRune(v)
			a[n] = uint16(r1)
			a[n+1] = uint16(r2)
			n += 2
		}
	}
	return a[:n]
}

func DecodeUTF16IntToString(s []uint16) string {
	n := 0
	for i := 0; i < len(s); i++ {
		switch r := s[i]; {
		case r < surr1, surr3 <= r:
			// normal rune
			n += utf8.RuneLen(rune(r))
		case surr1 <= r && r < surr2 && i+1 < len(s) &&
			surr2 <= s[i+1] && s[i+1] < surr3:
			// valid surrogate sequence
			n += utf8.RuneLen(utf16.DecodeRune(rune(r), rune(s[i+1])))
			i++
		default:
			// invalid surrogate sequence
			n += utf8.RuneLen(replacementChar)
		}
	}
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(s); i++ {
		switch r := s[i]; {
		case r < surr1, surr3 <= r:
			// normal rune
			b.WriteRune(rune(r))
		case surr1 <= r && r < surr2 && i+1 < len(s) &&
			surr2 <= s[i+1] && s[i+1] < surr3:
			// valid surrogate sequence
			b.WriteRune(utf16.DecodeRune(rune(r), rune(s[i+1])))
			i++
		default:
			// invalid surrogate sequence
			b.WriteRune(replacementChar)
		}
	}
	return b.String()
}

func EmojiToUTF16(emoji string) []byte {
	midUTF16 := EncodeStringToUTF16Int(emoji)
	res := make([]byte, 2)
	for _, code := range midUTF16 {
		res = append(res, byte(code))
	}
	return res
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
