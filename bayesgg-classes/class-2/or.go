package main

import "fmt"

func main() {
	number := 1
	if number > 0 || number < 10 {
		fmt.Println("The value of the variable number is between\n1 and 10\nTo be honest the value is", number)
	}
}
