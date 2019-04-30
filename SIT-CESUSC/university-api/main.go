package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/university-api/api"
	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/university-api/armazenamento"
	"github.com/thiagotrennepohl/go-talks/SIT-CESUSC/university-api/servico"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	motorWeb := gin.Default()

	sessaoDeBanco, err := mgo.Dial("mongodb://localhost:27017/universidade")
	if err != nil {
		panic(err)
	}

	camadaDeArmazenamento := armazenamento.NovoArmazenamento(sessaoDeBanco)
	camadaDeRegras := servico.NovaUniversidade(camadaDeArmazenamento)
	api.NovaAPI(motorWeb, camadaDeRegras)

	motorWeb.Run()
}
