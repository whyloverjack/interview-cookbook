package main

type person struct {
	name string
	age  int
}

func parsePerson1() {
	m := make(map[string]*person)
	persons := []person{
		{name: "Tom", age: 18},
		{name: "Jerry", age: 20},
		{name: "Jack", age: 22},
	}
	for _, p := range persons {
		m[p.name] = &p
	}
	for k, v := range m {
		println(k, "=>", v.name, v.age)
	}
}

func parsePerson2() {
	m := make(map[string]*person)
	persons := []person{
		{name: "Tom", age: 18},
		{name: "Jerry", age: 20},
		{name: "Jack", age: 22},
	}
	for i, p := range persons {
		m[p.name] = &persons[i]
	}
	for k, v := range m {
		println(k, "=>", v.name, v.age)
	}
}

func main() {
	parsePerson1()
	parsePerson2()
}
