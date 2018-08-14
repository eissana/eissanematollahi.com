package main

//go:generate ./gen_name.sh

import (
	"fmt"
	"person"
)

func main() {
	fmt.Printf("Person: %+v\n", person.Person{"Eissa", 44})
}
