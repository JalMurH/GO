package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
	//name string
}

type FullTimeEmployee struct {
	Employee
	Person
}

func (e *Employee) SetId(id int) {
	e.id = id
}

// func (e *Employee) SetName(name string) {
// 	e.name = name
// }

func (e *Employee) GetId() int {
	return e.id
}

// func (e *Employee) GetName() string {
// 	return e.name
// }

// func NewEmployee(id int, name string) *Employee {
// 	return &Employee{
// 		id:   id,
// 		name: name,
// 	}
// }

func main() {
	// e := Employee{}
	// //fmt.Printf("%v", e)
	// e.id = 1
	// e.name = "Jorge"
	// //fmt.Printf("%v", e)
	// e.SetId(5)
	// //fmt.Printf("%v", e)
	// e.SetName("Alejandro")
	// //fmt.Printf("%v", e)
	// fmt.Println(e.GetId())
	// fmt.Println(e.GetName())

	// //constructor forma 2
	// e2 := Employee{
	// 	id:   1,
	// 	name: "Name",
	// }
	// fmt.Printf("%v\n", e2)

	// // contructor forma 3 equivalente a la forma 1 devuelve referencia al valore de e3 se usa el * para acceder al valor
	// e3 := new(Employee)
	// fmt.Printf("%v", *e3)

	// // forma 4 similar a la forma 3 pero prebiamente se definio la funcion que le da valores
	// e4 := NewEmployee(8, "Murillo")
	// fmt.Printf("%v", *e4)

	//Herencia
	e := FullTimeEmployee{}
	e.name = "Jorge"
	e.age = 23
	e.id = 2
	fmt.Printf("%v", e)
}
