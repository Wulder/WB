package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var arr []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ { //заполнение массива рандомными значениями
		arr = append(arr, rand.Intn(100))
	}

	fmt.Println(arr)
	quickSort(arr)

	fmt.Println(arr)
}

func quickSort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return
}
