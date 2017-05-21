package main

import (
	"time"
)

func main() {
	for i := 0; i < 100000; i++ {
		go func() {
			time.Sleep(5 * time.Second)
		}()
	}

	time.Sleep(10 * time.Second)
}