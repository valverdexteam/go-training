package main

import (
	"testing"
)

func TestShouldAddNewPerson(t *testing.T) {
	initialLength := len(persons)
	expectedLength := initialLength + 1
	personToAdd := *NewPerson(0, "James", "Luis", 19)
	error := add(personToAdd)
	if error != nil {
		t.Error(error)
	}
	actualLength := len(persons)
	if actualLength != expectedLength {
		t.Errorf("expected %d but got %d", expectedLength, actualLength)
	}
}
func TestShouldNotAddNewPersonWithValidId(t *testing.T) {
	personToAdd := *NewPerson(3, "James", "Luis", 19)
	error := add(personToAdd)
	if error == nil {
		t.Error("It should return an error but it didn't ")
	}
}
func TestShouldReturnPedro(t *testing.T) {
	expectedPerson := persons[0]
	pedro, error := findById(1)
	if error != nil {
		t.Error("It should return the person but returned an error")
	}
	if pedro.Id != expectedPerson.Id {
		t.Errorf("Returned object should have id = %d but was %d", expectedPerson.Id, pedro.Id)
	}
	if pedro.FirstName != expectedPerson.FirstName {
		t.Errorf("Returned object should have first name = %s but was %s", expectedPerson.FirstName, pedro.FirstName)
	}
	if pedro.LastName != expectedPerson.LastName {
		t.Errorf("Returned object should have last name = %s but was %s", expectedPerson.LastName, pedro.LastName)
	}
}

func TestShouldReturnAnErrorWithInvalidId(t *testing.T) {
	person, error := findById(99)
	_ = person
	if error == nil {
		t.Error("It should return an error but it didn't ")
	}
}

func TestShouldUpdatePedro(t *testing.T) {
	pedro := persons[0]
	expectedPerson := person{
		Id:         pedro.Id,
		FirstName:  pedro.FirstName + "",
		LastName:   pedro.LastName + "",
		InitialAge: pedro.InitialAge,
		CurrentAge: pedro.CurrentAge,
	}
	expectedPerson.FirstName = pedro.FirstName + "updated"
	error := update(expectedPerson)
	if error != nil {
		t.Error(error)
	}
	pedro = persons[0]
	if pedro.Id != expectedPerson.Id {
		t.Errorf("Returned object should have id = %d but was %d", expectedPerson.Id, pedro.Id)
	}
	if pedro.FirstName != expectedPerson.FirstName {
		t.Errorf("Returned object should have first name = %s but was %s", expectedPerson.FirstName, pedro.FirstName)
	}
	if pedro.LastName != expectedPerson.LastName {
		t.Errorf("Returned object should have last name = %s but was %s", expectedPerson.LastName, pedro.LastName)
	}
}

func TestShouldNotUpdateWithInvalidId(t *testing.T) {
	pedro := persons[0]
	expectedPerson := person{
		Id:         99,
		FirstName:  pedro.FirstName + "",
		LastName:   pedro.LastName + "",
		InitialAge: pedro.InitialAge,
		CurrentAge: pedro.CurrentAge,
	}
	error := update(expectedPerson)
	if error == nil {
		t.Error("It should return an error but it didn't ")
	}
}
func TestShouldReturnAnErrorOnDeleteWithInvalidId(t *testing.T) {
	error := delete(99)
	if error == nil {
		t.Error("It should return an error but it didn't ")
	}
}

func TestShouldDeletePedro(t *testing.T) {
	pedro := persons[0]
	error := delete(pedro.Id)
	if error != nil {
		t.Error(error)
	}
	shouldNotBePedro := persons[0]
	if pedro.Id == shouldNotBePedro.Id {
		t.Errorf("Expected to be %d but was %d", pedro.Id, shouldNotBePedro.Id)
	}
}
