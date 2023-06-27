package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scanln(&n)

	// Lendo o array
	fmt.Println("Digite os elementos do array:")
	array := lerArray(n)

	// Verificando se o array está ordenado em ordem crescente
	ordenado := verificarOrdenacaoCrescente(array)

	// Imprimindo o resultado
	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}

func lerArray(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scanln(&array[i])
	}
	return array
}

func verificarOrdenacaoCrescente(array []int) bool {
	n := len(array)
	for i := 1; i < n; i++ {
		if array[i] < array[i-1] {
			return false
		}
	}
	return true
}
