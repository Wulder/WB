package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//2 множества
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})

	intersection := make(map[int]struct{}) //пересечение

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		m1[rand.Intn(20)] = struct{}{}
		m2[rand.Intn(20)] = struct{}{}
	}

	for k, _ := range m1 {

		if _, ok := m2[k]; ok {
			intersection[k] = struct{}{}
		}
	}

	printM(m1)
	printM(m2)
	printM(intersection)
}

func printM(m map[int]struct{}) { //удобный принт множеств (ключей мап)
	var s string
	for k, _ := range m {
		s += fmt.Sprint(k, ";")
	}
	fmt.Println(s)
}
