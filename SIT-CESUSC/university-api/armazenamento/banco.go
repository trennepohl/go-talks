package armazenamento

import (
	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/university-api/contratos"
	mgo "gopkg.in/mgo.v2"
)

type mongo struct {
	sessao *mgo.Session
}

func NovoArmazenamento(sessao *mgo.Session) contratos.Armazenamento {
	return &mongo{
		sessao: sessao,
	}
}

func (self *mongo) Salvar(objeto interface{}) error {
	sessao := self.sessao.Copy()
	defer sessao.Close()

	conn := sessao.DB("").C("Alunos")
	err := conn.Insert(objeto)
	return err
}

func (self *mongo) Buscar() {

}
