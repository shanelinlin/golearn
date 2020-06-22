//读写锁  sync.RWMutex
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwmutex sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("%d trying to lock \n", i)
			//读取锁
			rwmutex.RLock()
			fmt.Printf("reading data....%d \n ", i)
			//time.Sleep(time.Second)
			defer rwmutex.RUnlock()
			fmt.Printf("%d reading finished\n", i)

		}(i)
	}

	//主协程休眠一秒，其他3个协程不休眠，3个可以先执行，读与读不互斥，读写互斥
	time.Sleep(time.Second)
	//一般锁，写锁
	fmt.Println("main trying lock for writing ")
	rwmutex.Lock()
	fmt.Println("main writing")
	defer rwmutex.Unlock()
	fmt.Println("main finished")

}
