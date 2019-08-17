package main

import "fmt"

func main() {
	alphabet := []string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(alphabet); i++ {
		fmt.Printf("The letter %s is in the index %d", alphabet[i], i)
	}

}
