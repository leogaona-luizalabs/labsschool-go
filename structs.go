package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// jsonExemplo é uma constante (string). Em Go não é permitido
// utilizar ' para string e como foi preciso criar uma string
// no formato JSON, contento ", foi necessário usar `, o que
// permite criar uma string com " e quebra de linha.
const jsonExemplo = `{"nome":"Kyojuu Tokusou Jaspion",
		"telefone_fixo":"(999) 999-999-999",
		"idade":45}`

// Pessoa é uma struct com uma coleção de campos (propriedades)
type Pessoa struct {
	Nome     string `json:"nome"`                    // Esta á uma tag, o campo Nome irá aparecer como "nome".
	Telefone string `json:"telefone_fixo,omitempty"` // omitempty é uma tag que omite o campo caso ele seja vazio.
	Idade    int    `json:"idade,omitempty"`
}

// RetornaDados é um método da struct Pessoa
// Todo método precisa de um "receiver" que é o tipo ao qual esta função
// pertence.
// Neste exemplo nós definimos o receiver como  (p Pessoa) porque nós
// queremos manipular alguma propriedade da Pessoa através de "p".
// Caso seja necessário definir um método mas não é necessário alterar
// alguma propriedade, o receiver pode ser definido como (Pessoa).
func (p Pessoa) RetornaDados() string {
	return p.Nome
}

// DaCadeia é um método da struct Pessoa com um "receiver" (p Pessoa)
// porque os atributos de Pessoa estão sendo usados.
// Este método retorna dois valores, uma string e um erro.
// É obrigatório que o método sempre retorne estes dois valores.
// O tipo erro pode ser retornado como "nil" porque ele é uma interface
// e vamos explicar como elas funcionam nas próximas aulas.
// O tipo string pode ser retornado como "" mas não como nil.
func (p Pessoa) DaCadeia() (string, error) {
	if p.Idade >= 18 {
		return p.Telefone, nil
	}
	// Como precisamos retornar um erro, precisamos instanciar
	// um novo erro através da lib "errors" (veja no import)
	return p.Telefone, errors.New("Di menor")
}

// Aniversario é um método da struct Pessoa que não tem nenhum retorno.
// O "receiver" (p *Pessoa) é um ponteiro do tipo Pessoa.
// Neste exemplo é necessário que o "receiver" seja um ponteiro porque
// em Go todos os parâmetros e "receivers" são passados, por padrão,
// como cópia.
// Se quiser, testar, altere o "receiver" deste método para (p Pessoa)
// e veja o resultado, a idade será incrementada somente no contexto
// deste método.
func (p *Pessoa) Aniversario() {
	p.Idade = p.Idade + 1
}

// ponteiroString é uma função (não um método) que recebe, como
// parâmetro um ponteiro do tipo string e retorna uma string.
func ponteiroString(s *string) string {
	return ""
}

// main é o ponto de entrada do programa.
// Existe uma outra função que pode ser usada para executar algo
// antes da função main(), é a função init().
func mainAula1() {
	// superMan é uma instância do tipo Pessoa.
	// Se alguma propriedade não for preenchida, o compilador
	// irá atribuir o valor padrão:
	// int = 0
	// bool = false
	// string = ""
	// struct = {}
	superMan := Pessoa{
		Nome:     "Clark Kent",
		Telefone: "(555) 555-555-555",
		Idade:    2356,
	}

	// json.Marshal é o método do pacote JSON responsável por
	// converter a struct (objeto) em um JSON.
	superManJSON, err := json.Marshal(&superMan)
	if err != nil {
		fmt.Println(err)
	}

	// O retorno da chamada json.Marshal() é uma lista de bytes
	// e por isso é preciso converter para string.
	fmt.Println(string(jsonExemplo))
	fmt.Println(string(superManJSON))

	// jaspion é uma instância do tipo Pessoa.
	// As propriedades podem ser adicionadas usando
	// jaspion.Nome = "Jaspion", etc.
	jaspion := Pessoa{}

	// json.Unmarshal é o método do pacote JSON responsável por
	// converter um JSON em uma struct (objeto).
	err = json.Unmarshal([]byte(jsonExemplo), &jaspion)
	if err != nil {
		fmt.Println(err)
	}

	// O pacote fmt oferece vários métodos para trabalhar com
	// strings.
	// A função Printf permite o uso de verbos para imprimir
	// o valor da variável. Neste exemplo, o uso do %#v faz
	// com que o método retorne tanto os valores da variável
	// jáspion como também o seu tipo e propriedades.
	fmt.Printf("Pessoa: %#v", jaspion)

	// Se você quiser testar o uso do pacote XML como nós fizemos
	// na primeira aula, é preciso trocar todas as chamadas do
	// pacote json por xml, ou seja, substitua todo json. por xml.
	// é preciso também criar uma variável como a jsonExemplo
	// só que com uma string XML.

	// Tente ai e se não der certo nos chame no Slack (@gerep |
	// @leonardo.gaona)
}
