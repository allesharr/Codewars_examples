package main

import (
	"fmt"

	"/ToCamelCase/tocamelcase"
)

func main() {

	ss := tocamelcase.ToCamelCase("the-stealth-warrior")
	fmt.Println(ss)
	ss = tocamelcase.ToCamelCase("The_Stealth_Warrior")
	fmt.Println(ss)
}
