package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(i)
			fmt.Println("test cringe")
			wg.Done()
		}()
	}

	wg.Wait()
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
