package main

import "fmt"

func main() {
	for n := 0; n < 30; n++ {
	    fmt.Printf("Fibonacci(%d)=%d,\t", n, fib1(n))
		fmt.Printf("Fibonacci(%d)=%d\n", n, fib2(n))
	}
}

func fib1(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x + y
	}
	return x
}

func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}
