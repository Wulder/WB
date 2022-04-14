package main

import (
	"fmt"
	"sync"
)

func main() {

	var m sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ { //создаем 20 врайтеров, которые пишут в одну мапу из разных горутин
		wg.Add(1)
		go writer(i, &m, &wg)
	}

	usualMap := make(map[int]string)
	wg.Wait()
	m.Range(func(k, v interface{}) bool { //приведение sync.Map к обычному map
		usualMap[k.(int)] = v.(string)
		return true
	})
	fmt.Println(usualMap)

}

func writer(id int, m *sync.Map, wg *sync.WaitGroup) {

	for i := 0; i < 100; i++ {

		m.LoadOrStore(i, fmt.Sprint("from writer", id))

	}
	wg.Done()

}
