package main

import (
	"fmt"
)

// Condicionais
func main() {
	idade := 17

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
	// A variável maiorDeIdade é declarada e inicializada dentro do if, e só existe dentro do escopo do if. Se tentarmos usar esta variável fora do if, teremos um erro de compilação.
}