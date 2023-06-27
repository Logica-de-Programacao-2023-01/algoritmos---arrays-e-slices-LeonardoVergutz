package main

import "fmt"

func main() {
	var inteiros = [10]int{10, 23, 432, 2234, 43, 23, 54, 121, 1643, 98}
	var pares []int
	for _, num := range inteiros {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}
	fmt.Print("O Array original é:", inteiros, "\n")
	fmt.Print("O Array dos numeros pares que estão presentes no Array acima são:", pares)
}
