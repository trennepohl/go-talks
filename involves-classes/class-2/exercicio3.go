package main

import (
	"fmt"
)

func main() {
	var notas []float32
	for i := 1; i <= 3; i++ {
		var nota float32
		fmt.Printf("Digite a nota %d\n", i)
		fmt.Scanf("%f", &nota)
		notas = append(notas, nota)
	}

	media := calculaMedia(notas[0], notas[1], notas[2])
	situacao := verificaSituaçãoAluno(media)
	fmt.Printf("O Aluno foi %s com a média %f", situacao, media)
}

func calculaMedia(n1, n2, n3 float32) float32 {
	return (n1 + n2 + n3) / 3
}

func verificaSituaçãoAluno(media float32) string {
	if media > 7.0 {
		return fmt.Sprintf("Aprovado")

	} else if media <= 5.0 {

		return fmt.Sprintf("Reprovado")
	} else {
		return fmt.Sprintf("Aluno na média")
	}
}
