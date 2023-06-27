package main

import "fmt"

func main() {
	// Criar o slice de inteiros com os valores 1, 2, 3, 4 e 5
	slice := []int{1, 2, 3, 4, 5}

	// Remover o terceiro elemento do slice
	indice := 2
	slice = append(slice[:indice], slice[indice+1:]...)

	// Imprimir o slice resultante
	fmt.Println("Slice resultante:", slice)
}
