package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	arr := []rune(in.Text())

	l := len(arr) - 1
	for i := 0; i <= l; i++ {
		arr[l], arr[i] = arr[i], arr[l]
		l--
	}
	fmt.Println(string(arr))

}
