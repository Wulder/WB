package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value uint64
	sync.RWMutex
}

func (c *Counter) Increment() {
	c.RWMutex.Lock()
	c.value++
	c.RWMutex.Unlock()
}

func (c *Counter) Get() uint64 {
	c.RWMutex.RLock()
	defer c.RUnlock()
	return c.value
}

func main() {

	c := Counter{value: 0}

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(c *Counter, id int) {
			c.Increment()
			fmt.Println("Goroutine", id, "-", c.Get())
			wg.Done()
		}(&c, i)
	}

	wg.Wait()
	fmt.Println("result -", c.Get())
}
