package main

import "fmt"

func fullName() (firstName, lastName string) {
	firstName = "Ikbar"
	lastName = "Maulana"

	return
}

func main() {
	fisrtName, lastName := fullName()

	fmt.Println(fisrtName)
	fmt.Println(lastName)
}
