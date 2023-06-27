package main

import "fmt"

func main() {
	matriz := [2][3]int{{10, 20, 30}, {40, 50, 60}}
	var indicel, indicec int
	fmt.Println(matriz)
linha:
	fmt.Print("Digite um indice de linha:")
	fmt.Scan(&indicel)
	if indicel > 1 {
		fmt.Println("Número invalido!")
		goto linha
	}

coluna:
	fmt.Print("Digite um indice de coluna:")
	fmt.Scan(&indicec)

	if indicec > 2 {
		fmt.Println("Número invalido!")
		goto coluna
	}

	fmt.Println(matriz[indicel][indicec])

}
