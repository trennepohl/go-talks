package main

import "fmt"

func main() {
	var numbers []int

	fmt.Println("At anytime type 5709 to proceed to the result")
	for {
		counter := 1
		var number int
		fmt.Printf("Type in the %d number\n", counter)
		fmt.Scanln(&number)
		if number == 5709 {
			break
		}
		numbers = append(numbers, number)
	}

	var operation string
	fmt.Println("Now select the operation:\n* for multiplication\n+ for sum")
	fmt.Scanln(&operation)

	var result int
	switch operation {
	case "*":
		result = 1
		for _, number := range numbers {
			result *= number
		}
	case "+":
		for _, number := range numbers {
			result += number
		}

	}

	fmt.Println("The result is ", result)
}
