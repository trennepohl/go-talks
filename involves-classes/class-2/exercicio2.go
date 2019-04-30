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
	verificaSituaçãoAluno(media)
}

func calculaMedia(n1, n2, n3 float32) float32 {
	return (n1 + n2 + n3) / 3
}

func verificaSituaçãoAluno(media float32) {
	if media > 7.0 {
		fmt.Printf("Aprovado, media " + media)
		return
	} else if media <= 5.0 {
		fmt.Println("Reprovado, media %f", media)
		return
	} else {
		fmt.Printf("Aluno na média, media: %f", media)
	}
}
