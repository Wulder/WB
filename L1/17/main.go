package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 34, 576, 788} //отсортированный набор данных

	println("Recursive:")
	fmt.Println(binarySearchRecursive(arr, 3))  //true
	fmt.Println(binarySearchRecursive(arr, 33)) //false

	println("\nLoops:")
	fmt.Println(binarySearchLoop(arr, 3))
	fmt.Println(binarySearchLoop(arr, 33))
}

func binarySearchRecursive(arr []int, x int) bool { //рекурсивный метод
	pivot := len(arr) / 2

	if x < arr[pivot] && len(arr) > 1 {
		return binarySearchRecursive(arr[:len(arr)/2], x)
	} else if x > arr[pivot] && len(arr) > 1 {
		return binarySearchRecursive(arr[len(arr)/2:], x)
	} else if x == arr[pivot] {
		return true
	} else {
		return false
	}

}

func binarySearchLoop(arr []int, x int) bool { //метод с помощью цикла (экономится память в отличие от рекурсивного, так как не создаются дополнительные слайсы)
	pivot := len(arr) / 2
	i := 0
	for {
		i++
		if -1 < pivot && pivot <= len(arr)-1 {
			if x < arr[pivot] {
				pivot -= len(arr) / (i * 2)
				continue
			} else if x > arr[pivot] {
				pivot += len(arr) / (i * 2)
				continue
			} else if x == arr[pivot] {
				return true
			}
		}
		return false
	}
}
