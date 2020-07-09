package main

import (
	"fmt"
	"time"
)

func say(s string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(s)

}

func main() {
	for i := 0; i < 5; i++ {
		go say("床前明月光")
	}
	//主协程结束执行前调用一次say
	defer say("低头思故乡")
	//主协程休眠一秒，其他子协程先执行
	time.Sleep(time.Second)
}
