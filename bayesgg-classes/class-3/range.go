package main

import "fmt"

func main() {
	alphabet := []string{"a", "b", "c", "d", "e"}

	for index, letter := range alphabet {
		fmt.Printf("The letter %s is in the index %d", letter, index)
	}

}
