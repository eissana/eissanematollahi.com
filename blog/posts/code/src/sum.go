package main

import (
	"fmt"
	"sum"
)

var a []int

func init() {
	for i := 0; i < 100; i++ {
		a = append(a, 20)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {

	fmt.Printf("sum (serial)   = %d\n", sum.SumSerial(a, fib))
	fmt.Printf("sum (concurrent) = %d\n", sum.SumConcurrent(a, fib))
}
