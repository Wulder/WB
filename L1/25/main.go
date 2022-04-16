package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	sleep(5000)
	fmt.Println("end")
}

func sleep(d int) {
	<-time.After(time.Millisecond * time.Duration(d))
}
