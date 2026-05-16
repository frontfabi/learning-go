package main

import (
	"errors"
	"fmt"
)

func main() {
	// os tipos int podem ser especificados pelo numero de bits (int8, int16, int32, int64) ou somente int (usa a arquitetura do computador como base)
	var numero int = 1000000000000000000
	fmt.Println(numero)

	// mesma lógica do int, sendo este tipo um inteiro sem sinal
	var numeroSemSinal uint = 6151651984651
	fmt.Println(numeroSemSinal)

	// alias
	var inteiroComAlias rune = 56456456 // rune = int32
	fmt.Println(inteiroComAlias)

	var int8Bits byte = 132 // byte = int8
	fmt.Println(int8Bits)

	// flutuante puro não existe. então, ao declarar numeros decimais, usamos float32 ou float64
	var f32 float32 = 654654654.45
	fmt.Println(f32)

	var f64 float64 = 6546546544965812616214.64
	fmt.Println(f64)

	// sempre atribuir valor com aspas duplas
	var str = "String comum. No Go não existe char"
	fmt.Println(str)

	// ao atribuir texto com aspas simples, é retornado o valor ASCII do caractere informado
	char := 'F'
	fmt.Println(char)

	// valor ZERO: valor atribuído a variavel quando não inicializada
	var naoInicializada int
	fmt.Println(naoInicializada) 

	var isLogged bool = true
	fmt.Println(isLogged)

	// go tem um tipo ERRO e seu valor zero é <nil>
	var erro error = errors.New("Erro interno") // forma correta de declarar um erro
	fmt.Println(erro)
}

//Go não precisa de ; pois o proprio compilador remove