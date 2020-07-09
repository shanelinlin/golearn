package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a int32
	a = a + 10
	//原子操作变量a
	atomic.AddInt32(&a, 20)
	fmt.Println(a)
	fmt.Printf("variable a is %d \n", a)
}
