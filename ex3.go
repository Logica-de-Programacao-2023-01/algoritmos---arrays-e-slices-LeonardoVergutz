package main

import "fmt"

func main() {
	numeros := [4]float64{2.5, 4.5, 6.5, 8.5}
	var multiplicacao float64 = 1

	for _, x := range numeros {
		multiplicacao *= x
		fmt.Println(x)
	}
	fmt.Println(multiplicacao)
}
