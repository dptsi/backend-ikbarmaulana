package main

import "fmt"

func getFullname() (string, string) {
	return "ikbar", "maulana"
}

func main() {
	//firstName, lastName := getFullname()
	firstName, _ := getFullname() // _ untuk mengambil satu return
	fmt.Println(firstName)
}
