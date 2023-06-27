package main

import "fmt"

func main() {
	// Função para verificar se um número é primo
	func isPrime(num int) bool {
		if num <= 1 {
		return false
	}

		for i := 2; i*i <= num; i++ {
		if num%i == 0 {
		return false
	}
	}

		return true
	}

	func main() {
		var n int

		fmt.Print("Digite um número inteiro positivo: ")
		fmt.Scanln(&n)

		primeNumbers := make([]int, 0)

		i := 2
		for len(primeNumbers) < n {
			if isPrime(i) {
				primeNumbers = append(primeNumbers, i)
			}
			i++
		}

		fmt.Printf("Os %d primeiros números primos são: %v\n", n, primeNumbers)
}
