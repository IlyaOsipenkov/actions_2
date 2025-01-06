package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

func main() {
	var wg sync.WaitGroup
	numRoutines := 1000
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func(i int) {
			defer wg.Done()

			// Добавим задержку для повышения вероятности гонки
			time.Sleep(time.Millisecond * time.Duration(i%10))

			// Увеличим значение счетчика (область гонки)
			counter++
			fmt.Printf("Goroutine %d, Counter: %d\n", i, counter)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
