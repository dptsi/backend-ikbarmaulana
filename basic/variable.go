package main

import "fmt"

func main() {
	var nama string // basic

	nama = "ikbar"
	fmt.Println(nama)

	nama = "maulana"
	fmt.Println(nama)

	var surname = "Maulana Ikbar" // v1
	fmt.Println(surname)

	name := "Alexander Young" // v2 lebih singkat
	fmt.Println(name)

	var ( // multiple variable
		firstName = "ikbar"
		lastName  = "maulana"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
