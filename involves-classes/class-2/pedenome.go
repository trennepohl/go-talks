package main

import (
	"fmt"
)

func main() {
	nome := pedeNome()
}

func pedeNome() string {
	fmt.Println("Por favor digite seu nome")
	var nome string
	fmt.Scanln(&nome)
	return nome
}
