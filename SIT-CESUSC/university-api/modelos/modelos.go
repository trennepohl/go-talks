package modelos

type Pessoa struct {
	Nome      string
	CPF       string
	Endereco  string
	Idade     int
	Profissao string
}

type Aluno struct {
	InformaecoesPessoais Pessoa
	Turma                string
	Fase                 string
	Curso                Curso
	StatusDoCurso        string
}

type Curso struct {
	Nome         string
	Custo        float64
	CargaHoraria int
}
