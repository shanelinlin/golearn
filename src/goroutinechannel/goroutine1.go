package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s { // _不用切片的索引，这里只用到切片元素值，不关心索引
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, 9, 4, 3}
	c := make(chan int, 2)
	go sum(s[:len(s)/2], c) //求和：len(s)/2=3，[:3]取前三个元素
	go sum(s[len(s)/2:], c) //求和：[3:]取第三个元素之后的所有元素，3:表示第三个元素之后的元素
	fmt.Println("--------------------------------------")
	x, y := <-c, <-c // 从通道 c 中取值，每次<-c取一个值，并且位移到下一个值，通道里的值读取时，按filo先进后出处理
	close(c)         //关闭缓冲管道,不关闭通道，主协程会认为协程没有完成执行，一直等待，只有其他协程都完成了主协程才关闭
	fmt.Println(x, y, x+y)
	fmt.Println("--------------------------------------")
}
