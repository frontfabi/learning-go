package main

import "fmt"

// funções go precisam ter seu retorno tipado também
func somar(n1 int, n2 int) int {
	return n1 + n2
}

// Funções go podem ter mais de um retorno
func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(56, 47)
	fmt.Println(soma)

	// funções também são um tipo
	var fun = func() {
		fmt.Println("Função fun")
	}

	fun()

	fmt.Println(calculosMatematicos(10, 20))

	soma, subtracao := calculosMatematicos(15, 10)
	fmt.Println(soma, subtracao)

	// podemos ignorar qualquer um dos dois retornos usando _ na declaração de variáveis
	_, sub2 := calculosMatematicos(100, 50)
	fmt.Println(sub2)
	so2, _ := calculosMatematicos(150, 50)
	fmt.Println(so2)
}