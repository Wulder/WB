package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	powResult := make(chan int, len(arr)) //создание бурезированного канала с размером буфера равным длинне массива чисел
	for _, i := range arr {
		go func(n int) {
			powResult <- Pow(n) //запись квадрата числа в канал
		}(i)
	}

	var result int //сумма всех квадратов
	for i := 0; i < len(arr); i++ {
		result += <-powResult //чтение квадрата и прибавление его к сумме всех квадратов
	}

	fmt.Println("result =", result)

}

func Pow(i int) int {
	return i * i
}
