package main

import (
	"bytes"
	"fmt"
)

func main(){

	var numeros [10]int {10,20,30,40,50,60,70,80,90,100}
	var check int

	fmt.Print("Digite um número:")
	fmt.Scan(&check)


	if range numeros = check {
		fmt.Print("O Array contém o número informado!")
	}else{
		fmt.Print("O Array não contém o número informado!")
	}

	fmt.Println("O numero inserido:",check)
	fmt.Print("Array",numeros)
}
