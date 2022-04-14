package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(4)
	for i := 0; i < 5; i++ {
		go func(id int) {

			fmt.Println(id)
			runtime.LockOSThread()
			time.Sleep(time.Second * 2)

		}(i)
	}
	fmt.Println(runtime.NumCPU(), runtime.GOMAXPROCS(-1))
	time.Sleep(time.Second * 2)
}
