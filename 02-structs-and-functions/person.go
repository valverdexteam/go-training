package main

import "fmt"

func main() {

	pedro := NewPerson("Pedro", "Nascimento", 25)
	pedro.GetMoreOld(50)
	Print(pedro)

	gerson := NewPerson("Gerson", "Souza", 35)
	for i := 1; i <= 15; i++ {
		gerson.GetOld()
	}
	Print(gerson)

	lucas := NewPerson("Lucas", "Smith", 6)
	for i := 1; i <= 6; i++ {
		lucas.GetMoreOld(3)
	}
	Print(lucas)

}

type person struct {
	FirstName  string
	LastName   string
	InitialAge int
	CurrentAge int
}

func NewPerson(FirstName string, LastName string, InitialAge int) *person {
	return &person{FirstName, LastName, InitialAge, InitialAge}
}

func (p *person) GetOld() {
	p.CurrentAge++
}

func (p *person) GetMoreOld(years int) {
	for i := 1; i <= years; i++ {
		p.GetOld()
	}
}

func Print(p *person) {
	fmt.Println(fmt.Sprintf("%s %s arrived with %d years old, but now he(she)'s %d years old. The difference is %d", p.FirstName, p.LastName, p.InitialAge, p.CurrentAge, p.CurrentAge-p.InitialAge))
}
