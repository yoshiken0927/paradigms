package main

import (
	"flag"
	"fmt"

	"github.com/yoshiken0927/paradigms/functional/service"
)

var (
	name = flag.String("name", "", "mysql user name")
	age  = flag.Int("age", 0, "mysql user name")
)

func main() {
	flag.Parse()
	fmt.Println(
		service.UpdateEmployee(*name)(*age).Age,
	)
}
