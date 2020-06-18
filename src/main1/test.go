package main

import (
	"fmt"
	"structdemo"
)

type MyStruct struct {
	member string
}

func main() {
	var p structdemo.Person
	p.Name = "000"
	fmt.Println(p.Name)
	var my MyStruct
	my.member = "成员"
	fmt.Println(my.member)
	structdemo.GetAge()
}
