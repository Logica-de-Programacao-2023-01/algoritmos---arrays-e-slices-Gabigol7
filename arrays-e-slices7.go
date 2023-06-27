package main

import "fmt"

func main() {
	slice := make([]int, 5)

	fmt.Println("Digite um número inteiro:")
	var numero int
	fmt.Scan(&numero)

	if !contains(slice, numero) {
		slice = append(slice, numero)
	}

	fmt.Println("Slice resultante:", slice)
}

func contains(slice []int, numero int) bool {
	for _, valor := range slice {
		if valor == numero {
			return true
		}
	}
	return false
}
