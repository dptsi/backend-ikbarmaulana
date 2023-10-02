package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)
	//
	//resultInt := result.(int) // panic *tidak bisa asal konversi
	//fmt.Println(resultInt)

	// switch statement untuk mengatasi type data
	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is int")
	default:
		fmt.Println("Ups something wrong")
	}
}
