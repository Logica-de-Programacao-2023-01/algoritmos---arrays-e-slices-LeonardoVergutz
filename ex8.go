package main

import (
	"fmt"
	strings2 "strings"
)

func main() {
	var strings = []string{"João", "Carlos", "Maria", "Pedro", "Gerson", "Patrick", "Livia", "Violeta", "Paulo"}
	var letra string

	fmt.Println("Slice inicial:", strings)

	fmt.Print("Digite uma letra:")
	fmt.Scan(&letra)

	novoSlice := []string{}
	for _, palavra := range strings {
		novaPalavra := strings2.ReplaceAll(palavra, letra, "")
		novoSlice = append(novoSlice, novaPalavra)
	}

	fmt.Println("Slice resultante:", novoSlice)

}
