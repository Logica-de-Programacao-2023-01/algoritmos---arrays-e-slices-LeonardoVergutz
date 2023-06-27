package main

import "fmt"

func main() {
	inteiros := [10]int{10, 29, 234, 23, 65, 82, 15, 23, 65, 92}
	var soma int
	for i, num := range inteiros {
		if i%2 == 0 {
			soma += num
		}
	}
	fmt.Print("O Array original é:", inteiros, "\n")
	fmt.Print("A soma dos números localizados em posições pares é:", soma)
}
