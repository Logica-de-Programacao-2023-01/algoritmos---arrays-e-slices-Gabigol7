package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("Informe o tamanho do slice: ")
	fmt.Scanln(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Printf("Informe o valor do elemento %d: ", i)
		fmt.Scanln(&slice[i])
	}

	fmt.Println("O slice informado é:", slice)

	soma := 0
	for _, valor := range slice {
		soma += valor
	}

	fmt.Println("A soma dos valores é:", soma)
}
