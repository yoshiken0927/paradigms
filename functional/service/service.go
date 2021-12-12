package service

import (
	"github.com/yoshiken0927/paradigms/functional/model"
	"github.com/yoshiken0927/paradigms/functional/repository"
)

func PutEmployee(name string) func(int) model.Employee {
	return func(age int) model.Employee {
		return repository.UpdateEmployee(model.SetAge(repository.FindEmployee(name))(age))
	}
}

// func Callback(e model.Employee) func(func(model.Employee)) {
// 	return func(f func(model.Employee)) {
// 		f(e)
// 	}
// }
