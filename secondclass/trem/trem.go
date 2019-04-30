package trem

import (
	"fmt"
	"time"
)

type Trem struct {
	Vagoes     int
	maquinista string
}

type Sensor struct {
	Ativo bool
}

func (s *Sensor) AtivaSensor() {
	s.Ativo = true
	fmt.Println("Sensor ativado!")
}

func (s *Sensor) DesativaSensor() {
	s.Ativo = false
	fmt.Println("Sensor desativado!")
}

func (t *Trem) SetMaquinistaName(name string) {
	t.maquinista = name
	fmt.Println("O nome do maquinista é " + t.maquinista)

}

type Cancela struct {
	Status bool
}

func (c *Cancela) AbreCancela() {
	c.Status = true
	fmt.Println("Cancela aberta!")
}

func (c *Cancela) FechaCancela() {
	c.Status = false
	fmt.Println("Cancela fechada!")
}

func (t *Trem) Passagem() {
	time.Sleep(1 * time.Second)
	fmt.Println("Estoy passando...")
}
