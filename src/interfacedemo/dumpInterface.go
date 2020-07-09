package main

import (
	"fmt"
)

//空接口变量，interface{}
func main() {
	//定义一个空接口变量，空接口变量可以接受任意基本类型数据，但是不能是字面量类型如struct，map，slice，有点类似泛型，或者dart的dynamic类型
	var dumpi interface{}
	dumpi = 90
	fmt.Println(dumpi)
	dumpi = "空接口变量"
	fmt.Println(dumpi)
	dumpi = false
	fmt.Println(dumpi)

}
