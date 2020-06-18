package main

import (
	"fmt"
	"reflect"
)

func main() {
	/* 	var arr [5]int
	   	arr[0] = 1
	   	arr[1]=2
	   	arr[2] =3
	   	arr[3] =4
	   	arr[4] =5 */
	//声明变量时同时初始化赋值，变量名后可以不用指定数据类型，也可不用指定数组的大小
	//var arr = [5]int{1,2,3,4,5}
	// 不用指定数组长度，会根据{}初始化值的大小来确定,但是定义数组时没有指定数组长度时，这种数组是可以改变长度的，可变长度的数组在golang里叫slice切片。
	var arr []int = []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(arr))
	fmt.Printf("%d,%d,%v\n", len(arr), cap(arr), arr)
	//固定长度的数组，不可变
	var arr1 = [2]int{7, 8}
	//固定长度的数组不能追加元素，只有可变长度数组才可以
	fmt.Println(reflect.TypeOf(arr1))
	//arr1 = append(arr1, 9) //会报错，固定长度数组不能追加元素
	fmt.Println(len(arr1))
	arr = append(arr, 6)
	//var arr = []int{1,2,3,4,5}
	//1代表第一个个元素，数组索引从0开始，这里从第二个因素开始取，直至结束
	fmt.Println(arr[1:])
	//从开头取，共取三个元素，这里3不是索引，索引是3-1，一定要注意
	fmt.Println(arr[:3])
	fmt.Println(len(arr))
}
