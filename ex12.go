package main

import "fmt"

func main() {
	inteiros := [5]int{1, 3, 6, 8, 9}
	var multiplostres []int

	for _, num := range inteiros {
		if num%3 == 0 {
			multiplostres = append(multiplostres, num)
		}
	}
	fmt.Println("Os números multiplos de três presentes no array são:", multiplostres)
}
