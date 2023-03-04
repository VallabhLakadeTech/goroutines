package main

import (
	"fmt"
	"sync"
)

func waitgroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Vallabh")
		}
		wg.Done()
	}()
	wg.Wait()
}
