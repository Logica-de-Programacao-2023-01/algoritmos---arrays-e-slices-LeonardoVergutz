package main

import "fmt"

func main() {

	var numeros [6]int

	for i := 0; i < 6; i++ {
		fmt.Print("Digite um número que sera adicionado a posição ", i, ":")
		fmt.Scan(&numeros[i])

	}
	fmt.Println("O Array resultante é:", numeros)
}
