package mapr

import "fmt"

func CreateMap() {
	m := make(map[int]string)
	m[1] = "Sublime"
	m[2] = "Text"

	a := m[1]
	fmt.Println(a)
}

func NestMap() {
	m := make(map[int]map[int]string)
	m[1] = make(map[int]string)
	m[1][1] = "ok"
	fmt.Println(m)
}
