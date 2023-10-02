package main

import "fmt"

type Cust struct {
	Name, Address string
	Age           int
}

func (cust Cust) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", cust.Name)
}

func main() {
	ikbar := Cust{
		Name:    "Ikbar",
		Address: "Surabaya",
		Age:     26,
	}

	ikbar.sayHi("Maulana")
}
