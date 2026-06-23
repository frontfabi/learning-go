package main

import (
	"fmt"
)

func diaDaSemana(dia int) string {
	//em go, não precisamos da clausula break no final de cada case, pois o switch não "cai" para o próximo case automaticamente, como em outras linguagens. Se quisermos que isso aconteça, podemos usar a palavra-chave fallthrough.
	var diaDaSemana string
	switch dia {
	  case 1:
			diaDaSemana = "Domingo"
		case 2:
			diaDaSemana = "Segunda-feira"
		case 3:
			diaDaSemana = "Terça-feira"
		case 4:
			diaDaSemana = "Quarta-feira"
		case 5:
			diaDaSemana = "Quinta-feira"
		case 6:
			diaDaSemana = "Sexta-feira"
		case 7:
			diaDaSemana = "Sábado"
			fallthrough // o fallthrough faz com que o próximo case seja executado, mesmo que a condição não seja satisfeita. No caso, o case 7 será executado, e o case default também será executado.
		default:
			diaDaSemana = "Dia inválido"
	}
	return diaDaSemana
}
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

	// Switch
	dia := 7
	fmt.Println("Dia da semana:", diaDaSemana(dia))
}