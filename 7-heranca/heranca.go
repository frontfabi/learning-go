package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	idade uint8
	altura float32
}

// Em go não existe herança (conceito de OO). Para herdar campos de um struct, passa-se o nome de outro como campo, sem tipar
type estudante struct {
	pessoa 
	curso string
	turno string
}

func main() {
	fmt.Println("Herança")
	
	p1 := pessoa{"Maria", 38, 1.62}
	fmt.Println(p1)

	e1 := estudante{p1,"Engenharia","Manhã"}
	fmt.Println(e1)
}