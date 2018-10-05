package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/contratos"
	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/modelos"
)

type Api struct {
	universidade contratos.UniversidadeService
}

func (self *Api) RealizaMatricula(contexto *gin.Context) {
	aluno := modelos.Aluno{}

	err := contexto.BindJSON(&aluno)

	if err != nil {
		contexto.AbortWithError(500, err)
		return
	}

	fmt.Println(aluno)

	self.universidade.RealizaMatricula(aluno)
}

func (self *Api) BuscaAluno(contexto *gin.Context) {

}

func NovaAPI(motor *gin.Engine, universadeService contratos.UniversidadeService) {
	api := &Api{
		universidade: universadeService,
	}
	motor.POST("/aluno", api.RealizaMatricula)
	motor.GET("/aluno", api.BuscaAluno)
}
