package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o tamanho dos arrays: ")
	fmt.Scanln(&n)

	// Lendo o primeiro array
	fmt.Println("Digite os elementos do primeiro array:")
	array1 := lerArray(n)

	// Lendo o segundo array
	fmt.Println("Digite os elementos do segundo array:")
	array2 := lerArray(n)

	// Somando os dois arrays
	array3 := somarArrays(array1, array2)

	// Imprimindo o terceiro array
	fmt.Println("A soma dos dois arrays é:")
	imprimirArray(array3)
}

func lerArray(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scanln(&array[i])
	}
	return array
}

func somarArrays(array1, array2 []int) []int {
	n := len(array1)
	array3 := make([]int, n)
	for i := 0; i < n; i++ {
		array3[i] = array1[i] + array2[i]
	}
	return array3
}

func imprimirArray(array []int) {
	for _, elemento := range array {
		fmt.Printf("%d ", elemento)
	}
	fmt.Println()
}
