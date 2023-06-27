package main

import "fmt"

func main() {
	// Criação do array de inteiros com 7 elementos
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	// Solicita ao usuário um número
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	// Adiciona o número ao primeiro e último elemento do array
	array[0] += numero
	array[len(array)-1] += numero

	// Imprime o array resultante
	fmt.Println("Array resultante:", array)
}
