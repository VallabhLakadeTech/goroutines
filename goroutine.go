package main

import (
	"fmt"
	"time"
)

func simplegoroutine() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Vallabh")
		}
	}()
	time.Sleep(1 * time.Second)
}
