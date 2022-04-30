package py

import (
	"strings"

	"golang.org/x/exp/slices"
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
		ret[i] = RuneOption(han, style, i, runes)[0]
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
		ret[i] = RuneOption(han, style, i, runes)
	}

	return ret
}

func Initials(str string) string {
	return Py(str, "", Init)
}

func ToneIndex(tone string) int {
	// return slices.Index(tones, FromStyle(tone))
	return slices.Index(tones, tone)
}
