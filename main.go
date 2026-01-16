package main
import (
	"fmt"
	
	"package_scope/packag_scope"

	"package_scope/variableShadowing"

	"package_scope/functions"

	
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

var I int=100

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


	// package scope function call
	addSUm:=packag_scope.Add(10,20)//function Add package packag_scope এর অংশ।
	fmt.Println("add from package scope=>",addSUm)

	// variable shadowing
	fmt.Println("variable shadowing example=>",)
	variableShadowing.Shadow()

	//types of functions

	//1.named or standard function
	functions.NamedOrStandardFunction()

	//init function I
	fmt.Println(I) // I=200 e hobe karon init function global scope e jeye I er value change kore felse
	
}


//init function-->you can not call this .computer autoatically call this function.it is executed before main function and after global scope allocation.you can not pass any parameter in init function and no return value .you can declare multiple init function in a single package .all init function will be executed before main function.

// onno package er init main function er age execute hobe.
/* Execution Order (Important!)
1.Imported packages এর init()
2.main package এর init()
3.main() function
*/

/*
multiple inint in single packge-->
Order: file top → bottom
*/


func init(){
	
	fmt.Println(I)  //eikhane I=100 hbe.
	I=200
	functions.Init()
}
func init(){
	
	functions.Init()
}





