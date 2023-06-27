package main

import "fmt"

func main() {
	var primos []int
	var numero int
	fmt.Print("Digite um número:")
	fmt.Scan(&numero)

	if numero > 1 {
		for i := 2; i < numero; i++ {
			primo := true
			for j := 2; j < i; j++ {
				if i&j == 0 {
					primo = false
					break
				}
			}
			if primo {
				primos = append(primos, i)
			}

		}
	}
	fmt.Print("Números que são primos ate o numero informado:", primos)
}
