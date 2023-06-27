package main

import "fmt"

func main() {
	// Criando o array de inteiros com 10 elementos
	array := [10]int{2, 7, 4, 10, 6, 3, 8, 1, 9, 5}

	// Criando um novo slice contendo apenas os elementos pares
	var slice []int
	for _, value := range array {
		if value%2 == 0 {
			slice = append(slice, value)
		}
	}

	// Imprimindo o novo slice
	fmt.Println(slice)
}
