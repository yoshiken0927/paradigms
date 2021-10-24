package main

import (
	"flag"
	"fmt"

	"github.com/yoshiken0927/paradigms/object/service"
)

var (
	name = flag.String("name", "", "mysql user name")
	age  = flag.Int("age", 0, "mysql user name")
)

func main() {
	flag.Parse()

	e, err := service.GetOlder(*name, *age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(e.Salary())
}
