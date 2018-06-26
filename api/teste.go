package api

import "fmt"

func init() {
	fmt.Println("init api")
}

// Pessoa ...
type Pessoa struct {
	Nome      string
	Sobrenome string
	senha     string
}

// GetSenha retorna a senha do tipo `p`.
func (p *Pessoa) GetSenha() string {
	return p.senha
}

// SetSenha define a senha do usuário
func (p *Pessoa) SetSenha(s string) {
	p.senha = s
}
