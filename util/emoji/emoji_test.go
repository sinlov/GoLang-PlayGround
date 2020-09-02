package emoji

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

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

func Test_UnicodeToEmoji(t *testing.T) {
	target := "[\\u1f601]more test of 中文"
	toEmoji := UnicodeToEmoji(target)
	t.Logf("target %v UnicodeToEmoji %v", target, toEmoji)
	assert.Equal(t, `😁more test of 中文`, toEmoji)
}

func Test_EmojiToUnicode(t *testing.T) {
	target := `😁more test of 中文`
	toUnicode := EmojiToUnicode(target)
	t.Logf("target %v EmojiToUnicode %v", target, toUnicode)
	assert.Equal(t, `[\u1f601]more test of 中文`, toUnicode)
}

func Test_EmojiModifiedStringToUnicode(t *testing.T) {
	target := `eda0bdedb0924d794170706c69636174696f6ee4b8ade69687`
	targetByte, err := HexStr2ByteArr(target)
	t.Logf("target %v toHex %v", targetByte, target)
	assert.Nil(t, err)
	unicode, err := ModifiedEmojiStringToUnicode(string(targetByte))
	t.Logf("target %v ModifiedEmojiStringToUnicode %v", string(targetByte), unicode)
	assert.Nil(t, err)
}

func Test_EncodeStringToUTF16(t *testing.T) {
	monkey := `🐒`
	t.Logf("%v", fmt.Sprintf("monkey %v to HEX %x", monkey, []byte(monkey)))
	stringToUTF16 := EncodeStringToUTF16Int(monkey)

	t.Logf("%v", fmt.Sprintf("stringToUTF16 %x", stringToUTF16))

	utf16IntToString := DecodeUTF16IntToString(stringToUTF16)
	t.Logf("%v", fmt.Sprintf("utf16IntToString %v", utf16IntToString))
}

func Test_EmojiToUTF16(t *testing.T) {
	target := `🐒`
	emojiToUTF16 := EmojiToUTF16(target)
	t.Logf("target %v EmojiToUTF16 %v", target, emojiToUTF16)
	t.Logf("%v", fmt.Sprintf("emojiToUTF16 %x", emojiToUTF16))
}

func Test_EmojiToCodePoint(t *testing.T) {
	// use js at https://github.com/channg/umoji
	target := `😁`
	toCodePoint, err := EmojiToCodePoint(target)
	assert.Nil(t, err)
	t.Logf("target %v EmojiToCodePoint %v", target, toCodePoint)
	assert.Equal(t, uint64(128513), toCodePoint)
}

func Test_CodePoint2Emoji(t *testing.T) {
	// use js at https://github.com/channg/umoji
	target := 128513
	codePoint2Emoji := CodePoint2Emoji(uint64(target))
	t.Logf("target %v CodePoint2Emoji %v", target, codePoint2Emoji)
	assert.Equal(t, `😁`, codePoint2Emoji)
}
