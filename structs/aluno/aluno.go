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
	return 10
}

type Renan struct {
}

func (r *Renan) CalculaMedia() float32 {
	return 10
}

func (r *Renan) TrazMaterial() {

}
