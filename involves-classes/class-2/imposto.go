package main

import "fmt"

// Escreva um algoritmo para realizar o calculo de imposto de renda de uma pessoa.
// Deverá ser solicitado nome e salario.
// Apresente o salario buto, liquido, classe de desconto e total descontado do salário.

// té 1.903,98	–	–
// De 1.903,99 até 2.826,65	7,5	142,80
// De 2.826,66 até 3.751,05	15	354,80
// De 3.751,06 até 4.664,68	22,5	636,13
// Acima de 4.664,68	27,5	869,36

const (
	tetoClasse1 = 1903.98
	tetoClasse2 = 2826.65
	tetoClasse3 = 3751.05
	tetoClasse4 = 4664.68
)

var (
	impostoClasse1 = 0.0
	impostoClasse2 = 7.5
	impostoClasse3 = 15.0
	impostoClasse4 = 22.5
	impostoClasse5 = 27.5
)

func main() {
	var nome, cpf string
	var salario float64

	fmt.Println("Olá, por favor digite seu primeiro nome:")
	fmt.Scanln(&nome)
	fmt.Println("Agora o seu CPF:")
	fmt.Scanln(&cpf)
	fmt.Println("Por fim o seu salário:")
	fmt.Scanln(&salario)

	classeDesconto, salarioLiquido := calculaImposto(salario)

	fmt.Printf("Seu salario bruto é de: %f\nSeu salário liquido é de: %f\nO desconto foi de: %f\nA classe de desconto é: %s", salario, salarioLiquido, salario-salarioLiquido, classeDesconto)
}

func calculaImposto(salario float64) (string, float64) {
	switch {
	case salario <= tetoClasse1:
		return "Classe 1", salario - (salario * (impostoClasse1 / 100))
	case salario <= tetoClasse2:
		return "Classe 2", salario - (salario * (impostoClasse2 / 100))
	case salario <= tetoClasse3:
		return "Classe 3", salario - (salario * (impostoClasse3 / 100))
	case salario <= tetoClasse4:
		return "Classe 4", salario - (salario * (impostoClasse4 / 100))
	default:
		return "Classe 5", salario - (salario * (impostoClasse5 / 100))

	}
}
