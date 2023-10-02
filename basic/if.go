package main

import "fmt"

func main() {
	age := 21
	name := "maulana"

	if age == 20 {
		fmt.Println("betul")
	} else {
		fmt.Println("salah")
	}

	if length := len(name); length > 5 { // short statement + inisialisasi diawal
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("nama benar")
	}
}
