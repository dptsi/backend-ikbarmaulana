package main

import "fmt"

type Filter func(string) string // Type declaration

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "prittt"
	} else {
		return name
	}
}

func main() {
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter) // mengirim parameter berupa function
	sayHelloWithFilter("wkwkwk", filter)
}
