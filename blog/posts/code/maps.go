package main

import "fmt"

func main() {
//    heights := map[string]float32 {
//	    "Sarah": 162.42
//    }
    heights := make(map[string]float32)
//    var heights map[string]float32
    heights["Mike"] = 180.35

    person := "Mike"

    if height, ok := heights[person]; !ok {
        fmt.Printf("Height of %s is not available.", person)
    } else {
        fmt.Printf("Height of %s: %0.2f", person, height)
    }
}
