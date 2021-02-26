package main

import (
	"fmt"
)

type empActivity interface {
	assignProject(project string)
	applyLeave(i int)
}
type Person struct {
	firstName string
	lastName  string
	age       int
	address   string
}
type Employee struct {
	Person
	empNo        string
	dept         string
	project      string
	leaveApplied int
}

func (p Person) Print() {
	fmt.Printf("\n Person Detail  {First Name: %s, Lastname: %s, Age: %d, Address: %s}", p.firstName, p.lastName, p.age, p.address)
}
func (e Employee) Print() {
	e.Person.Print()
	fmt.Printf("\n Employee Detail  {empNo: %s, dept: %s project: %s leaves: %d}", e.empNo, e.dept, e.project, e.leaveApplied)
}
func (e Employee) assignProject(project string) {
	e.project = project
}

// func (e Employee) applyLeave(i int) {
// 	e.leaveApplied = i
// }

// func assign(e interface {}, p Project){
// 	e.assignProject(p)
// }

func main() {
	var per1 Person
	per1 = Person{firstName: "Anand", lastName: "Babu", age: 20, address: "Bangalore"}
	//per1.Print()
	var emp2 Employee
	emp2.empNo = "201"
	emp2.dept = "Healthcare"
	emp2.Person = per1
	emp2.Print()
	// var per2 Person = emp2.Person
	// per2.Print()
	var empA1 empActivity
	empA1 = emp2
	empA1.applyLeave(2)

}
