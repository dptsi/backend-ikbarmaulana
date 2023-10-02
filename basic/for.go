package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	if counter%2 == 0 {
	//		fmt.Println(counter, "genap")
	//	} else {
	//		fmt.Println(counter, "ganjil")
	//	}
	//
	//	counter++
	//}

	slice := []string{
		"ikbar",
		"maulana",
		"surabaya",
		"119237812638",
	}

	for _, value := range slice { // _ untuk skip value
		//fmt.Println("index ", i, "=", value)
		fmt.Println(value)
	}
}
