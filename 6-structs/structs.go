package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	var u usuario
	fmt.Println(u) // o valor zero de um struct é o valor zero de seus respectivos campos

	u.nome = "Fabi"
	u.idade = 36
	fmt.Println(u)

	u2 := usuario{"Fabi", 36}
	fmt.Println(u2)

	u3 := usuario{nome: "Fabi"}
	fmt.Println(u3)
}