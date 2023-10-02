package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Studied() {
	man.Name = man.Name + " S.Tr.Kom"
}

func main() {
	ikbar := Man{"Ikbar"}

	ikbar.Studied()

	fmt.Println(ikbar)
}
