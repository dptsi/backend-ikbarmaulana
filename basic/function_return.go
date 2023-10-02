package main

import "fmt"

func getHello(name string) string {
	return "hello " + name
}

func main() {
	result := getHello("wkwkwk")
	fmt.Println(result)
}
