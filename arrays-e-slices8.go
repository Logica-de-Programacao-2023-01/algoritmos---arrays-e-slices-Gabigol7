package main

import "fmt"

func main() {
	// Criando o slice de strings com tamanho 8
	slice := make([]string, 8)

	// Solicitando ao usuário que informe um valor
	var valor string
	fmt.Print("Informe um valor: ")
	fmt.Scanln(&valor)

	// Preenchendo o slice com alguns valores para fins de exemplo
	slice[0] = "valor1"
	slice[1] = "valor2"
	slice[2] = "valor3"
	slice[3] = "valor2"
	slice[4] = "valor4"
	slice[5] = "valor2"
	slice[6] = "valor5"
	slice[7] = "valor6"

	// Removendo todas as ocorrências do valor informado pelo usuário
	novoSlice := removeValor(slice, valor)

	// Imprimindo o resultado
	fmt.Println("Slice original:", slice)
	fmt.Println("Slice modificado:", novoSlice)
}

func removeValor(slice []string, valor string) []string {
	novoSlice := make([]string, 0)

	// Percorrendo o slice original
	for _, v := range slice {
		// Verificando se o valor atual é diferente do valor informado
		if v != valor {
			// Adicionando o valor ao novo slice se for diferente
			novoSlice = append(novoSlice, v)
		}
	}

	return novoSlice
}
