package ferrovia

type Trem struct {
	nomeDoMaquinista string
	vagoes           int
}

func (trem *Trem) DefineNomeMaquinista(nome string) {
	trem.nomeDoMaquinista = nome
}
