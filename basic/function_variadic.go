package main

import "fmt"

func sumAll(numbers ...int) int { //... dinamic parameter
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := sumAll(2, 34, 21, 34, 4)
	fmt.Println(result)

	slice := []int{4, 4, 123, 323, 45}
	sum := sumAll(slice...) //gunakan ... di akhir jika ingin mengirim slice parameter
	fmt.Println(sum)
}
