package py_test

import (
	"testing"

	"github.com/xuender/go-py"
	"github.com/xuender/oils/assert"
)

func TestRuneOption(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"1"}, py.RuneOption('1', py.Normal, -1, nil))
	assert.Equals(t, []string{"jia4o", "jia1o"}, py.RuneOption('教', py.Normal, -1, nil))
	assert.Equal(t, "a1", py.RuneOption('阿', py.Normal, -1, nil)[0])
	assert.Equals(t, []string{"zha3ng", "cha2ng"}, py.RuneOption('长', py.Normal, -1, nil))
	assert.Equals(t, []string{"b", "f"}, py.RuneOption('不', py.Init, -1, nil))
	assert.Equals(t, []string{"y"}, py.RuneOption('与', py.Init, -1, nil))
	assert.Equals(t, []string{"d", "c", "t", "z", "s"}, py.RuneOption('撣', py.Init, -1, nil))
	assert.Equals(t, []string{"d", "t", "z"}, py.RuneOption('敦', py.Init, -1, nil))
	assert.Equals(t, []string{"b", "f", "l", "p"}, py.RuneOption('賁', py.Init, -1, nil))
	assert.Equals(t, []string{"m", "j", "l"}, py.RuneOption('繆', py.Init, -1, nil))
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

func TestToPolyphonic(t *testing.T) {
	t.Parallel()

	poly := py.ToPolyphonic(py.Pinyins("长", py.Init)[0], '长', 0, []rune("长相"))
	assert.Equals(t, []string{"z", "c"}, poly)

	poly = py.ToPolyphonic(py.Pinyins("长", py.Init)[0], '长', 0, []rune("特长"))
	assert.Equals(t, []string{"c", "z"}, poly)
}
