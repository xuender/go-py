package gopy

import (
	"strings"
	"unicode"

	"github.com/xuender/oils/base"
)

type Dict map[rune]string

// nolint
var phonetic = map[string]string{
	"a1": "ā",
	"a2": "á",
	"a3": "ǎ",
	"a4": "à",
	"e1": "ē",
	"e2": "é",
	"e3": "ě",
	"e4": "è",
	"o1": "ō",
	"o2": "ó",
	"o3": "ǒ",
	"o4": "ò",
	"i1": "ī",
	"i2": "í",
	"i3": "ǐ",
	"i4": "ì",
	"u1": "ū",
	"u2": "ú",
	"u3": "ǔ",
	"u4": "ù",
	"v":  "ü",
	"v2": "ǘ",
	"v3": "ǚ",
	"v4": "ǜ",
	"n2": "ń",
	"n3": "ň",
	"n4": "ǹ",
	"m2": "ḿ",
}

func Style(pinyin string, option Option) string {
	switch option {
	case Tone:
		for tone, pho := range phonetic {
			pinyin = strings.ReplaceAll(pinyin, tone, pho)
		}

		return pinyin
	case NoTone:
		slice := base.NewSlice([]rune(pinyin)...)
		slice.Del([]rune("1234")...)

		return string(slice)
	case Init:
		return pinyin[:1]
	case Normal:
		return pinyin
	default:
		return pinyin
	}
}

func IsHan(han rune) bool {
	if !unicode.Is(unicode.Scripts["Han"], han) {
		return false
	}

	if han < 0x4e00 || han > 0xfa2d || (han < 0xf900 && han > 0x9FA5) {
		return false
	}

	return true
}

func (p Dict) Rune(han rune, option Option) []string {
	pinyin, has := p[han]
	if !has {
		return []string{string(han)}
	}

	pys := strings.Split(pinyin, ",")
	ret := make([]string, len(pys))
	for index, pyIndex := range pys {
		num, _ := base.ParseInteger[int](pyIndex)
		ret[index] = Style(tones[num], option)
	}

	return ret
}
