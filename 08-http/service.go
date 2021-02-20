package main

import (
	"errors"
	"fmt"
)

var (
	persons = []person{
		{
			Id:         1,
			FirstName:  "Pedro",
			LastName:   "Nascimento",
			InitialAge: 25,
		},
		{
			Id:         2,
			FirstName:  "Gerson",
			LastName:   "Souza",
			InitialAge: 35,
		},
		{
			Id:         3,
			FirstName:  "Lucas",
			LastName:   "Smith",
			InitialAge: 6,
		},
	}
)

func listAllPersons() []person {
	return persons
}

func findById(Id int) (person, error) {
	var result person
	IndexOf := indexOf(Id)
	if IndexOf < 0 {
		return result, errors.New(fmt.Sprintf("Couldn't find any person with id %d", Id))
	}
	return persons[IndexOf], nil

}

func add(p person) error {
	if p.Id != 0 {
		return errors.New("You can't assign anything other than 0 value to id. It's an auto-generated value")
	}
	p.Id = generateNextId()
	persons = append(persons, p)
	return nil
}

func update(p person) error {
	if p.Id == 0 {
		return errors.New("0 is not a valid id for update. Please inform the correct id")
	}
	IndexOf := indexOf(p.Id)
	if IndexOf < 0 {
		return errors.New(fmt.Sprintf("Couldn't find any person with id %d", p.Id))
	}
	persons[IndexOf] = p
	return nil
}

func delete(Id int) error {
	if Id == 0 {
		return errors.New("0 is not a valid id for deletion. Please inform the correct id")
	}
	IndexOf := indexOf(Id)
	if IndexOf < 0 {
		return errors.New(fmt.Sprintf("Couldn't find any person with id %d", Id))
	}
	persons = remove(persons, IndexOf)
	return nil
}

func generateNextId() int {
	biggestId := 0
	for _, p := range persons {
		if p.Id > biggestId {
			biggestId = p.Id
		}
	}
	return biggestId + 1
}

func indexOf(Id int) int {
	for i, p := range persons {
		if p.Id == Id {
			return i
		}
	}
	return -1
}

func remove(slice []person, index int) []person {
	return append(slice[:index], slice[index+1:]...)
}
