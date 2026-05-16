package main

import "fmt"

// go não permite criar variáveis não utilizadas. Dá erro no código

func main() {
	var variavel1 string = "Declarando string de forma explícita"
	fmt.Println(variavel1)

	variavel2 := "Inferindo tipo string numa variável"
	fmt.Println(variavel2)

	var (
		variavel3 string = "Criando variaveis em grupo"
		variavel4 string = "Criando mais uma"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Inferindo em grupo", "De novo"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	//go não precisa de variavel auxiliar pra inverter os valores
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5)
	fmt.Println(variavel6)
}