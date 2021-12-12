package model

type Employee struct {
	Name    string
	Age     int
	Version int
}

func SetAge(e Employee) func(int) Employee {
	return func(age int) Employee {
		e.Age = age
		return e
	}
}
