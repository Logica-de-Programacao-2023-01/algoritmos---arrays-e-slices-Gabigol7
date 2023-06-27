package main

import "fmt"

func main() {
	// Criação do array de floats com 6 elementos
	array := [6]float64{}

	// Solicita ao usuário que informe um número
	var numero float64
	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)

	// Adiciona o número em todas as posições do array
	for i := 0; i < len(array); i++ {
		array[i] = numero
	}

	// Imprime o array resultante
	fmt.Println("Array resultante:", array)
}
