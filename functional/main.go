package main

import (
	"fmt"
	"strconv"

	"github.com/yoshiken0927/paradigms/functional/model"
	"github.com/yoshiken0927/paradigms/functional/service"
)

func main() {
	fmt.Println(parse(service.PutEmployee("yoshiken")(31)))
}

func parse(e model.Employee) string {
	return e.Name + ":" + strconv.Itoa(e.Age) + ":" + strconv.Itoa(e.Version)
}
