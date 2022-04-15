package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "abcd абвг"
	for k, v := range []rune(str) {
		fmt.Println(k, "-", v)
	}
	fmt.Println(strings.Repeat("-", 20))
	for k, v := range []byte(str) {
		fmt.Println(k, "-", v)
	}
}
