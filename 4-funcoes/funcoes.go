package main

import "fmt"

// funções go precisam ter seu retorno tipado também
func somar(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	soma := somar(56, 47)
	fmt.Println(soma)

	// funções também são um tipo
	var fun = func() {
		fmt.Println("Função fun")
	}

	fun()

}