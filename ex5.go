package main

import "fmt"

func main() {
	var matriz [3][2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("Digite o valor para a linha ", i, " e para a coluna ", j, ":")
			fmt.Scan(&matriz[i][j])
		}
	}
	fmt.Println("A matriz final é essa:", matriz)
}
