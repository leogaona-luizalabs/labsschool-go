package main

// Triatleta interface que declara 3 métodos
type Triatleta interface {
	Corre() string
	Pedala() string
	Nada() string
}

// Componente outro exemplo de interfaces
type Componente interface {
	Renderiza(s string) (string, error)
}

// Velocista ...
type Velocista struct {
	Nome string
}

// Corre ...
func (s Velocista) Corre() string {
	return "Corre!"
}

// Pedala ...
func (s Velocista) Pedala() string {
	return "Pedala!"
}

// Nada ...
func (s Velocista) Nada() string {
	return "Nada!"
}

// IronMan função que recebe uma interface como argumento. Apenas os métodos declarados na interface ficam acessíveis daqui
func IronMan(t Triatleta) {
	t.Corre()
	t.Nada()
	t.Pedala()
}

func mainAula2() {
	joao := Velocista{Nome: "Joao"}
	IronMan(joao)
}
