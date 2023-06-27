package main

import "fmt"

func main() {
	// Criando o slice de inteiros
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Solicitando ao usuário os índices dos elementos para troca
	var index1, index2 int
	fmt.Print("Informe o primeiro índice: ")
	fmt.Scanln(&index1)
	fmt.Print("Informe o segundo índice: ")
	fmt.Scanln(&index2)

	// Verificando se os índices são válidos
	if index1 < 0 || index1 >= len(slice) || index2 < 0 || index2 >= len(slice) {
		fmt.Println("Índices inválidos. Programa encerrado.")
		return
	}

	// Realizando a troca de posição dos elementos
	slice[index1], slice[index2] = slice[index2], slice[index1]

	// Imprimindo o slice resultante
	fmt.Println("Slice resultante:", slice)
}
