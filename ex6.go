package main

import (
	"fmt"
)

func main() {

	var numeros = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 95}
	var check int

	fmt.Print("Digite um número:")
	fmt.Scan(&check)

	encontrado := false
	for _, num := range numeros {
		if num == check {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println("O Array contém o número informado!")
	} else {
		fmt.Println("O Array não contém o número informado!")
	}

	fmt.Println("O numero inserido:", check)
	fmt.Print("Array", numeros)
}
