package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Ups. Something wrong")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasilBagi, err := pembagian(81, 8)

	if err == nil {
		fmt.Println(hasilBagi)
	} else {
		fmt.Println(err.Error())
	}
}
