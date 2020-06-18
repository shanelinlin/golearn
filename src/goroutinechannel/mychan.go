package main

import "fmt"

func main2() {
	var c chan int         // 创建通道，通道保存数据类型未int
	c = make(chan int, 10) // 通道的缓冲区大小为10，最多保存10个值
	c <- 10                //把数据写入通道
	c <- 20
	/* for v := range c {
		fmt.Println(v)
	} */
	fmt.Println(<-c) //从通道读取数据，每读一次，自己向前位移一个位置
	fmt.Println(<-c)
	// 如果位置管道中数值的个数，如何判断还
	//  v, ok := <-ch // 当无法获取管道中的值后，ok==false

}
