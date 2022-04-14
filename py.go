package gopy

import (
	"strings"
)

type Option int

const (
	Normal Option = iota
	Tone
	NoTone
	Initial
)

func Pinyin(str string, options ...Option) []string {
	runes := []rune(str)
	ret := make([]string, len(runes))
	style := Normal

	if len(options) > 0 {
		style = options[0]
	}

	for i, han := range runes {
		ret[i] = data.Rune(han, style)[0]
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
		ret[i] = data.Rune(han, style)
	}

	return ret
}

func Initials(str string) string {
	return strings.Join(Pinyin(str, Initial), "")
}

func Rune(han rune) []string {
	return data.Rune(han, Tone)
}
