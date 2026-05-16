// go mod init [nome_do_modulo] inicia um modulo e cria o arquivo go.mod
// go.mod funciona como o package.json

// go build - compila o modulo em um arquivo executável
// go mod tidy atualiza o go.mod, removendo dependencias não utilizadas
package main

import (
	"fmt"
	"modulo/auxiliar"
	
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá, mundo!")
	auxiliar.Escrever()
	
	erro := checkmail.ValidateFormat("frontfab")
	fmt.Println(erro)
}