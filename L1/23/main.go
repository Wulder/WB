package main

import "fmt"

func main() {
	s := []string{"a", "b", "c", "d", "e", "f", "G", "h", "i", "j"}

	fmt.Println(s)
	deleteEl2(s, 6)
	fmt.Println(s)
}

func deleteEl(s []string, i int) { //медленный способ
	s = append(s[:i], s[i+1:]...)
}

func deleteEl2(s []string, i int) { //способ быстрее
	s[i], s[len(s)-1] = s[len(s)-1], s[i]
	s = s[:len(s)-1]
}
