package main

import "fmt"

// Escreva um algoritmo onde o usuário informa:
// Nome do funcionário
// CPF
// Salário
func main() {
	var nome, cpf string
	var salario float32

	fmt.Println("Olá, por favor digite seu primeiro nome:")
	fmt.Scanln(&nome)
	fmt.Println("Agora o seu CPF:")
	fmt.Scanln(&cpf)
	fmt.Println("Por fim o seu salário:")
	fmt.Scanln(&salario)

	fmt.Printf("\nMuito bem, cadastramos você com sucesso\nNome: %s\nCPF: %s\nSalário: %f", nome, cpf, salario)
}
