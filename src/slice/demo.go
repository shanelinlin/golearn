/*
	切片slice是一种特殊的数组，是可变的数组，有两种定义方式 ，用数组定义的方式，但是[]datatype的[]不能带数组长度，带长度的就是固定数组了。
*/
package main

import "fmt"

func main() {
	//用数组定义的方式，注意[]没有指定数组长度的是可变数组，指定长度[n]是固定长度为n的固定数组
	var my = []string{"one", "two", "three", "four"}
	// make方式创建切片，长度2，容量10 长度是初始是分配的大小（元素个数）追加元素后长度是变化的，
	//容量是预设大小，当追加元素到切片，切片长没有超过容量,容量不变，当长度超过容量，容量以长度的倍数增加。
	var ml = make([]string, 2)
	//取切片索引
	for i := range ml {
		ml[i] = "he"
	}
	ml = append(ml, "so")
	//ml = append(ml, "so")

	printSlice(ml)
	printSlice(my)
	//给变长数组追加元素，新切片的长度是长度变化，容量是原来长度的两倍。
	my = append(my, "five")
	printSlice(my)
	//用切片的部分元素重建赋值给切片，被赋值的切片的所有元素全部先删除，在重建
	ml = my[:2]
	printSlice(ml)
	ml = append(ml, "fuck")
	printSlice(ml)
}

func printSlice(x []string) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
