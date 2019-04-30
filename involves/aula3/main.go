package main

import (
	"fmt"

	"github.com/thiagotrennepohl/go-talks/involves/aula3/aluno"
	"github.com/thiagotrennepohl/go-talks/involves/aula3/ferrovia"
)

func main() {
	trem := ferrovia.Trem{}
	trem.DefineNomeMaquinista("joao")

	aluno2 := aluno.Aluno{
		Nome: "Joao",
	}

	fmt.Println(aluno2.FazMedia())

}
