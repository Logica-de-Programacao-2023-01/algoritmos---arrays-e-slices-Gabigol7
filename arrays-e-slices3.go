package main

import "fmt"

func main() {
	array := [4]float64{2.5, 3.2, 1.8, 4.7}
	produto := 1.0

	for _, valor := range array {
		produto *= valor
	}

	fmt.Printf("O produto dos valores no array é: %.2f\n", produto)
}
