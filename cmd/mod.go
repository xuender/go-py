package cmd

import (
	"runtime/debug"
)

// nolint: gochecknoglobals
var (
	// ModVersion 版本号.
	ModVersion = "(devel)"
	// ModSum check sum.
	ModSum = "00000000"
	// ModPath 模块路径.
	ModPath = ""
)

// nolint
func init() {
	if ModVersion != "(devel)" {
		return
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	mod := &info.Main
	if mod.Replace != nil {
		mod = mod.Replace
	}

	ModVersion = mod.Version
	ModSum = mod.Sum
	ModPath = mod.Path
}
