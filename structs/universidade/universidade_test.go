package universidade

import (
	"testing"
)

type AlunoMock struct {
	RetornaMedia func() float32
}

func (a *AlunoMock) CalculaMedia() float32 {
	return a.RetornaMedia()
}

func TestAlunoAprovado(t *testing.T) {
	aluno := &AlunoMock{
		RetornaMedia: func() float32 {
			return 10
		},
	}
	situacao := SituacaoFinal(aluno)

	if situacao != "Aprovado" {
		t.Errorf("Expected Aprovado, got %s", situacao)
	}
}

func TestAlunoReprovado(t *testing.T) {
	aluno := &AlunoMock{
		RetornaMedia: func() float32 {
			return 2
		},
	}
	situacao := SituacaoFinal(aluno)

	if situacao != "Reprovado" {
		t.Errorf("Expected Aprovado, got %s", situacao)
	}
}
