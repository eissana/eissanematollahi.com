package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	// runtime.GOMAXPROCS(4)
	go spinner(100 * time.Millisecond)
    counter(10, 500 * time.Millisecond)
}

func counter(n int, delay time.Duration) {
	for i := 1; i <= n; i++ {
		fmt.Printf("\r\t%d", i)
		time.Sleep(delay)
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
