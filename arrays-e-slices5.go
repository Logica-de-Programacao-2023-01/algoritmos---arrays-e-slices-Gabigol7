package main

import "fmt"

func main() {
	// Criando a matriz
	matriz := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matriz[i] = make([]int, 2)
	}

	// Solicitando os valores ao usuário
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Digite o valor para a posição [%d][%d]: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	// Imprimindo a matriz resultante
	fmt.Println("Matriz Resultante:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}
}
