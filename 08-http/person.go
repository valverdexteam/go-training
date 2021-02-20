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
