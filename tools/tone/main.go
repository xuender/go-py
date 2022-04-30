package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/xuender/go-py"
	"github.com/xuender/oils/base"
)

func main() {
	poly := base.Panic1(os.Open(filepath.Join("data", "polyphonic.txt")))
	defer poly.Close()

	bytes := base.Panic1(io.ReadAll(poly))
	newStr := py.FromStyle(string(bytes))

	nf := base.Panic1(os.Create(filepath.Join("data", "p.txt")))
	defer nf.Close()

	_, _ = io.WriteString(nf, newStr)
}
