package main

import (
	"fmt"
	"strconv"
)

var validUPC = "12345678910" //Começa em 1 não em 0

func main() {
	evenNumbersSum := featureFlagNumbersSum(0, validUPC)
	oddNumbersSum := featureFlagNumbersSum(1, validUPC)
	x := (oddNumbersSum + (evenNumbersSum * 3)) % 10
	checkDigit := 10 - x
	fmt.Println(checkDigit)

}

func featureFlagNumbersSum(startPoint int, UPC string) int {
	var sum int
	for i := startPoint; i < len(UPC); i += 2 {
		number, err := strconv.Atoi(string(UPC[i]))
		if err != nil {
			panic(err)
		}
		sum += number
	}
	return sum
}
