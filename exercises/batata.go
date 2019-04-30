// 2. Crie uma função que calcula a média de notas, a função deve receber 3 parâmetros (int ou float) calcular a média e imprimir no console se o aluno foi aprovado ou não:
// média >= 7, aprovado;
// 5 < média < 7, recuperação;
// média < 5, reprovado.

package main

import (
	"fmt"
)

var ni, nii, niii, res float32

func main() {
	fmt.Println("Pow cara digita as nota ae:")
	fmt.Scan(&ni, &nii, &niii)
	fmt.Println(ni, nii, niii)
	minhamedia := media(ni, nii, niii)
	if minhamedia < 6.0 {
		fmt.Println("hey vc rodou")
	} else {
		fmt.Println("parabens vc é inteligente")
	}
}

func media(ni, nii, niii float32) float32 {
	return (ni + nii + niii) / 3
}
