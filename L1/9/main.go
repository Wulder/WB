package main

import "fmt"

func main() {

	var arr [100]int
	for i := range arr {
		arr[i] = i
	}

	chan1 := make(chan int) //канал для чисел из массива
	chan2 := make(chan int) //канал для записи в него квадратов чисел из первого канала

	go func() {
		for _, i := range arr {
			chan1 <- i //кладем элемент массива в первый канал
		}
		close(chan1) //закрываем первый канал после отправки туда каждого элемента массива
	}()

	go func() {
		for x := range chan1 { //ожидаем данные из первого канала
			chan2 <- x * x //кладем квадрат числв из первого канала во второй
		}
		close(chan2) //закрываем второй канал, после отправки туда всех данных с первого канала
	}()

	for i := range chan2 { //ожидаем данные со второго канала
		fmt.Println(i) //кидаем пришедшие с канала данные в вывод
	}
}
