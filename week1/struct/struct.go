package main

import "fmt"

/*
	type NameStruct struct {
		Field1 dataType			<- Field, Property
		FieldN otherDataType
	}
	func (NameStruct) NameOfMethod()
	func (n *NameStruct) NameOfMethodWithStructAlias() {
		// Membutuhkan nilai Field1
		var field1 = n.Field1
	}
*/

type Biodata struct {
	Name    string
	Age     uint32
	Hobbies []string
}

type Biodata2 Biodata   // Menurunkan fields saja
type Biodata3 = Biodata // Menurunkan semuanya (field2 + methods)
type Biodata4 struct {
	*Biodata
	Address string
}

func (b *Biodata) GetName() string {
	return b.Name
}

func (b *Biodata) GetHobbies() []string {
	return b.Hobbies
}

func (b *Biodata) AddHobby(hobby ...string) {
	b.Hobbies = append(b.Hobbies, hobby...)
}

func main() {
	biodata := Biodata4{
		Biodata: &Biodata{
			Name:    "John Wick",
			Age:     28,
			Hobbies: []string{"bersepeda", "berlari"},
		},
		Address: "Bondowoso",
	}
	fmt.Println("Get from field:", biodata.Name)
	fmt.Println("Get name:", biodata.GetName())
	fmt.Println("Get hobbies :", biodata.GetHobbies())

	biodata.AddHobby("berenang", "memancing", "berkebun")
	fmt.Println("Get hobbies #2 :", biodata.GetHobbies())
}
