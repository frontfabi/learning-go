package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1 // atribuindo valor por cópia

	fmt.Println(var1, var2)

	var1++ // ao incrementar só muda var1
	fmt.Println(var1, var2) // 11 10

	var var3 int
	var pont *int

	var3 = 10
	pont = &var3 // atribuição de ponteiro para var3

	var3++

	fmt.Println(var3, pont, *pont) // valor de var3, entereço de memória de var3 (para onde está apontando), desreferenciação para exibir o valor nesse endereço
}