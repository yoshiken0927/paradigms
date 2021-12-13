package main

import (
	"fmt"
	"strconv"

	"github.com/yoshiken0927/paradigms/object/model"
	"github.com/yoshiken0927/paradigms/object/service"
)

func main() {
	e, err := service.PutEmployee("yoshiken", 31)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(format(e))
}

func format(e model.Employee) string {
	return e.Name() + ":" + strconv.Itoa(e.Age()) + ":" + strconv.Itoa(e.Version())
}
