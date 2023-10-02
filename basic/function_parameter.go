package main

import "fmt"

func sayHelloTo(fisrtName string, lastName string) {
	fmt.Println("hello to", fisrtName, lastName)
}

func main() {
	sayHelloTo("ikbar", "maulana")
}
