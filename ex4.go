package main

import "fmt"

func main() {
	var quantidade, num, soma int
	var numeros []int

	fmt.Print("Digite a quantidade de elementos:")
	fmt.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		fmt.Print("Digite um número:")
		fmt.Scan(&num)
		soma += num
		numeros = append(numeros, num)
	}
	fmt.Println("Os números informados são:", numeros)
	fmt.Println("A soma desses números é:", soma)
}
