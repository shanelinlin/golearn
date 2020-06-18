package main

import (
	"fmt"
	//"time"
)

//非缓冲通道不用关闭,必须有2个以上的协程同时读写通道，否则就会永久堵塞导致死锁
func main() {

	nch := make(chan string)

	//写通道协程
	go func(sch chan string) {
		for i := 0; i < 6; i++ {
			nch <- "hello"
		}
	}(nch)

	//读取协程
	go func(sch chan string) {

		var str string

		for i := 0; i < 6; i++ {
			//time.Sleep(1 * time.Second)
			str = <-nch
			fmt.Println(str)
		}
	}(nch)
	//这一步很重要，主协程也参与操作非缓冲通道，这样就有3个协程一起并发执行，不会造成死锁
	mystr := <-nch
	fmt.Println(mystr)
	fmt.Println("main ends")
}
