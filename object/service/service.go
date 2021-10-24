package service

import (
	"errors"

	"github.com/yoshiken0927/paradigms/object/model"
	"github.com/yoshiken0927/paradigms/object/repository"
)

func GetOlder(name string, age int) (model.Employee, error) {
	e := repository.FindEmployee(name)
	if e == nil {
		return nil, errors.New("employee is not found")
	}

	err := e.SetAge(age)
	if err != nil {
		return nil, err
	}

	repository.UpdateEmployee(e)
	return e, nil
}
