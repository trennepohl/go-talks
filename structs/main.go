package main

import (
	"fmt"

	"github.com/thiagotrennepohl/go-talks/structs/aluno"
	"github.com/thiagotrennepohl/go-talks/structs/universidade"
)

func main() {
	raffa := &aluno.Aluno{
		Nome:  "Raffa",
		Idade: 25,
		Notas: []float32{1, 2, 3},
		Curso: "Como desviar de balas",
		Turma: 12,
	}
	renan := &aluno.Renan{}

	situacao := universidade.SituacaoFinal(raffa)
	fmt.Println(situacao)

}
