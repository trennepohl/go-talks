package main

import "fmt"

func main() {

	var output = "\n"
	counter := 1
	for {
		var name, profession string
		var age int
		fmt.Println("Please inform the name")
		fmt.Scanln(&name)
		fmt.Println("Please inform the age")
		fmt.Scanln(&age)
		fmt.Println("Please inform the profession")
		fmt.Scanln(&profession)

		output += fmt.Sprintf("Record %d\n================\nName:%s\nAge: %d\nProfession: %s\n---\n", counter, name, age, profession)
		counter++

		var answer string
		fmt.Println("Do you want to continue to add people? (yes|no)")
		fmt.Scanln(&answer)
		if answer != "yes" {
			break
		}

	}
	fmt.Println(output)
}
