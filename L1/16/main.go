package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var arr [1000]int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ { //заполнение массива рандомными значениями
		arr[i] = rand.Intn(100000)
	}

	fmt.Println(quickSort(arr[:]))
}

func quickSort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivotIndex := rand.Int() % len(a)

	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
