package main

import (
	"fmt"
	"sync"
	"time"
)

//普通互斥锁，相互排斥
var mutex sync.Mutex

func main() {
	for i := 0; i < 6; i++ {
		go func(n int) {
			fmt.Printf("%d 准备获取锁。。。\n", n)
			mutex.Lock()
			fmt.Printf("%d 获得锁，进行操作\n", n)
			//保证协程结束前Mutex会释放锁
			defer mutex.Unlock()
		}(i)
	}
	//主协程休眠一秒,其他6个协程不休眠，先获得执行
	time.Sleep(time.Millisecond * 1000)
	mutex.Lock()
	//主协程结束前释放锁
	defer mutex.Unlock()
	fmt.Println("main finished")

}
