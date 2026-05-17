package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint16
}

func main() {
	var u usuario
	fmt.Println(u) // o valor zero de um struct é o valor zero de seus respectivos campos

	u.nome = "Fabi"
	u.idade = 36
	fmt.Println(u)

	endExemplo := endereco{"Rua dos Bobos", 0}

	u2 := usuario{"Fabi", 36, endExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Fabi"}
	fmt.Println(u3)
}