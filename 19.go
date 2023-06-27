package main

import "fmt"

func main() {
	var quantidade1, quantidade2 int

	fmt.Print("Digite a quantidade de elementos do primeiro array: ")
	fmt.Scan(&quantidade1)
	fmt.Print("Digite a quantidade de elementos do segundo array: ")
	fmt.Scan(&quantidade2)

	array1 := make([]int, quantidade1)
	array2 := make([]int, quantidade2)

	for i := 0; i < quantidade1; i++ {
		fmt.Print("Digite o elemento ", i+1, " do primeiro array: ")
		fmt.Scan(&array1[i])
	}
	fmt.Println("O primeiro array:", array1)

	for i := 0; i < quantidade2; i++ {
		fmt.Print("Digite o elemento ", i+1, " do segundo array: ")
		fmt.Scan(&array2[i])
	}
	fmt.Println("O segundo array:", array2)

	tamanhoUniao := quantidade1 + quantidade2
	uniao := make([]int, tamanhoUniao)

	for i := 0; i < quantidade1; i++ {
		uniao[i] = array1[i]
	}

	for i := 0; i < quantidade2; i++ {
		uniao[quantidade1+i] = array2[i]
	}

	fmt.Println("A união dos dois arrays é:", uniao)
}
