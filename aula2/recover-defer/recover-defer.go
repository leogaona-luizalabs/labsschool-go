package main

import (
	"fmt"
)

// A função main vai terminar a execução caso um erro seja retornado pela função retornaErro()
// Note que usamos o pacote `os`.
func main() {
	fazerAlgoImportante()

	fmt.Println("Estou na main e a função fazerAlgoImportante() retornou normalmente. Note que a aplicação foi finalizada sem retornar erro.")
}

func fazerAlgoImportante() {
	defer recuperdaDoPanic()

	fmt.Println("chamando causaPanic()")

	causaPanic(0)

	fmt.Println("retornou normalmente de causaPanic()")
}

// causaPanic força um panic quando o parâmetro for igual a 3.
// O panic é uma função que para o todo o fluxo de execução do código e começa a devolver na pilha
// de execução.
// Mesmo no caso do panic, qualquer chamada usando `defer` será executada.
// Note que causaPanic está sendo chamada recursivamente.
func causaPanic(i int) {
	if i > 3 {
		fmt.Println("Causando panic")
		panic("PANIC!!!")
	}

	defer fmt.Println("defer em causaPanic()", i)

	fmt.Println("causaPanic() no valor", i)
	causaPanic(i + 1)
}

// Esta função é responsável por verificar se existe um panic para recuperar.
// Durante a execução normal, sem panic, `recover` vai retornar nil, por isso o bloco `if`.
// Caso exista um panic na rotina, a chamada `recover` vai capturar o valor dado ao panic
// e dar continuidade à execução da aplicação normalmente.
func recuperdaDoPanic() {
	err := recover()
	if err != nil {
		fmt.Println("capturei o panic:", err)
	}

	// Outra maneira de fazer o bloco `if`, com a atribuição e validação.

	// if err := recover(); err != nil {
	// 	fmt.Println("capturei o panic:", err)
	// }
}
