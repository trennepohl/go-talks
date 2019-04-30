package main

import (
	"fmt"
)

func main() {
	fmt.Println("Olá")
	nome := "thiago"

	imprimeNomeHacker("Banana")

}

func imprimeNomeHacker(nome string) {
	fmt.Println("Meu nome hacker é ", nome)
}
