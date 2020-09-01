package emoji

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
