package py

import (
	"strings"
)

type Option int

const (
	Normal Option = iota
	Tone
	NoTone
	Init
)

func Py(str, sep string, options ...Option) string {
	return strings.Join(Pinyin(str, options...), sep)
}

func Pinyin(str string, options ...Option) []string {
	runes := []rune(str)
	ret := make([]string, len(runes))
	style := Normal

	if len(options) > 0 {
		style = options[0]
	}

	for i, han := range runes {
		ret[i] = Runes(han, style)[0]
	}

	return ret
}

func Pinyins(str string, options ...Option) [][]string {
	runes := []rune(str)
	ret := make([][]string, len(runes))
	style := Normal

	if len(options) > 0 {
		style = options[0]
	}

	for i, han := range runes {
		ret[i] = Runes(han, style)
	}

	return ret
}

func Initials(str string) string {
	return Py(str, "", Init)
}

func Rune(han rune) []string {
	return Runes(han, Tone)
}
