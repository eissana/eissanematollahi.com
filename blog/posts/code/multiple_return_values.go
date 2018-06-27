package main

import "fmt"

func MinMax(arr []int) (min, max int, err error) {
    if len(arr) == 0 {
       err = fmt.Errorf("Input slice is empty")
       return
    }
    min, max = arr[0], arr[0]
    for _, a := range arr {
        if a < min {
            min = a
        } else if a > max {
            max = a
        }
    }
    return
}

func main() {
    arr := []int{5,2,8,3,6}
    if min, max, err := MinMax(arr); err != nil {
        fmt.Printf("An error occured: %v", err)
    } else {
	fmt.Printf("Min=%d, Max=%d\n", min, max)
    }

    arr = []int{}
    min, max, err := MinMax(arr)
    if err != nil {
	fmt.Printf("An error occured: %v", err)
    } else {
	fmt.Printf("Min=%d, Max=%d\n", min, max)
    }
}


