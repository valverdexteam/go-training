package main

type person struct {
	Id         int
	FirstName  string
	LastName   string
	InitialAge int
	CurrentAge int
}

func NewPerson(Id int, FirstName string, LastName string, InitialAge int) *person {
	return &person{Id, FirstName, LastName, InitialAge, InitialAge}
}

func (p *person) GetOld() {
	p.CurrentAge++
}

func (p *person) GetMoreOld(years int) {
	for i := 1; i <= years; i++ {
		p.GetOld()
	}
}
