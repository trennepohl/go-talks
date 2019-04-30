package main

import (
	"fmt"

	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/workshop/modelos"
)

type ErroInterfaceVazia struct {
	Message string
}

func (self *ErroInterfaceVazia) Error() string {
	return self.Message
}

type ErroAlunoSemNome struct {
	Message string
}

func (self *ErroAlunoSemNome) Error() string {
	return self.Message
}

func main() {
	// CadastraAluno("Aline", "ADS")
	joaozinho := modelos.Aluno{
		InfoPessoal: modelos.Pessoa{
			CPF:      "9999999-99",
			Endereco: "Planeta Terra",
		},
	}
	err := CadastraAluno(joaozinho)
	if err != nil {
		fmt.Println(err)
	}
}

func CadastraAluno(aluno modelos.Aluno) error {
	if aluno.InfoPessoal.Nome == "" {
		return &ErroAlunoSemNome{Message: "Opa deu erro o aluno nao tem nome"}

	}
	return nil
}
