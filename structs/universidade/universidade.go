package universidade

type Estudante interface {
	CalculaMedia() float32
}

func SituacaoFinal(pessoa Estudante) string {
	media := pessoa.CalculaMedia()

	if media > 7 {
		return "Aprovado"
	}

	return "Reprovado"
}
