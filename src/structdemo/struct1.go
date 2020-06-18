package structdemo

import "fmt"

type Person struct {
	Name string
}

func Say(o One) {
	//return o.age
}

func GetAge() {
	var o One = One{900000}
	//通过成员:value这种map的方式初始化结构体
	var objo = One{age: 55}

	fmt.Println(o.age)
	fmt.Println(objo.age)

}
