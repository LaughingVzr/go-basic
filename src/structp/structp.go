package structp

import "fmt"

type human struct {
	Age int
}

type person struct {
	Name string
	Age  int
}

type pers struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

type per struct {
	string
	int
}

type farmer struct {
	Name string
	human
}

type worker struct {
	Name string
	human
}

func StructBasic() {
	a := person{}
	a.Name = "Laughing"
	a.Age = 25
	fmt.Println(a)
}

func StructPointer() {
	a := &person{
		Name: "Laughing_Vzr",
		Age:  25,
	}
	a.Name = "Vzr"
	a.Age = 26
	fmt.Println(a)
}

func UpdateStruct() {
	a := person{
		Name: "Laughing_Vzr",
		Age:  25,
	}
	fmt.Println(a)
	A(a)
	fmt.Println(a)
	B(&a)
	fmt.Println(a)
}

func A(p person) {
	p.Name = "Laughing"
	p.Age = 15
	fmt.Println("A Person:", p)
}

func B(p *person) {
	p.Name = "Laughing"
	p.Age = 26
	fmt.Println("B Person:", p)
}

func AnonymousStruct() {
	a := struct {
		Name string
		Age  int
	}{
		Name: "Vzr",
		Age:  25,
	}
	fmt.Println(a)
}

func NestAnonymousStruct() {
	a := pers{
		Name: "Vzr",
		Age:  26,
	}
	a.Contact.Phone = "13333333333"
	a.Contact.City = "Beijing"
	fmt.Println(a)
}

func AnonymousColStruct() {
	a := per{"Laughing", 26}
	fmt.Println(a)
}

func MixtureStruct() {
	f := farmer{Name: "Laughing", human: human{Age: 35}}
	w := worker{Name: "Vzr", human: human{Age: 40}}
	fmt.Println(f, w)
	f.Age = 39
	w.Age = 35
	fmt.Println(f, w)
}
