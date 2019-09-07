package main

import (
	"fmt"

	"github.com/thiagotrennepohl/go-talks/bayesgg-classes/class-4/models"
	"github.com/thiagotrennepohl/go-talks/bayesgg-classes/class-4/ui"
)

func main() {
	someone := models.Person{Name: "Thiago"}
	fmt.Println(someone.GetName())
	fmt.Println(someone.Name)
	ui.ShowGreetingsMessage()
	ui.ShowGreetingsMessage()
	ui.ShowGreetingsMessage()
	ui.ShowGreetingsMessage()

	if ui.IsPlayerReady() {
		continue the game
	}

}
