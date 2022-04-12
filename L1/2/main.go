package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup //объект синхронизации

	for _, i := range arr {
		wg.Add(1)
		go func(n int) { //запуск горутины для просчета квадрата элемента из arr
			fmt.Println("square of", n, "=", Pow(n))
			wg.Done()
		}(i)
	}

	wg.Wait() //ожидание завершения всех корутин вызванных в цикле выше
	fmt.Println("execute time:", time.Since(start))
}

func Pow(i int) int {
	return i * i
}
