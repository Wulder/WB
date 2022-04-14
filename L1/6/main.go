package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var isEnd bool //глобальная переменная

func main() {

	var wg sync.WaitGroup

	chanel1 := make(chan int)
	go goroutine1(chanel1) //запуск горутины
	wg.Add(1)
	chanel1 <- 1 //окончание горутины
	close(chanel1)
	wg.Done()

	ctx, cancel := context.WithCancel(context.Background()) //создание контекста для горутины
	go goroutine2(ctx)                                      //запуск горутины
	wg.Add(1)
	time.Sleep(time.Second * 1)
	cancel() //окончание горутины
	wg.Done()

	go goroutine3() //запуск горутины
	wg.Add(1)
	time.Sleep(time.Second * 1)
	isEnd = true //окончание горутины
	wg.Done()

	var goroutine4wg sync.WaitGroup // WaitGroup для горутины
	goroutine4wg.Add(1)
	go goroutine4(&goroutine4wg) //начало горутины
	wg.Add(1)
	goroutine4wg.Done() //окончание горутины
	time.Sleep(time.Second * 1)
	wg.Done()

	wg.Wait() //ожидание завершения всех горутин
}

func goroutine1(c <-chan int) { //ожидание записи в канал и его последующего прочтения

	fmt.Println("goroutine1 - Some work...")
	<-c //ожидание записи в канал
	fmt.Println("goroutine1 is over")
	return
}

func goroutine2(ctx context.Context) { //завершение выполнения горутины по при окончании выполнения контекста

	select {
	case <-ctx.Done():
		break
	default:
		fmt.Println("goroutine2 - Some work...")
	}

	fmt.Println("goroutine2 is over")
	return
}

func goroutine3() { //завершение работы горутины тупой проверкой значения переменной
	for {
		if isEnd {
			fmt.Println("goroutine3 is over")
			return
		}
		time.Sleep(time.Second * 1)
		fmt.Println("goroutine3 - Some work...")
	}

}

func goroutine4(w *sync.WaitGroup) { //завершение горутины используя объект синхронизации в видео WiteGroup
	fmt.Println("goroutine4 - Some work...")
	w.Wait()
	fmt.Println("goroutine4 is over")
	return
}
