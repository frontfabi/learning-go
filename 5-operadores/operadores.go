package main

import "fmt"

func main() {
	// aritméticos
	soma := 2 + 2
	subtracao := 10 - 2
	divisao := 100 / 5
	multiplicacao := 35 * 7
	restoDivisaoMod := 57 % 13

	fmt.Println(soma, subtracao,divisao, multiplicacao, restoDivisaoMod)

	// em go não podemos fazer operações entre tipos diferentes, até mesmo entre int16 e int32

	// atribuição
	var var1 string = "atribuindo com ="
	inferindoComGopher := "inferencia de tipo"

	fmt.Println(var1)
	fmt.Println(inferindoComGopher)

	// relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//lógicos
	v, f := true, false
	fmt.Println(v && f)
	fmt.Println(f || v)
	fmt.Println(!v)
	fmt.Println(!f)

	//unários
	numero := 100
	numero++
	numero += 25
	fmt.Println(numero)
	numero--
	numero -= 30
	fmt.Println(numero)
	numero /= 5
	fmt.Println(numero)
	numero *= 20
	fmt.Println(numero)

	// ternário não existe em Go :'(	pois sua premissa é que só haja uma forma de fazer as coisas
}