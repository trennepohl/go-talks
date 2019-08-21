package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var age int
	var name, position string

	for i := 0; i < 2; i++ {
		fmt.Println("This is interaction 1")
		fmt.Println("Insert name")
		in := bufio.NewReader(os.Stdin)
		name, _ = in.ReadString('\n')
		fmt.Println("age")
		fmt.Scanln(&age)
		fmt.Println("position")
		fmt.Scanln(&position)
		fmt.Println(name)
	}

	fmt.Println("Record 2")
	fmt.Println("===========================")
	fmt.Println("Name:", x)
	fmt.Println("Age:", y)
	fmt.Println("Position:", z)

}
