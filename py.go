package gopy

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
	dic := Dict(dict)
	runes := []rune(str)
	ret := make([]string, len(runes))
	style := Normal

	if len(options) > 0 {
		style = options[0]
	}

	for i, han := range runes {
		ret[i] = dic.Rune(han, style)[0]
	}

	return ret
}

func Pinyins(str string, options ...Option) [][]string {
	dic := Dict(dict)
	runes := []rune(str)
	ret := make([][]string, len(runes))
	style := Normal

	if len(options) > 0 {
		style = options[0]
	}

	for i, han := range runes {
		ret[i] = dic.Rune(han, style)
	}

	return ret
}

func Initials(str string) string {
	return Py(str, "", Init)
}

func Rune(han rune) []string {
	dic := Dict(dict)

	return dic.Rune(han, Tone)
}
