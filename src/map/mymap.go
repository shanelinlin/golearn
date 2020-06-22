package main

import "fmt"

//map字典和数组不同，数组只要有大写，元素都会自动初始化值，int初始化0，string初始化"",map字典不会初始化
func main() {

	var mymap = map[int]string{1: "first value", 2: "second value", 3: "third value"}
	for k, v := range mymap {
		fmt.Println(k, v)
	}
	var dummymap map[int]int
	fmt.Println(dummymap == nil)
}
