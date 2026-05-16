package auxiliar

import "fmt"

// definir uma função com letra maiúscula informa que ela é pública, já que go não possui public/private
// se a função tem letra minuscula, ela fica private
// quando temos uma função pública, devemos ter comentários em cima dela para documantá-la
// go get [url_do_modulo] instala módulos externos
func Escrever() {
	fmt.Println("Escrevendo direto do Auxiliar")
	escrever2()
}