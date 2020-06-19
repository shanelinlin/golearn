package main

import (
	"fmt"
	"sync"
	//"time"
)

//WaitGroup类似java里的CountDownLatch倒数计时器，Add(int)是初始化等待的协程数，用于倒数的计数，Done()计数减少一，并判断计数是否为0，如果不为0，阻塞等待，知道计数为0
var wg sync.WaitGroup

func main() {
	wg.Add(3)

	go func(v int) {
		//time.Sleep(time.Second)
		defer wg.Done()
		fmt.Println("go", v)

	}(1)

	go func(v int) {
		//time.Sleep(time.Second)
		defer wg.Done()
		fmt.Println("go", v)

	}(2)

	go func(v int) {
		//time.Sleep(time.Second)
		defer wg.Done()
		fmt.Println("go", v)

	}(3)
	//主协程执行到这里，会判断计数是否已经是0，如果不为0就阻塞等待，
	wg.Wait()
	fmt.Println("main goroutine ended")
}
