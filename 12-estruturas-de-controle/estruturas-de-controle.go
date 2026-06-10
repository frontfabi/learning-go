package main

import (
	"fmt"
)

// Condicionais
func main() {
	idade := 18

	if idade >= 16 {
		fmt.Println("Pode votar.")
	} else {
		fmt.Println("Não pode votar.")
	}

	if maiorDeIdade := idade; maiorDeIdade >= 18 {
		fmt.Println("É maior de idade.")
	} else {
		fmt.Println("É menor de idade.")
	}
}