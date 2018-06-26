package main

import (
	"fmt"

	"github.com/luizalabs/labs-school-csc/api"
)

// A função init() é executada antes da função main().
// A função init() das dependências (imports) é executada primeiro.
func init() {
	fmt.Println("main")
}

func main() {
	var p api.Pessoa
	var d api.Pessoa

	d = api.Pessoa{} // Cria uma pessoa com os valores padrões para cada atributo.
	fmt.Println(d.Nome)

	p = api.Pessoa{
		Nome:      "Lucas",
		Sobrenome: "Devito",
	}

	// Faz uso do método público para definir a senha da pessoa `p`
	p.SetSenha("xablau")

	// Map é o mesmo que o hash, recebe uma chave e um valor.
	var m = make(map[string]int)
	m["Lucas"] = 13
	fmt.Println(m["Lucas"])

	// Apesar de também utilizar make, aqui criamos um slice com 10 espaços iniciais,
	// todos inicializados com o valor padrão de um inteiro, 0 (zero).
	var a = make([]int, 10)
	a[0] = 1000

	// A função append retorna uma cópia do slice mais o que foi acrescentado e por
	// isso é preciso fazer uma atribuição.
	a = append(a, 1)

	// Assim como no Python, é possível fazer consulta de slice utilizando este método.
	fmt.Println(a[1:])

	// Chamada de função com parâmetro variádico.
	// teste("Daniel", "Joao", "Paulo")
	teste()
}

// Esta função recebe um parâmetro chamado `nomes` que é variádico, ou seja, ele recebe
// qualquer número de valores do tipo string. Este parâmetro é opcional.
// Apesar de ter um nome diferente, o parâmetro variádico nada mais é do quem uma lista do tipo.
// Ex: `nomes ...string` == `nomes []string` ou `idades ...int` == `idades []int`
func teste(nomes ...string) {
	fmt.Println(nomes)
}
