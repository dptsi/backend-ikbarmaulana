package main

import "fmt"

type Address struct {
	City, Province, Countery string
}

func main() {
	address1 := Address{"Sidoarjo", "Jatim", "Indonesia"}
	address2 := &address1

	address2.City = "Surabaya"

	//address2 = &Address{"Jakarta", "DKI", "Indo"} // membuat memory baru
	*address2 = Address{"Jakarta", "DKI", "Indo"} // merubah semua variabel ke memory 2

	fmt.Println(address1) // address 1 tidak berubah
	fmt.Println(address2)
}
