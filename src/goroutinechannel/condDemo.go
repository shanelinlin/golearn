package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

var capacity = 10
var consumerNum = 2
var producerNum = 2

func producer(out chan<- int) {
	for i := 0; i < producerNum; i++ {
		go func(nu int) {
			//无条件循环
			for {
				cond.L.Lock()
				for len(out) == capacity {
					fmt.Println("Capacity Full, stop Produce")
					cond.Wait()
				}
				num := rand.Intn(100)
				out <- num
				fmt.Printf("Produce %d produce: num %d\n", nu, num)
				cond.L.Unlock()
				cond.Signal()

				time.Sleep(time.Second)
			}
		}(i)
	}
}

func consumer(in <-chan int) {
	for i := 0; i < consumerNum; i++ {
		go func(nu int) {
			//无条件循环
			for {
				cond.L.Lock()
				for len(in) == 0 {
					fmt.Println("Capacity Empty, stop Consume")
					cond.Wait()
				}
				num := <-in
				fmt.Printf("Goroutine %d: consume num %d\n", nu, num)
				cond.L.Unlock()
				time.Sleep(time.Millisecond * 500)
				cond.Signal()
			}
		}(i)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//创建一个无缓冲通道，目的让主goroutine进入阻塞等待
	quit := make(chan bool)
	product := make(chan int, capacity)
	producer(product)
	consumer(product)
	//只有主goroutine操作quit通道，主goroutine进入永久阻塞等待
	<-quit
}
