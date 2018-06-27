package main

import "fmt"

func main() {
	// Criamos um slice com 6 valores inteiros.
	s := []int{1, 2, 3, -4, 5, -6}

	// Criamos um channel que trabalha somente com tipos inteiros.
	c := make(chan int)

	// Chamamos, em um processo separado do `main` a função soma
	// passando a lista de inteiros e o channel.
	go soma(s, c)

	// A leitura do channel, `<-c` é bloqueante, ou seja, o processo `main`
	// vai ficar parado neste ponto até que algo seja retornado através do channel.
	x := <-c

	fmt.Println(x)

	// Poderíamos ler o retorno do channel diretamente no Println, assim: fmt.Println(<-c)
}

func soma(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	// Somente depois de passar por todos os números eu escrevo a soma no channel.
	c <- total
}
