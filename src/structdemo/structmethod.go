package main //structdemo

import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {
	// 成员变量，首字母是小写，在当前源码种可以使用 :结构体变量.成员直接访问，外部源码要使用，成员名第一字母必须大写，golang约定，首字母大写的是导出给外部源码访问
	radius float64
}

func main() {

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius

}
