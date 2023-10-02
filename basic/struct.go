package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var cust Customer

	cust.Age = 12
	cust.Name = "Ikbar"
	cust.Address = "Surabaya"

	fmt.Println(cust)

	ikbar := Customer{ // cara 2
		Name:    "ikbar",
		Address: "Sidoarjo",
		Age:     26,
	}

	fmt.Println(ikbar)

	maulana := Customer{"maulana", "sidoarjo", 26} // cara 3

	fmt.Println(maulana)
}
