package model

import "errors"

type Employee interface {
	Salary() int
	SetAge(age int) error
}

func NewEmployee(name string, age int) Employee {
	return &employeeImpl{
		name: name,
		age:  age,
	}
}

type employeeImpl struct {
	name string
	age  int
}

func (e *employeeImpl) Salary() int {
	return e.age * 10000
}

func (e *employeeImpl) SetAge(age int) error {
	if e.age >= age {
		return errors.New("input age is less than current age")
	}
	e.age = age
	return nil
}
