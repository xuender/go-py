package gopy_test

import (
	_ "embed"
	"testing"

	"gitee.com/xuender/oils/assert"
	"github.com/xuender/gopy"
)

func TestDict_Rune(t *testing.T) {
	t.Parallel()

	data := gopy.Dict{}
	data[38463] = []int16{1}
	data[38271] = []int16{116, 1405}

	assert.Equals(t, []string{"1"}, data.Rune('1', gopy.Normal))
	assert.Equals(t, []string{"教"}, data.Rune('教', gopy.Normal))
	assert.Equals(t, []string{"a1"}, data.Rune('阿', gopy.Normal))
	assert.Equals(t, []string{"cha2ng", "zha3ng"}, data.Rune('长', gopy.Normal))
}

func TestStyle(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "a", gopy.Style("a1i", gopy.Initial))
	assert.Equal(t, "a1", gopy.Style("a1", gopy.Normal))
	assert.Equal(t, "ā", gopy.Style("a1", gopy.Tone))
	assert.Equal(t, "ai", gopy.Style("a1i", gopy.NoTone))
}

func TestIsHan(t *testing.T) {
	t.Parallel()

	assert.False(t, gopy.IsHan(0xf890))
	assert.False(t, gopy.IsHan(0xfa2f))
	assert.True(t, gopy.IsHan('阿'))
	assert.False(t, gopy.IsHan('a'))
}
