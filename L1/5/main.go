package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	seconds := 5 //N секунд после которых программа прекращает работу

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(seconds)) //создание контекста с таймаутом

	channel := make(chan string) //создание канала

	go sender(ctx, channel) //запуск горутины с сендером данных в канал
	reciever(ctx, channel)  //запуск слушателя в текущей горутине(main)
	close(channel)          //закрытие канала
	cancel()                //(!)IDE пишет предупреждение что нужно использовать функцию завершения контекста, не понял зачем ее использовать второй раз, если таймаут и так завершает контекст, поэтому добавил сюда*

}

func sender(ctx context.Context, c chan<- string) { //функция отправляющая в канал значения
	i := 0
	for {
		select {
		case <-ctx.Done(): //кейс, если контекст вызвал завершение
			return
		default:
			time.Sleep(time.Millisecond * 200)
			i++
			c <- fmt.Sprint("msg", i)
		}
	}
}

func reciever(ctx context.Context, c <-chan string) { //функция принимающая сообщения
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(<-c)

		}
	}
}
