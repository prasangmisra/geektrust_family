package models

import (
	"fmt"
	"geektrust/helpers"
)

//Family struct is the holder of the family tree
type Family struct {
	Person *Person
}

//Init initializes the family tree
func (f *Family) Init() {
	//Going bottom up
	Yodhan := Person{Gender: helpers.GenderMale, Name: "Yodhan"}

	Dritha := Person{Gender: helpers.GenderFemale, Name: "Dritha"}
	Jaya := Person{Gender: helpers.GenderMale, Name: "Jaya"}

	Dritha.Children = append(Dritha.Children, &Yodhan)
	Dritha.Spouce = &Jaya
	Jaya.Spouce = &Dritha
	Tritha := Person{Gender: helpers.GenderFemale, Name: "Tritha"}
	Vritha := Person{Gender: helpers.GenderMale, Name: "Vritha"}

	Chit := Person{Gender: helpers.GenderMale, Name: "Chit"}
	Amba := Person{Gender: helpers.GenderFemale, Name: "Amba"}

	Amba.Spouce = &Chit
	Chit.Spouce = &Amba
	Amba.Children = append(Amba.Children, &Dritha, &Tritha, &Vritha)

	Ish := Person{Gender: helpers.GenderMale, Name: "Ish"}

	Vila := Person{Gender: helpers.GenderFemale, Name: "Vila"}
	Chika := Person{Gender: helpers.GenderFemale, Name: "Chika"}

	Lika := Person{Gender: helpers.GenderFemale, Name: "Lika"}
	Vich := Person{Gender: helpers.GenderMale, Name: "Vich"}

	Lika.Spouce = &Vich
	Vich.Spouce = &Lika

	Lika.Children = append(Lika.Children, &Vila, &Chika)

	Laki := Person{Gender: helpers.GenderMale, Name: "Laki"}
	Lavnya := Person{Gender: helpers.GenderFemale, Name: "Lavnya"}

	Jnki := Person{Gender: helpers.GenderFemale, Name: "Jnki"}
	Arit := Person{Gender: helpers.GenderMale, Name: "Arit"}

	Jnki.Spouce = &Arit
	Arit.Spouce = &Jnki

	Jnki.Children = append(Jnki.Children, &Laki, &Lavnya)

	Ahit := Person{Gender: helpers.GenderMale, Name: "Ahit"}

	Chitra := Person{Gender: helpers.GenderFemale, Name: "Chitra"}
	Aras := Person{Gender: helpers.GenderMale, Name: "Aras"}

	Chitra.Spouce = &Aras
	Aras.Spouce = &Chitra
	Chitra.Children = append(Chitra.Children, &Jnki, &Ahit)

	Vasa := Person{Gender: helpers.GenderMale, Name: "Vasa"}

	Satvy := Person{Gender: helpers.GenderFemale, Name: "Satvy"}
	Asva := Person{Gender: helpers.GenderMale, Name: "Asva"}

	Satvy.Spouce = &Asva
	Asva.Spouce = &Satvy
	Satvy.Children = append(Satvy.Children, &Vasa)

	Kriya := Person{Gender: helpers.GenderMale, Name: "Kriya"}
	Krithi := Person{Gender: helpers.GenderFemale, Name: "Krithi"}

	Kripi := Person{Gender: helpers.GenderFemale, Name: "Krpi"}
	Vyas := Person{Gender: helpers.GenderMale, Name: "Vyas"}

	Kripi.Spouce = &Vyas
	Vyas.Spouce = &Kripi
	Kripi.Children = append(Kripi.Children, &Kriya, &Krithi)

	Atya := Person{Gender: helpers.GenderFemale, Name: "Atya"}

	Satya := Person{Gender: helpers.GenderFemale, Name: "Satya"}
	Vyan := Person{Gender: helpers.GenderMale, Name: "Vyan"}
	Satya.Spouce = &Vyan
	Vyan.Spouce = &Satya
	Satya.Children = append(Satya.Children, &Asva, &Vyas, &Atya)

	Shan := Person{Gender: helpers.GenderMale, Name: "Shan"}
	Anga := Person{Gender: helpers.GenderFemale, Name: "Anga"}
	Anga.Spouce = &Shan
	Shan.Spouce = &Anga
	Anga.Children = append(Anga.Children, &Chit, &Ish, &Vich, &Aras, &Satya)

	f.Person = &Anga
}

//AddChild function will add the child to the family
func (f *Family) AddChild(motherName string, childName string, gender string) (string, bool) {
	//Check if the child name already exist in the family tree
	_, childExists := f.Person.IfPersonInFamily(childName)
	if childExists {
		return helpers.ResponseChildAdditionFail, false
	}
	//Find the mother node
	mother, motherExists := f.Person.IfPersonInFamily(motherName)
	if !motherExists {
		return helpers.ResponseNotFound, false
	}
	if mother.Gender == helpers.GenderMale {
		return helpers.ResponseChildAdditionFail, false
	}
	//If found, check if female
	//Add the person struct as the child of the mother node
	if motherExists {
		mother.Children = append(mother.Children, CreatePerson(childName, gender))
	}
	return helpers.ResponseChildAdditionSuccess, true
}

