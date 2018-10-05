package contratos

import (
	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/university-api/modelos"
)

type UniversidadeService interface {
	RealizaMatricula(modelos.Aluno) error
	CalculaMensalidade()
	TrancaDisciplina()
	FormaAluno()
}

type Armazenamento interface {
	Salvar(interface{}) error
	Buscar()
}
