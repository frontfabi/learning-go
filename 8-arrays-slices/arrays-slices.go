package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	// Arrays em Go são tipados também e não podem ter tipos diferentes entre os elementos
	var array1[5] int
	fmt.Println(array1) // imprime 5 itens com valor zero
	
	array1[0] = 16
	fmt.Println(array1, array1[0])

	array2 := [3]string{"A", "B", "C"} // atribuição direta
	fmt.Println(array2)

	//Go não permite incrementar a quantidade de itens depois de definida

	array3 := [...]int{1,2,3,4,5} // define a quantidade de itens de um array pela quantidade de itens atribuídos
	fmt.Println(array3)

	//Slices são como ponteiros para arrays, mas não têm tamanho fixo
	slice1 := []int{7,8,9,4,5,6,70,120,33,58,156,200}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array1))

	slice1 = append(slice1, 28) // cria um slice novo, adicionando o 28
	fmt.Println(slice1)

	slice2 := array3[1:3] // pega uma parte de um array definida pelos indices mencionados e transforma em slice
	fmt.Println(slice2) // o primeiro indice é inclusivo e o segundo é exclusivo
}