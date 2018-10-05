package main

import "fmt"

func mostraUltimo(alfabeto []int, corte int) int {
	return alfabeto[corte]
}

func main() {
	var someIntSlice []int = []int{1, 2, 3, 4, 5, 6, 7, 8}

	var n int

	fmt.Println("Bota o n ae mermao")
	fmt.Scanf("%d", &n)
	fmt.Println(mostraUltimo(someIntSlice, n))

}
