package aluno

type Aluno struct {
	Nome string
	N1   int
	N2   int
	N3   int
}

func (a *Aluno) FazMedia() float32 {
	return float32((a.N1 + a.N2 + a.N3)) / 3
}
