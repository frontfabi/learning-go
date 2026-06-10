package main

import "fmt"

func main() {
	
	usuario := map[string]string {
		"nome": "João",
		"sobrenome": "Silva",
		"email": "joao.silva@example.com",
	}

	aluno := map[string]map[string]string {
		"dadosPessoais": {
			"nome": "Maria",
			"sobrenome": "Santos",
			"email": "maria.santos@example.com",
		},
		"curso": {
			"nome": "Engenharia de Software",
			"instituicao": "Universidade XYZ",
		},
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])
	fmt.Println(usuario["email"])
	// Para acessar os campos, não pode usar ponto, como em struct, deve usar colchetes e a chave entre aspas.
	
	aluno["signo"] = map[string]string {
		"nome": "Leão",
	}

	fmt.Println(aluno)
	// Para adicionar um novo campo, basta usar a sintaxe de colchetes e atribuir um valor a ele. No caso, estamos adicionando um campo "signo" ao mapa "aluno", que é um mapa aninhado com uma chave "nome" e o valor "Leão".
}