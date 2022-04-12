package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

UpLoop:
	for {
		in := bufio.NewScanner(os.Stdin)
		in.Scan()

		for _, s := range in.Text() {
			if strings.Count(in.Text(), string(s)) > 1 {
				fmt.Println(in.Text(), "-", false)
				continue UpLoop

			}
		}
		fmt.Println(in.Text(), "-", true)
	}

}
