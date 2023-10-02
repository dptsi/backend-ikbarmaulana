package main

import "fmt"

type Alamat struct {
	City, Province, Countery string
}

//func ChangeCountry(alamat Alamat) { // jika menggunakan ini maka hanya variable di function ini yg akan berubah
//	alamat.Countery = "Indonesia"
//}

func ChangeCountry(alamat *Alamat) { // jika menggunakan pointer, maka semua akan berubah
	alamat.Countery = "Indonesia"
}

func main() {
	var adress = Alamat{
		City:     "California",
		Province: "CA",
		Countery: "America",
	}
	fmt.Println(adress)

	ChangeCountry(&adress)

	fmt.Println(adress)
}
