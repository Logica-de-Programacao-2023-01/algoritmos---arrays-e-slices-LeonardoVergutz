package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&n)

	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scan(&array[i])
	}

	ordenado := true
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			ordenado = false
			break
		}
	}

	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}
