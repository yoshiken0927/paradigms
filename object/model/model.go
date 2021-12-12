package model

import (
	"errors"
)

type Employee interface {
	Name() string
	Age() int
	SetAge(age int) error
	Salary() int
	SetVersion(version int) error
	Version() int
}

func NewEmployee(name string, age int) Employee {
	return &employeeImpl{
		name: name,
		age:  age,
	}
}

type employeeImpl struct {
	name    string
	age     int
	version int
}

func (e *employeeImpl) Salary() int {
	return e.age * 10000
}

func (e *employeeImpl) Name() string {
	return e.name
}

func (e *employeeImpl) Age() int {
	return e.age
}

func (e *employeeImpl) SetAge(age int) error {
	if e.age >= age {
		return errors.New("input age is less than current age")
	}
	e.age = age
	return nil
}

func (e *employeeImpl) SetVersion(v int) error {
	if e.version >= v {
		return errors.New("input version is less than current version")
	}
	e.version = v
	return nil
}

func (e *employeeImpl) Version() int {
	return e.version
}
