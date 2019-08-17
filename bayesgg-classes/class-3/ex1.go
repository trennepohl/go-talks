package main

import "fmt"

func main() {
	nL := []int{43, 22, 31, 47, 53, 60, 71, 84, 96, 10, 112, 12, 13}
	for _, n := range nL {
		if n%2 != 0 {
			fmt.Printf("%d\n", n)
		}
	}
}
