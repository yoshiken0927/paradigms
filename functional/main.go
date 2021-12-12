package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/yoshiken0927/paradigms/functional/model"
	"github.com/yoshiken0927/paradigms/functional/service"
)

var (
	name = flag.String("name", "", "mysql user name")
	age  = flag.Int("age", 0, "mysql user name")
)

func main() {
	fmt.Println(parse(service.UpdateEmployee("yoshiken")(31)))
}

func parse(e model.Employee) string {
	return e.Name + ":" + strconv.Itoa(e.Age) + ":" + strconv.Itoa(e.Version)
}
