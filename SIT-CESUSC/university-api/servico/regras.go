package servico

import (
	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/university-api/contratos"
	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/university-api/modelos"
)

//Universidade é a entidade resposnável pelas regras de negócio
type universidade struct {
	armazenamento contratos.Armazenamento
}

func NovaUniversidade(armazenamentos contratos.Armazenamento) contratos.UniversidadeService {
	return &universidade{
		armazenamento: armazenamentos,
	}
}

func (self *universidade) CalculaMensalidade() {

}

func (self *universidade) FormaAluno() {

}

func (self *universidade) RealizaMatricula(aluno modelos.Aluno) error {
	err := self.armazenamento.Salvar(aluno)
	return err
}
func (self *universidade) TrancaDisciplina() {

}

func (self *universidade) BuscaAlunoPorMatricula(matricula int) {
	self.armazenamento.Buscar()
}
