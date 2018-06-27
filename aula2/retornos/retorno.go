package main

import (
	"errors"
	"fmt"
)

func main() {

	err := retornaErro()
	// Como o erro implementa uma interface, é possível fazer uma comparação
	// do seu valor com nulo e o mesmo vale para ponteiros.
	if err != nil {
		fmt.Println("erro retornaErro():", err)
	}

	// Aqui nós recebemos todos os valores retornados.
	texto, numero, booleano := multiplosRetornos()
	fmt.Println(texto, numero, booleano)

	// Aqui nós ignoramos o segundo valor retornado.
	texto2, _, booleano2 := multiplosRetornos()
	fmt.Println(texto2, booleano2)

	joao := retornaStruct()
	fmt.Println(joao.nome)

	x := retornoNomeado(10, 11)
	fmt.Println(x)
}

func retornaErro() error {
	return errors.New("não foi possível fazer X")
}

func multiplosRetornos() (string, int, bool) {
	return "texto", 10, true
}

type pessoa struct {
	nome  string
	idade int
}

// pessoa é um tipo, por isso é possível usar, também, como um retorno.
// caso a pessoa não tenha sido criada dentro da função, será retornada
// uma instância de pessoa com os valores padrões em cada propriedade:
// nome = "" e idade = 0
func retornaStruct() pessoa {
	p := pessoa{
		nome:  "João",
		idade: 10,
	}

	return p

	// Outra maneira, neste caso, é retornar a struct diretamente,
	// sem a necessidade de criar uma variável.

	// return pessoa{
	// 	nome:  "João",
	// 	idade: 10,
	// }
}

// Note que como o tipo dos dois parâmetros são do mesmo tipo (int),
// é possível separá-los por vírgula e colocar o tipo no final, ex.:
// algumaFuncao(nome, sobrenome, endereco string)
// A mesma regra pode ser aplicada no retorno nomeado.
// O retorno nomeado já cria a variável `x`.
// Note também que no final foi usado somente `return`, a função já
// sabe o que precisa ser retornado, no caso, a variável x.
func retornoNomeado(a, b int) (x int) {
	x = a + b
	return
}
