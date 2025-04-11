package races

import (
	"fmt"
	"sync"
)

var counter = 0

func read(wg *sync.WaitGroup, mtx *sync.RWMutex) {
	defer wg.Done()
	defer mtx.RUnlock()
	fmt.Println(counter)
}

func increment(wg *sync.WaitGroup, mtx *sync.RWMutex) {
	defer wg.Done()
	defer mtx.Unlock()
	counter++
}

func Run() {
	fmt.Println("Start Counting")
	defer fmt.Println("Finish Counting")
	var wg sync.WaitGroup
	var mtx sync.RWMutex
	for counter <= 10 {
		wg.Add(2)

		mtx.RLock()
		go read(&wg, &mtx)

		mtx.Lock()
		go increment(&wg, &mtx)

		wg.Wait()
	}
}
