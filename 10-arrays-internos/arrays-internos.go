package main

import "fmt"

func main() {
	slice3 := make([]float32, 10, 11) // função make: aloca um espaço em memória recebe tipo, tamanho e capacidade

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) // capacidade máxima (opcional)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) 
	fmt.Println(cap(slice3)) // quando estouramos a capacidade máxima, o go realoca o dobro de memória como nova capacidade máxima do slice
}