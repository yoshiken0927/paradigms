package repository

import (
	"fmt"

	"github.com/yoshiken0927/paradigms/functional/model"
)

func FindEmployee(name string) model.Employee {
	return model.Employee{
		Name: name,
		Age:  29,
	}
}

func UpdateEmployee(e model.Employee) {
	fmt.Println("updated employee.")
}
