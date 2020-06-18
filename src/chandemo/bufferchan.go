package main

import (
	"fmt"
)

//缓冲通道，写不会blcok，读取时，如果通道没有数据，会block，如果多协程，读写同时进行，写数据到通道的协程，完成写时，必须关闭通道，不然也会
func main1() {
	ch := make(chan string, 2)

	go func(sch chan string) {
		sch <- "hello channel of golang"
		//fmt.Printf("the chan value is:%s\n", <-sch)
		//通道用完记得关闭，不然main 协程遍历for 遍历语句读到缓冲通道没有数据时，会一直等它写被写入数据，直至被关闭，关闭时等于告诉等待读取的协程已经不写数据了
		//然后等待的协程也停止读取通道，继续往下执行。
		close(sch)
	}(ch)
	//当缓冲通道无数据时，被阻塞，等待其他goroutine写入数据，在读取
	for t := range ch {
		fmt.Println(t)
	}
	//缓冲通道关闭后，可读但不可写
	//ch <- "write after closed"
	fmt.Println("main go finished")

}
