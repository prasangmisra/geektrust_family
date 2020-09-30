package models

import (
	"geektrust/helpers"
)

//Person struct are the characters of a person in the family tree
type Person struct {
	Name     string
	Gender   string
	Spouce   *Person
	Children []*Person
}

//CreatePerson function will create a new person and return the pointer
func CreatePerson(name string, gender string) *Person {
	var person Person
	person.Name = name
	person.Gender = gender
	return &person
}

//FindChildrenWithName gives the pointer to the children node
func (p *Person) FindChildrenWithName(input string) *Person {
	result := p
	if p.Gender == helpers.GenderMale {
		result = p.Spouce
	}
	for _, value := range result.Children {
		if value.Name == input {
			return value
		}
	}
	result = nil
	return result
}

//FindParent function will try to find the parent node in the family
func (p *Person) FindParent(name string) (*Person, bool) {
	if p.Gender == helpers.GenderMale && p.Spouce != nil {
		return p.Spouce.FindParent(name)
	}
	for _, value := range p.Children {

		if value.Name == name {
			return p, true
		}
		value, result := value.FindParent(name)
		if result {
			return value, result
		}
	}
	return &Person{}, false
}

//IfPersonInFamily function tells whether the person is a part of the family
func (p *Person) IfPersonInFamily(name string) (*Person, bool) {
	if p.Name == name {
		return p, true
	}
	if p.Spouce != nil && p.Spouce.Name == name {
		return p.Spouce, true
	}
	if p.Gender == helpers.GenderMale && p.Spouce != nil {
		p = p.Spouce
	}
	for _, value := range p.Children {
		resultPerson, result := value.IfPersonInFamily(name)
		if result {
			return resultPerson, result
		}
	}

	return &Person{}, false
}
