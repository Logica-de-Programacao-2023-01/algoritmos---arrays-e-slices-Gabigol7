package main

import "fmt"

func main() {
	// Criando o array de floats com 10 elementos
	array := [10]float64{2.3, 7.8, 4.5, 10.2, 6.1, 3.9, 8.7, 1.2, 9.6, 5.4}

	// Criando um novo slice contendo apenas os elementos maiores que 5
	var slice []float64
	for _, value := range array {
		if value > 5 {
			slice = append(slice, value)
		}
	}

	// Imprimindo o novo slice
	fmt.Println(slice)
}
