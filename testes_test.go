package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

// VelocistaMock é uma struct para mock. Para que as funções de um mock possam ser utilizadas à partir desse objeto,
// ele "herda" mock.Mock. Depois de feita essa declaração, basta implementar os métodos da interface na struct
type VelocistaMock struct {
	mock.Mock
}

// Renderiza é a versão mockada do método. O Mock receberá no caso de teste o que será retornado. É recomendável que o mock
// receba o mesmo número de argumentos de retorno no caso de teste. Assim fica fácil inferir se o mock está programado para
// retornar um erro ou uma resposta de sucesso
func (v *VelocistaMock) Renderiza(s string) (string, error) {
	args := v.Called(s)

	// como o método retorna uma string no indice 0 e um erro no indice 1, só precisamos checar se o parâmetro no indice 1
	// (erro) veio nulo. Se não veio, devemos retornar um erro. Do contrário, retornamos uma resposta de sucesso recuperando
	// o valor do índice 0
	erro := args.Get(1)
	if erro != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), nil
}

// Caso de teste. Por padrão precisa começar com Test e ter a seguinte assinatura:
// func TestAlgumaCoisa(t *testing.T)
// Lembrando que arquivos com sufixo _test.go não são incluídos no binário final
func TestRenderiza(t *testing.T) {
	v := VelocistaMock{}
	// nesta linha informamos que o mock, ao ter chamado o método Renderiza, passando como argumento "123", deverá retornar
	// uma string "123" e um nil (ou seja, resposta sem erro)
	v.On("Renderiza", "123").Return("123", nil)

	_, err := v.Renderiza("123")
	if err != nil {
		t.Error("Erro retornado:", err)
		t.Fail()
	}
}

// TestRenderizaErro caso de teste onde o mock está configurado para retornar um erro
func TestRenderizaErro(t *testing.T) {
	v := VelocistaMock{}
	// nesta linha informamos que o mock, ao ter chamado o método Renderiza, passando como argumento "123", deverá retornar
	// um erro (verifique que passamos um erro no índice 1)
	v.On("Renderiza", "123").Return("", errors.New("ocorreu um erro"))

	_, err := v.Renderiza("123")
	if err == nil {
		t.Error("Nenhum erro retornado")
		t.Fail()
	}
}
