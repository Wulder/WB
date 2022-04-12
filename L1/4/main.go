package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	workersCount := ReadWorkesCount()

	fmt.Println(workersCount)
}

func ReadWorkesCount() int {
	in := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите количество воркеров: ")
	in.Scan()
	i, err := strconv.Atoi(in.Text())
	if err != nil {
		fmt.Println(err.Error())
		return ReadWorkesCount()
	} else if i == 0 {
		return ReadWorkesCount()
	}

	return i
}
