package main

import "fmt"

func main() {
	var numeros = [7]int{1, 2, 3, 4, 5, 6, 7}
	var num int
	fmt.Print("Digite um número que sera adicionado a primeira e a ultima posição do array:")
	fmt.Scan(&num)

	numeros[0] = num
	numeros[6] = num
	fmt.Print("O array final ficou assim:", numeros)
}
