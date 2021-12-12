package repository

import (
	"github.com/yoshiken0927/paradigms/functional/model"
)

func FindEmployee(name string) model.Employee {
	return model.Employee{
		Name:    name,
		Age:     29,
		Version: 0,
	}
}

func UpdateEmployee(e model.Employee) model.Employee {
	return model.Employee{
		Name:    e.Name,
		Age:     e.Age,
		Version: e.Version + 1,
	}

}
