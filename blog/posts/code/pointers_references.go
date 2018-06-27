package main

import "fmt"

type Person struct {
        name string
        height float32
}

func CreatePerson(name string, height float32) *Person {
    // person := new(Person)
    // person.name = name
    // person.height = height
    // return person
        return &Person {
                name: name,
                height: height,
	}
}

func main() {
   person := CreatePerson("Mike", 180.35)
   fmt.Printf("Person is %+v", *person)
}
