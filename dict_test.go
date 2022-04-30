package py_test

import (
	_ "embed"
	"testing"

	"github.com/xuender/go-py"
	"github.com/xuender/oils/assert"
)

func TestDict_Rune(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"1"}, py.Runes('1', py.Normal))
	assert.Equals(t, []string{"jia4o", "jia1o"}, py.Runes('教', py.Normal))
	assert.Equal(t, "a1", py.Runes('阿', py.Normal)[0])
	assert.Equals(t, []string{"zha3ng", "cha2ng"}, py.Runes('长', py.Normal))
}

func TestStyle(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "a", py.Style("a1i", py.Init))
	assert.Equal(t, "a1", py.Style("a1", py.Normal))
	assert.Equal(t, "ā", py.Style("a1", py.Tone))
	assert.Equal(t, "ai", py.Style("a1i", py.NoTone))
}

func TestIsHan(t *testing.T) {
	t.Parallel()

	assert.False(t, py.IsHan(0xf890))
	assert.False(t, py.IsHan(0xfa2f))
	assert.True(t, py.IsHan('阿'))
	assert.False(t, py.IsHan('a'))
}
