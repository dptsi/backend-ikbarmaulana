package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayyHello(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var ikbar Person
	ikbar.Name = "ikbar"

	sayyHello(ikbar)
}
