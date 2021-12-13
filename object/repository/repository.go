package repository

import "github.com/yoshiken0927/paradigms/object/model"

func FindEmployee(name string) model.Employee {
	return model.NewEmployee(name, 30)
}

func UpdateEmployee(e model.Employee) model.Employee {
	e.SetVersion(e.Version() + 1)
	return e
}
