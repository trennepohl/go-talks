package universidade

type Estudante interface {
	CalculaMedia() float32
<<<<<<< HEAD
	TrazMaterial()
=======
>>>>>>> ab341018a63bd802cd1d901f9679c7dd3abe1085
}

func SituacaoFinal(pessoa Estudante) string {
	media := pessoa.CalculaMedia()

	if media > 7 {
		return "Aprovado"
	}

	return "Reprovado"
}
