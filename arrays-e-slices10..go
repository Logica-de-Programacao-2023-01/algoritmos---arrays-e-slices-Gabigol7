package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 10) // Cria um slice de inteiros com tamanho 10

	// Preenche o slice com valores aleatórios
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100) // Gera um número aleatório entre 0 e 99
	}

	// Encontra o valor mínimo e máximo no slice
	min := slice[0]
	max := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	// Imprime o valor mínimo e máximo
	fmt.Println("Valor mínimo:", min)
	fmt.Println("Valor máximo:", max)
}
