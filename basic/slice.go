package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[2:4]

	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
	fmt.Println(slice1)
}
