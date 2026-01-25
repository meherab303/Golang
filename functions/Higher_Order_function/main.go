package main

import "fmt"

func process(a int ,b int, op func(p int ,q int)){  //op-->(parameter name) func(p int ,q int)-->(parameter type)
	op(a,b)
}

func add(x int, y int){
	z:=x+y
	fmt.Println(z)
}

func main(){
	process(4,5,add)
}


//First Order function-->normal value niye kaj kore like int,bool,float etc.
//Higher Order Function-->First order function niye kaj korbe.



/*higher order function--> any of the following 3

1. As  a parameter-->take function 
2.Return-->function 
3.Both

*/
