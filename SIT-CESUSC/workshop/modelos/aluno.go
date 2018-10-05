package modelos

type Pessoa struct {
	Nome     string
	CPF      string
	Endereco string
}
type Curso struct {
	Nome         string
	CargaHoraria int
	Professores  []Professor
}
type Aluno struct {
	InfoPessoal Pessoa
	Curso       []Curso
}
type Professor struct {
	Curso       []Curso
	InfoPessoal Pessoa
	Salario     float64
}
