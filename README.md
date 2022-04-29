# goPY

[![GoCI](https://github.com/xuender/gopy/workflows/Go/badge.svg)](https://github.com/xuender/gopy/actions)
[![codecov](https://codecov.io/gh/xuender/gopy/branch/main/graph/badge.svg?token=8CTpNIHxYT)](https://codecov.io/gh/xuender/gopy)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/gopy)](https://goreportcard.com/report/github.com/xuender/gopy)
[![GoDoc](https://godoc.org/github.com/xuender/gopy?status.svg)](https://pkg.go.dev/github.com/xuender/gopy)
[![Gitter](https://badges.gitter.im/xuender-gopy/community.svg)](https://gitter.im/xuender-gopy/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![GitHub license](https://img.shields.io/github/license/xuender/gopy)](https://github.com/xuender/gopy/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/xuender/gopy)](https://github.com/xuender/gopy/issues)
[![GitHub stars](https://img.shields.io/github/stars/xuender/gopy)](https://github.com/xuender/gopy/stargazers)

中文转拼音.

## use

```shell
go get github.com/xuender/gopy
```

```go
package main

import (
  "fmt"

  "github.com/xuender/gopy"
  )

func main(){
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
```

## cmd

```shell
gopy --help
```

### install

```shell
go install github.com/xuender/gopy/cmd/gopy@latest
```

### examples

```shell
gopy 阿弥陀佛
gopy -i -s= 阿弥陀佛
gopy -t -h -hs=\; 阿弥陀佛
gopy -n 阿弥陀佛
```
