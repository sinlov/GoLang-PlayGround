package go_regexp

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestMatchStringTargetChinese(t *testing.T) {
	// mock
	word := `今天是个好天气`
	regexpRight := `(\\u4eca)+(\\u5929)+`
	regexpError := `(\\u592a)+(\\u9633)+`
	// do
	actRight := MatchStringTargetChinese(word, regexpRight)
	actError := MatchStringTargetChinese(word, regexpError)
	// v
	assert.Equal(t, true, actRight)
	assert.Equal(t, false, actError)
}

func Test_tranWord2ASCII(t *testing.T) {
	word := `今天大小`
	t.Logf("%v tran -> %v", word, tranWord2ASCII(word))
	word2 := `太阳出来`
	t.Logf("%v tran -> %v", word2, tranWord2ASCII(word2))
	word3 := `今天是个好天气`
	t.Logf("%v tran -> %v", word3, tranWord2ASCII(word3))
}

func Test_regexp_target_chinese(t *testing.T) {
	word := `今天大小`
	regex := `[\\u4eca\\u5929]` + `.*`
	mustCompile := regexp.MustCompile(regex)
	match := mustCompile.MatchString(tranWord2ASCII(word))
	assert.Equal(t, true, match)
	t.Logf("word: %v , regexp %v", word, match)
}
