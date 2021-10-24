package service

import (
	"github.com/yoshiken0927/paradigms/functional/model"
	"github.com/yoshiken0927/paradigms/functional/repository"
)

func UpdateEmployee(name string) func(int) model.Employee {
	return func(age int) model.Employee {
		return SetAge(repository.FindEmployee(name))(age)
	}
}

func SetAge(e model.Employee) func(int) model.Employee {
	return func(age int) model.Employee {
		e.Age = age
		return e
	}
}

func Callback(e model.Employee) func(func(model.Employee)) {
	return func(f func(model.Employee)) {
		f(e)
	}
}
