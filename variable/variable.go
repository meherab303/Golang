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
}