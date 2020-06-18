package main

import (
	"fmt"
	"reflect"
)

func main() {
	var ivar int = 90
	t := reflect.TypeOf(ivar)
	var k reflect.Kind = t.Kind()
	fmt.Println(t.Name(), t.Kind().String())
	fmt.Println(k.String())
}
