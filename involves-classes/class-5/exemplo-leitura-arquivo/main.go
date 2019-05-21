package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Class struct {
	Name      string `json:"nome"`
	Genero    string
	Forca     int
	Agilidade int
}

func main() {

	arquivoGuerreiro, err := os.Open("guerreiro.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer arquivoGuerreiro.Close()

	byteArray, err := ioutil.ReadAll(arquivoGuerreiro)
	if err != nil {
		log.Fatal(err.Error())
	}

	guerreiro := Class{}

	err = json.Unmarshal(byteArray, &guerreiro)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(guerreiro.Forca)
}
