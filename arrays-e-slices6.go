package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var valor int
	fmt.Print("Informe um valor: ")
	fmt.Scan(&valor)

	encontrado := false
	for _, num := range array {
		if num == valor {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println("O valor está presente no array.")
	} else {
		fmt.Println("O valor não está presente no array.")
	}
}
