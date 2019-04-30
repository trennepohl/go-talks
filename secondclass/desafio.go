package main

import (
	. "./trem"
)

func main() {
	meutrem := Trem{Vagoes: 15}
	meutrem.SetMaquinistaName("Jose")

	meusensor := Sensor{}
	minhacancela := Cancela{}

	for i := 0; i < meutrem.Vagoes; i++ {
		if !meusensor.Ativo && i == 0 {
			meusensor.DesativaSensor()
			minhacancela.FechaCancela()
		}

		meutrem.Passagem()

	}
	meusensor.AtivaSensor()
	minhacancela.AbreCancela()

}
