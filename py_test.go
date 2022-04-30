package py_test

import (
	"fmt"
	"testing"

	"github.com/xuender/go-py"
	"github.com/xuender/oils/assert"
)

func ExamplePinyin() {
	fmt.Println(py.Pinyin("阿弥陀佛"))
	fmt.Println(py.Pinyin("阿弥陀佛", py.Tone))
	fmt.Println(py.Pinyin("阿弥陀佛", py.NoTone))
	fmt.Println(py.Pinyin("阿弥陀佛", py.Init))
	fmt.Println(py.Initials("阿弥陀佛"))

	// output:
	// [a1 mi2 tuo2 fu2]
	// [ā mí tuó fú]
	// [a mi tuo fu]
	// [a m t f]
	// amtf
}

func TestRune(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"1"}, py.Rune('1'))
}

func TestInitials(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "zs", py.Initials("张三"))
	assert.Equal(t, "123", py.Initials("123"))
}

func TestPinyin(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"1", "2", "3"}, py.Pinyin("123"))
	assert.Equals(t, []string{"zha3ng"}, py.Pinyin("长"))
	assert.Equals(t, []string{"zhang"}, py.Pinyin("长", py.NoTone))
}

func TestPinyins(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "1", py.Pinyins("123")[0][0])
	assert.Equals(t, []string{"zha3ng", "cha2ng"}, py.Pinyins("长")[0])
	assert.Equals(t, []string{"zhang", "chang"}, py.Pinyins("长", py.NoTone)[0])
}

func TestPy(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "chang zhang", py.Py("厂长", " ", py.NoTone))
}
