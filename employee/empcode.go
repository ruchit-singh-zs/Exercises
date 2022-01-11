package employee

//Create an entity of test

type Employee struct {
	desc string
	age  int
	name string
	id   string
}

//Check the age and compute the details of the test

func CheckAge(desc string, age int, name, id string) (Employee, bool) {
	if age < 22 {
		return Employee{}, false
	}

	return Employee{desc: desc, age: age, name: name, id: id}, true
}
