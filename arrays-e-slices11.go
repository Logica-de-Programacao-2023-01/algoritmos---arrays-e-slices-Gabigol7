package main

import "fmt"

func main() {
	// Criando o array bidimensional com 2 linhas e 3 colunas
	matriz := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	var linha, coluna int

	// Solicitando ao usuário os índices de linha e coluna
	fmt.Print("Informe o índice da linha (0 ou 1): ")
	fmt.Scanln(&linha)
	fmt.Print("Informe o índice da coluna (0, 1 ou 2): ")
	fmt.Scanln(&coluna)

	// Verificando se os índices estão dentro do intervalo válido
	if linha >= 0 && linha < 2 && coluna >= 0 && coluna < 3 {
		// Acessando o valor armazenado na posição informada
		valor := matriz[linha][coluna]
		fmt.Printf("O valor na posição [%d][%d] é: %d\n", linha, coluna, valor)
	} else {
		fmt.Println("Índices inválidos!")
	}
}
