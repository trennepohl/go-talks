package main

import "fmt"

func main() {
	calculaGastoCigarro(6.0, 15.0, 8.75)
}

func calculaGastoCigarro(qtCigarros float32, qtAnos float32, precoCigarro float32) {
	precoUnitarioCigarro := precoCigarro / 20
	dias := qtAnos * 365
	cigarrosFumados := qtCigarros * dias
	gasto := cigarrosFumados * precoUnitarioCigarro
	fmt.Printf("O gasto de cigarros até agora é %f", gasto)

}
