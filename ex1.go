package main

import "fmt"

func main() {

	numeros := [3]int{2, 4, 8}
	var soma int

	for _, x := range numeros {
		soma += x
	}

	fmt.Println(" A soma dos elementos dessa Array é", soma)
}
