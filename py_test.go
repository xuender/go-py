package gopy_test

import (
	"fmt"
	"testing"

	"github.com/xuender/gopy"
	"github.com/xuender/oils/assert"
)

func ExamplePinyin() {
	fmt.Println(gopy.Pinyin("阿弥陀佛"))
	fmt.Println(gopy.Pinyin("阿弥陀佛", gopy.Tone))
	fmt.Println(gopy.Pinyin("阿弥陀佛", gopy.NoTone))
	fmt.Println(gopy.Pinyin("阿弥陀佛", gopy.Init))
	fmt.Println(gopy.Initials("阿弥陀佛"))

	// output:
	// [a1 mi2 tuo2 fu2]
	// [ā mí tuó fú]
	// [a mi tuo fu]
	// [a m t f]
	// amtf
}

func TestRune(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"1"}, gopy.Rune('1'))
}

func TestInitials(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "zs", gopy.Initials("张三"))
	assert.Equal(t, "123", gopy.Initials("123"))
}

func TestPinyin(t *testing.T) {
	t.Parallel()

	assert.Equals(t, []string{"1", "2", "3"}, gopy.Pinyin("123"))
	assert.Equals(t, []string{"zha3ng"}, gopy.Pinyin("长"))
	assert.Equals(t, []string{"zhang"}, gopy.Pinyin("长", gopy.NoTone))
}

func TestPinyins(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "1", gopy.Pinyins("123")[0][0])
	assert.Equals(t, []string{"zha3ng", "cha2ng"}, gopy.Pinyins("长")[0])
	assert.Equals(t, []string{"zhang", "chang"}, gopy.Pinyins("长", gopy.NoTone)[0])
}

func TestPy(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "chang zhang", gopy.Py("厂长", " ", gopy.NoTone))
}
