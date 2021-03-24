package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	p := person{"Jo", 23}
	fmt.Println("p.name:", p.name)
	fmt.Println("p.age:", p.age)

	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})

	// Mesma coisa :D
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 51
	fmt.Println(sp.age)
}
