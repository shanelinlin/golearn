package main
import "fmt"

//传递指针作为参数
func funcWithPointerParam(p *string) string {
	return *p
}
func main(){
	var p *string 
	var p1 *string

	var name = "hello,pointer"
	p= &name

	var strarr [3]string =[3]string{"java","php","python"}
	fmt.Println(p1==nil)
	p1 = &strarr[0]
	fmt.Println(strarr)
	fmt.Println(*p1)
	fmt.Println(*p)
	
	var paramstr ="my param value"
	
	tstr := funcWithPointerParam(&paramstr)
	fmt.Println(tstr)

}

