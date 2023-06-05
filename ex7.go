package main

import (
	"fmt"
)

func main() {
	var numeros = []int{10, 20, 30, 40, 50}
	var novo int

	fmt.Print("Digite um número:")
	fmt.Scan(&novo)

	encontrado := false
	for _, num := range numeros {
		if num == novo {
			encontrado = true
			break
		}
	}
	if encontrado {
		fmt.Println("O número já esta na lista!")
	} else {
		final := append(numeros, novo)
		fmt.Println(final)
	}

}
