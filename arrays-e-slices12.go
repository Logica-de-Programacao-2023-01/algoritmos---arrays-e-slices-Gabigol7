package main

import "fmt"

func main() {
	// Criando o array de inteiros com 5 elementos
	array := [5]int{1, 3, 6, 8, 9}

	// Criando um novo slice para armazenar os elementos múltiplos de 3
	var slice []int

	// Percorrendo o array e adicionando os elementos múltiplos de 3 ao slice
	for _, num := range array {
		if num%3 == 0 {
			slice = append(slice, num)
		}
	}

	// Imprimindo o novo slice
	fmt.Println(slice)
}
