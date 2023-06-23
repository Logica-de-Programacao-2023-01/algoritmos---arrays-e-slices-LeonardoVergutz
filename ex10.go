package main

import "fmt"

func main() {

	var inteiros = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	maior := 0
	menor := 0

	for _, numeros := range inteiros {
		if numeros > maior {
			maior = numeros
		} else if numeros < menor {
			menor = numeros
		}
	}

	fmt.Println("O maior número do array é: ", maior)
	fmt.Println("O menor número do array é: ", menor)
	fmt.Println("O array é:", inteiros)

}
