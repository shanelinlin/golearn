package main

import (
	"fmt"
)

type D struct {
	Name string
	Age  int
}

func main() {
	//空结构类型
	var test struct{}
	var a = D{Name: "hello", Age: 20}
	fmt.Println(a.Name)
	//实例化空结构struct{}{}
	test = struct{}{}
	fmt.Printf("%v %v", test, a)

}
