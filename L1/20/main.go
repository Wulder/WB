package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	words := strings.Fields(in.Text())

	l := len(words) - 1
	for i := 0; i < l; i++ {
		words[i], words[l] = words[l], words[i]
		l--
	}

	fmt.Println(strings.Join(words, " "))

}
