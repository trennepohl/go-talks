package models

type Person struct {
	Name string
}

func (p *Person) GetName() string {
	p.Name = "Something"
	return p.Name
}
