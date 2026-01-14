package main
import (
	"fmt"
	
	"package_scope/packag_scope"

)

func add(a int,b int)int{
	sum:= a+b

	// fmt.Println("sum function=>",sum)
	if a+b>120 {
		return a+b+-1
	}else if a+b==120{
		return  a+b-2
	}else {
		return sum
	}
	
}
func sub_mul(a int,b int)(int,int){
	
	switch true{  // switch e 2ta variable is not allowed 
	case a>b:
		return a-b,a*b
	case b>a:
		return b-a,a*b
	default:
		return a-b+1,a*b		
	}
	

	
	
}

func main(){
	 a:=10
	//2. var a int =10
	//3. var a=10
	fmt.Println(a)
	//reassign
	a=20
	fmt.Println(a)
	
	const b=100
	fmt.Println(b)
	sum:=add(a,b)
	fmt.Println("sum function=>",sum)
	sub,mul:=sub_mul(a,b)
	fmt.Println("sub function=>",sub,mul)


	addSUm:=packag_scope.Add(10,20)//function Add package packag_scope এর অংশ।
	fmt.Println("add from package scope=>",addSUm)

}
