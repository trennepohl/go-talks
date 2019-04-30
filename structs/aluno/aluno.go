package aluno

//Aluno is something that represts aluno
type Aluno struct {
	Nome  string
	Idade int
	Turma int
	Curso string
	Notas []float32
}

func (a *Aluno) CalculaMedia() float32 {
	return (a.Notas[0] + a.Notas[1] + a.Notas[2]) / float32(len(a.Notas))
}
