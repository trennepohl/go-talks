package main

import "fmt"

func main() {
	if "gopher" != "Gopher" && "gopheR" != "Gopher" {
		fmt.Println("The result is TRUE because gopher\nand gopheR are different than Gopher")
	}
}
