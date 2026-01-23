package main


import "fmt"


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
	const b=100


	sum:=add(a,b)
	fmt.Println("sum function=>",sum)
	sub,mul:=sub_mul(a,b)
	fmt.Println("sub function=>",sub,mul)
	
}








