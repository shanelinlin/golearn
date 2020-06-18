package main

import "fmt"

// 新类型，
type mystring string

//string类型的别名
type aliastring = string
type myfunc func(string)

func say(v string) {
	fmt.Println(v)
}

func main() {
	var str mystring
	str = "string新类型：mystring"
	// mystring虽然可以存储string类型，但是golang认为他们不是同一种类型，要转换string(str)
	say(string(str))

	var s aliastring
	s = "string别名"
	// string的别名类型，golang认为类型的别名类型和原类型string还是同一种类型，不需要转换
	say(s)
	var myfun myfunc
	myfun = say
	myfun(s)

}
