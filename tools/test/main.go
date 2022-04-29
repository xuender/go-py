package main

import "fmt"

func main() {
	dic := map[rune][]int16{
		1: {3},
		2: {3, 2, 3},
	}

	fmt.Println(dic)
}
