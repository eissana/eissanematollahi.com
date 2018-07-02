package sum

// import "sync"

func SumSerial(a []int, f func(int) int) int {

	sum := 0
	for _, x := range a {
		sum += f(x)
	}
	return sum
}

func SumConcurrent(x []int, f func(int) int) int {
	entries := make(chan int)

	// x is passed as an argument to goroutine! (Avoid using shared values among goroutines)
	for _, xi := range x {
		go func(xi int) {
			entries <- f(xi)
		}(xi)
	}

	sum := 0
	for range x {
		sum += <-entries
	}
	return sum
}

/*
// The following technique (using wait groups) is used when the number of goroutines is not known a priori
func SumConcurrent(a []int, f func(int) int) int {
	entries := make(chan int)
	var wg sync.WaitGroup

	for _, x := range a {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			entries <- f(x)
		}(x)
	}

	go func() {
		wg.Wait()
		close(entries)
	}()

	sum := 0
	for entry := range entries {
		sum += entry
	}
	return sum
}
*/