//IfPersonInFamily function checks if the person exist in the family tree
func (f *Family) IfPersonInFamily(name string) (*Person, bool) {
	fmt.Println("Inside family, checking ", name)
	if f.Person.Name == name {
		return f.Person, true
	}
	return f.Person.IfPersonInFamily(name)
}

//FindParentUncleAunt checks if there is any parental/maternal uncle/aunt present for the given input name
func (f *Family) FindParentUncleAunt(name string, gender string, isPaternal bool) (*[]Person, bool, string) {
	//1. Find if name exists in the family
	fmt.Println("Checking for ", name)
	person, personExistBool := f.IfPersonInFamily(name)
	fmt.Println("Result is :", personExistBool, person.Name)
	if personExistBool {
		//2. Find father
		//fmt.Println("Person exists")
		parent, parentExistBool := f.FindParent(person.Name)
		fmt.Println("ParentExist bool:", parentExistBool)
		if parentExistBool && isPaternal {
			parent = parent.Spouce
			parentExistBool = parent != nil
		}
		if parentExistBool {
			//3. Find father's siblings
			//fmt.Println("Parent exists,", parent.Name)
			siblings, siblingsBool, _ := f.FindSiblings(parent.Name)
			if siblingsBool {
				//fmt.Println("Siblings exists, and they are", len(*siblings))
				var result []Person
				for _, value := range *siblings {
					//4. Shortlist father's siblings based on gender
					if value.Gender == gender {
						result = append(result, value)
					}
				}
				if len(result) > 0 {
					return &result, true, helpers.ResponseChildAdditionSuccess
				}
				return &result, false, helpers.ResponseNone
			}
		}
		fmt.Println("Returning NONE")
		return &[]Person{}, false, helpers.ResponseNone
	}
	return &[]Person{}, false, helpers.ResponseNotFound

}

//FindInLaw function returns the nodes of the brother/sister in laws of the given node
//gender = male for brother in law, female for sister in law
func (f *Family) FindInLaw(name string, gender string) (*[]Person, bool, string) {
	//1.  Find if name exists in the family
	//fmt.Println("Inside Findinlaw")
	person, personExistBool := f.IfPersonInFamily(name)
	//fmt.Println(personExistBool)
	if personExistBool {
		//fmt.Println("PersonExists and name is:", person.Name)
		var result []Person
		//2.1 Get Spouce node
		spouce := person.Spouce
		if spouce != nil {
			//2.2 If spouce exist, siblings of spouce which are male
			spouceSiblings, spouceSiblingsBool, _ := f.FindSiblings(spouce.Name)
			if spouceSiblingsBool {
				for _, value := range *spouceSiblings {
					if value.Gender == gender {
						result = append(result, value)
					}

				}
			}
		}
		//3.1 Get all female siblings
		siblings, siblingsBool, _ := f.FindSiblings(person.Name)
		if siblingsBool {
			for _, value := range *siblings {
				//3.2 All the male spouce of siblings in the result
				if value.Gender != gender && value.Spouce != nil {
					result = append(result, *value.Spouce)
				}
			}
		}
		if len(result) > 0 {
			return &result, true, helpers.ResponseChildAdditionSuccess
		}
		return &result, false, helpers.ResponseNone
	}

	//Return the result
	return &[]Person{}, personExistBool, helpers.ResponseNotFound
}

//FindChildren returns the nodes of all the children according to gender of the given node
func (f *Family) FindChildren(name string, gender string) (*[]Person, bool, string) {
	//Find the node from name
	person, personExistBool := f.IfPersonInFamily(name)
	if personExistBool {
		if person.Gender == helpers.GenderMale && person.Spouce != nil {
			person = person.Spouce
		}
		//fmt.Println("Checking for children of ", person.Name)
		var result []Person
		for _, value := range person.Children {
			if value.Gender == gender {
				result = append(result, *value)
			}
		}
		if len(result) > 0 {
			return &result, true, helpers.ResponseChildAdditionSuccess
		}
		return &result, false, helpers.ResponseNone
	}
	//Loop through the children according to gender and append to the result
	//Return the resule
	return &[]Person{}, personExistBool, helpers.ResponseNotFound
}

//FindSiblings returns all the siblings of the given node
func (f *Family) FindSiblings(name string) (*[]Person, bool, string) {
	//Find the node from name
	_, personBool := f.IfPersonInFamily(name)
	if personBool {
		//FindParent
		parent, parentResultBool := f.FindParent(name)
		if parentResultBool {
			var siblings []Person
			//If parent found, go through all the children apart from the node itself and append it to the result
			for _, sibling := range parent.Children {
				if sibling.Name != name {
					siblings = append(siblings, *sibling)
				}
			}
			if len(siblings) > 0 {
				return &siblings, true, helpers.ResponseChildAdditionSuccess
			}
			return &siblings, false, helpers.ResponseNone
		}
	}
	//return false
	return &[]Person{}, false, helpers.ResponseNotFound
}

//FindParent returns the parent node of the given node
func (f *Family) FindParent(name string) (*Person, bool) {
	if f.Person.Name == name {
		//fmt.Println("This is the root node, no parents")
		return &Person{}, false
	}
	return f.Person.FindParent(name)
}
