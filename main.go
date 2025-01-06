package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func(i int) {
			time.Sleep(time.Millisecond * 1) // Добавление задержки
			fmt.Println(i)
			fmt.Println("test cringe")
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
