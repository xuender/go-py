package py

import (
	"strings"
	"unicode"

	"github.com/xuender/oils/base"
)

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

func FromStyle(pinyin string) string {
	for tone, pho := range phonetic {
		pinyin = strings.ReplaceAll(pinyin, pho, tone)
	}

	return pinyin
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

func ToStrings(pinyin []uint16, option Option) []string {
	ret := base.NewSlice[string]()

	for _, num := range pinyin {
		ret.Add(Style(tones[num], option))
	}

	ret.Unique()

	return ret
}

func ToPolyphonic(tones []string, han rune, index int, runes []rune) []string {
	if index < 0 {
		return tones
	}

	poly, has := polyphonic[han]
	if !has {
		return tones
	}

	start := index - base.Two
	if start < 0 {
		start = 0
	}

	end := index + base.Two
	if end > len(runes) {
		end = len(runes)
	}

	sub := string(runes[start:end])

	for _, pol := range poly {
		for _, word := range pol.Words {
			if strings.Contains(sub, word) {
				top := tones[pol.Tone]
				tones = append(tones[0:pol.Tone], tones[pol.Tone+1:]...)

				return append([]string{top}, tones...)
			}
		}
	}

	return tones
}

func RuneOption(han rune, option Option, index int, runes []rune) []string {
	if array, has := dict1[han]; has {
		return ToStrings(array[:], option)
	}

	return ToPolyphonic(cool(han, option), han, index, runes)
}

func cool(han rune, option Option) []string {
	if array, has := dict2[han]; has {
		return ToStrings(array[:], option)
	}

	if array, has := dict3[han]; has {
		return ToStrings(array[:], option)
	}

	if array, has := dict4[han]; has {
		return ToStrings(array[:], option)
	}

	if array, has := dict5[han]; has {
		return ToStrings(array[:], option)
	}

	if array, has := dict6[han]; has {
		return ToStrings(array[:], option)
	}

	if array, has := dict7[han]; has {
		return ToStrings(array[:], option)
	}

	if array, has := dict8[han]; has {
		return ToStrings(array[:], option)
	}

	if array, has := dict9[han]; has {
		return ToStrings(array[:], option)
	}

	if array, has := dict10[han]; has {
		return ToStrings(array[:], option)
	}

	return []string{string(han)}
}
