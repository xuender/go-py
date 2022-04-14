package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gitee.com/xuender/oils/base"
	"github.com/xuender/gopy"
	"github.com/xuender/gopy/cmd"
)

func main() {
	var (
		tone      bool
		noTone    bool
		init      bool
		normal    bool
		heteronym bool
		sep       string
		hsep      string
	)

	flag.Usage = usage
	flag.BoolVar(&normal, "d", false, "default")
	flag.BoolVar(&heteronym, "h", false, "show heterony")
	flag.BoolVar(&init, "i", false, "initials")
	flag.BoolVar(&noTone, "n", false, "no tone")
	flag.BoolVar(&tone, "t", false, "tone")
	flag.StringVar(&sep, "s", " ", "separators")
	flag.StringVar(&hsep, "hs", ",", "heterony separators")
	flag.Parse()

	if len(flag.Args()) == 0 {
		usage()
		os.Exit(0)
	}

	defer Panic()

	input := strings.Join(flag.Args(), " ")
	output := base.NewSlice[string]()

	if normal {
		output.Add(toPinyin(input, sep, hsep, heteronym, gopy.Normal))
	}

	if init {
		output.Add(toPinyin(input, sep, hsep, heteronym, gopy.Init))
	}

	if noTone {
		output.Add(toPinyin(input, sep, hsep, heteronym, gopy.NoTone))
	}

	if tone {
		output.Add(toPinyin(input, sep, hsep, heteronym, gopy.Tone))
	}

	if len(output) == 0 {
		output.Add(gopy.Py(input, sep))
	}

	for _, out := range output {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", out)
	}
}

func toPinyin(str, sep, hsep string, heteronym bool, option gopy.Option) string {
	if heteronym {
		slice := base.NewSlice[string]()
		for _, pinyin := range gopy.Pinyins(str, option) {
			slice.Add(strings.Join(pinyin, hsep))
		}

		return slice.Join(sep)
	}

	return strings.Join(gopy.Pinyin(str, option), sep)
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "gopy[%s]\n\n", cmd.ModVersion)
	fmt.Fprintf(flag.CommandLine.Output(), "chinese to pinyin.\n\n")
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), `
Examples:
	gopy 阿弥陀佛
	gopy -i -s= 阿弥陀佛
	gopy -t -h -hs=\; 阿弥陀佛
	gopy -n 阿弥陀佛
`)
	fmt.Fprintf(flag.CommandLine.Output(), "\nmod: %s\ngit: %s\n", cmd.ModPath, cmd.ModSum)
}

func Panic() {
	if err := recover(); err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "error: %v\n", err)

		os.Exit(1)
	}
}
