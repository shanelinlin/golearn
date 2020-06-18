/**
实际开发时包名和目录保持一致，这里为了方便运行，只有报名为main才能作为入口
**/

package main
import "fmt"
//import "hello/structdemo"

//return single value 方法定义参数指定类型和返回类型
func sum(x int ,y int) int{
	 var t  =x+y
	
	return t
}

// return multi value返回多值的函数
func multiVal(str0 string ,str1 string )(string ,string){
	return"你好"+str0,"hello"+str1
}
func main(){
	//变量可以不用显性的指定类型，会根据值动态判断类型
	var s  =sum(90,88)
	fmt.Println(s)
	var t1 ,t2 =multiVal("go语言","python语言")
	fmt.Println(t1,t2)

}
