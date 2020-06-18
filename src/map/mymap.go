package main

import "fmt"

func main(){

	var mymap = map[int]string{1:"first value",2:"second value",3:"third value"}
	for k,v := range mymap{
		fmt.Println(k,v)
	}
	var dummymap map[int]int
	fmt.Println(dummymap==nil)
}

