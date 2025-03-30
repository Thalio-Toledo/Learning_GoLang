package main

import (
	"fmt"
)

func imprimirDados(nome string, idade int) {
	fmt.Printf("%s tem %d anos.", nome, idade)
}

func soma(n, m int) int {
	return n + m
}

func main() {
	imprimirDados("Thalio", 27)
	fmt.Println("\n A soma é", soma(3, 5))
}
