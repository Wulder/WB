package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	workersCount := readWorkesCount() //Ввод количество воркеров
	messages := make(chan string, 5)  //создание канала с сообщениями для воркеров

	fmt.Println(workersCount)

	ctx, cancel := context.WithCancel(context.Background())

	SetupCloseHandler(cancel) //обработчик нажатия комбинации ctrl + c
	for i := 0; i < workersCount; i++ {
		go worker(i, ctx, messages)
	}

	for {
		for i := 0; i < 5; i++ {
			messages <- getRandomString()
		}

	}

}

func worker(id int, ctx context.Context, messages <-chan string) {
	for msg := range messages {
		select {
		case <-ctx.Done():
			fmt.Println("Workder", id, "work over")
			return
		default:
			time.Sleep(time.Millisecond * 100) // какая-то полезная нагрузка
			fmt.Println("Worker", id, "write:", msg)
		}

	}
}

func SetupCloseHandler(cancel context.CancelFunc) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {

		<-c
		cancel() //вызов завершения контекста (горутины подписаны на этот контекст)
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		time.Sleep(time.Second * 1)
		os.Exit(0)
	}()
}

func getRandomString() string { //генерация рандомной строки
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var result string

	for i := 0; i < 6; i++ {
		result += string(letterRunes[rand.Intn(len(letterRunes))])
	}

	return result
}

func readWorkesCount() int {
	in := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите количество воркеров: ")
	in.Scan()
	i, err := strconv.Atoi(in.Text())
	if err != nil {
		fmt.Println(err.Error())
		return readWorkesCount()
	} else if i == 0 {
		return readWorkesCount()
	}

	return i
}
