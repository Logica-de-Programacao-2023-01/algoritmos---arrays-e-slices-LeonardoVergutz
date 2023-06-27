package main

import "fmt"

func main() {
	var numeros = [8]int{10, 20, 30, 63, 45, 87, 3, 12}
	var primeiro, segundo int
	fmt.Print("Digite o primeiro indice que sera trocado:")
	fmt.Scan(&primeiro)
	fmt.Print("Digite o segundo indice que sera trocado:")
	fmt.Scan(&segundo)

	fmt.Print("Slice antes da alteração:", numeros)

	numeros[primeiro], numeros[segundo] = numeros[segundo], numeros[primeiro]

	fmt.Print("Slice após a alteração:", numeros)
}
