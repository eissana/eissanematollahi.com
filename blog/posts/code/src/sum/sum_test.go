package sum

import "testing"

var a map[int][]int

func init() {
	n := 10
    a = make(map[int][]int)

	for i := 10; i <= 30; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i] = append(a[i], i)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}


func BenchmarkSum10(b *testing.B) {
	benchmarkSum(10, b)
}

func BenchmarkSum11(b *testing.B) {
	benchmarkSum(11, b)
}

func BenchmarkSum12(b *testing.B) {
	benchmarkSum(12, b)
}

func BenchmarkSum13(b *testing.B) {
	benchmarkSum(13, b)
}

func BenchmarkSum14(b *testing.B) {
	benchmarkSum(14, b)
}

func BenchmarkSum15(b *testing.B) {
	benchmarkSum(15, b)
}

func BenchmarkSum16(b *testing.B) {
	benchmarkSum(16, b)
}

func BenchmarkSum17(b *testing.B) {
	benchmarkSum(17, b)
}

func BenchmarkSum18(b *testing.B) {
	benchmarkSum(18, b)
}

func BenchmarkSum19(b *testing.B) {
	benchmarkSum(19, b)
}

func BenchmarkSum20(b *testing.B) {
	benchmarkSum(20, b)
}

func BenchmarkSum21(b *testing.B) {
	benchmarkSum(21, b)
}

func BenchmarkSum22(b *testing.B) {
	benchmarkSum(22, b)
}

func BenchmarkSum23(b *testing.B) {
	benchmarkSum(23, b)
}

func BenchmarkSum24(b *testing.B) {
	benchmarkSum(24, b)
}

func BenchmarkSum25(b *testing.B) {
	benchmarkSum(25, b)
}

func BenchmarkSum26(b *testing.B) {
	benchmarkSum(26, b)
}

func BenchmarkSum27(b *testing.B) {
	benchmarkSum(27, b)
}

func BenchmarkSum28(b *testing.B) {
	benchmarkSum(28, b)
}

func BenchmarkSum29(b *testing.B) {
	benchmarkSum(29, b)
}

func BenchmarkSum30(b *testing.B) {
	benchmarkSum(30, b)
}


func benchmarkSum(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		// uncomment this to produce serial execution benchmark
		SumSerial(a[n], fib)
		// uncomment this to produce concurrent execution benchmark
		// SumConcurrent(a[n], fib)
	}
}
