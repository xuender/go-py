# go-py

[![GoCI](https://github.com/xuender/go-py/workflows/Go/badge.svg)](https://github.com/xuender/gopy/actions)
[![codecov](https://codecov.io/gh/xuender/go-py/branch/main/graph/badge.svg?token=1QQNBH82CM)](https://codecov.io/gh/xuender/gopy)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/go-py)](https://goreportcard.com/report/github.com/xuender/gopy)
[![GoDoc](https://godoc.org/github.com/xuender/go-py?status.svg)](https://pkg.go.dev/github.com/xuender/gopy)
[![Gitter](https://badges.gitter.im/xuender-go-py/community.svg)](https://gitter.im/xuender-gopy/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![GitHub license](https://img.shields.io/github/license/xuender/go-py)](https://github.com/xuender/gopy/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/xuender/go-py)](https://github.com/xuender/gopy/issues)
[![GitHub stars](https://img.shields.io/github/stars/xuender/go-py)](https://github.com/xuender/gopy/stargazers)

中文转拼音.

## use

```shell
go get github.com/xuender/go-py
```

```go
package main

import (
  "fmt"

  "github.com/xuender/go-py"
  )

func main(){
  fmt.Println(py.Pinyin("阿弥陀佛"))
  fmt.Println(py.Pinyin("阿弥陀佛", gopy.Tone))
  fmt.Println(py.Pinyin("阿弥陀佛", gopy.NoTone))
  fmt.Println(py.Pinyin("阿弥陀佛", gopy.Init))
  fmt.Println(py.Initials("阿弥陀佛"))

  // output:
  // [a1 mi2 tuo2 fu2]
  // [ā mí tuó fú]
  // [a mi tuo fu]
  // [a m t f]
  // amtf
}
```

## cmd

```shell
py --help
```

### install

```shell
go install github.com/xuender/go-py/cmd/py@latest
```

### examples

```shell
py 阿弥陀佛
py -i -s= 阿弥陀佛
py -t -h -hs=\; 阿弥陀佛
py -n 阿弥陀佛
```
