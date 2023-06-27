package main

import "fmt"

func main() {
	floats := [10]float64{2, 3, 6, 7, 17, 10, 82, 1, 5, 11}
	var maiorescinco []float64
	fmt.Print("O Array atual é:", floats)
	fmt.Println()
	for _, num := range floats {
		if num > 5 {
			maiorescinco = append(maiorescinco, num)
		}
	}
	fmt.Print("Os números do Array que são maiores que 5 são:", maiorescinco)
}
