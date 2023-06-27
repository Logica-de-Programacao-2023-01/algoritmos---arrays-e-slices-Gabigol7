package main

import "fmt"

func main() {
	// Criar o array de inteiros com 3 elementos
	array := [3]int{1, 2, 3}

	// Calcular a soma dos valores armazenados no array
	soma := 0
	for _, valor := range array {
		soma += valor
	}

	// Imprimir a soma dos valores
	fmt.Println("A soma dos valores é:", soma)
}
