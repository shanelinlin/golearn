package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {
	mutex.Lock()
	fmt.Println("hello,mutex")

	defer mutex.Unlock()

}
