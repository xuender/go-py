# gopy

Chinese to pinyin by golang.

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
## cli

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
