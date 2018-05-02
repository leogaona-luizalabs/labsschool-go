package main

import (
	"fmt"
	"sync"
)

func concorrenciaComWaitGroup() {
	// waitGroup funciona como um semáforo. Informamos a ele a quantidade de tarefas pendentes através da função Add
	// e então, a cada finalização de tarefa, decrementamos o valor de tarefas pendentes com a função Done.
	//
	// Ao utilizamos a função Wait(), a routine atual fica presa até que todas as tarefas tenham sido executadas (ou seja,
	// o valor do waitgroup seja 0 novamente)
	var wg sync.WaitGroup

	inteiros := []int{1, 2, 3, 4}
	// adicionamos o tamanho do array no waitgroup, uma vez que desejamos processar cada um dos elementos do array em paralelo
	wg.Add(len(inteiros))
	for _, v := range inteiros {
		go func(wg *sync.WaitGroup, v int) {
			// não podemos esquecer de chamar wg.Done(). Do contrário, o processo entrará em deadlock pois ficará dormindo
			// para sempre
			defer wg.Done()
			fmt.Println(v * v)
		}(&wg, v)
	}

	// esta linha ficará bloqueada até que todos os wg.Done sejam chamados
	wg.Wait()
	fmt.Println("for paralelo com waitgroup executado com sucesso")
}

func concorrenciaComChannels() {
	// channels possuem um funcionamento muito parecido com filas. Existem 2 tipos de channel: buffered e o unbuffered.
	// a diferença entre os dois é que no buffered já sabemos a quantidade de mensagens que podem cair no channel, e o envio
	// de mensagens não é blocante. No unbufered, o emissor da mensagem fica preso até que alguma outra goroutine leia a
	// mensagem enviada
	inteiros := []int{1, 2, 3, 4}
	c := make(chan int)

	for _, v := range inteiros {
		go func(c chan int, v int) {
			// cada goroutine aberta irá processar o quadrado do número e enviará o resultado para o channel
			c <- v * v
		}(c, v)
	}

	for i := 0; i < len(inteiros); i++ {
		fmt.Println(<-c)
	}
	fmt.Println("for paralelo com channels executado com sucesso")

	// agora faremos novamente o loop, mudando a forma como escutamos as mensagens no channel
	for _, v := range inteiros {
		go func(c chan int, v int) {
			// cada goroutine aberta irá processar o quadrado do número e enviará o resultado para o channel
			c <- v * v
		}(c, v)
	}

	// é possível fazer loops com range em um canal. O loop continuará sendo executado até que o channel seja fechado.
	// portanto, precisamos de um controle sobre a quantidade de mensagens que já foram recebidas
	processados := 0
	for v := range c {
		fmt.Println(v)
		processados++

		// se já recebemos 4 mensagens (tamanho do array), fechamos o channel para que o loop seja quebrado e o fluxo
		// do programa continue
		if processados == len(inteiros) {
			close(c)
		}
	}

	fmt.Println("for paralelo com channels executado com sucesso novamente")
}
