package main

import (
	"fmt"
)

//take function as parameter
func process(a int ,b int, op func(p int ,q int)){  //op-->(parameter name) func(p int ,q int)-->(parameter type)
	op(a,b)
}

//return fanction
func returnFunc() func(x int,y int){  //func(x int,y int)-->return type
	return sub
}

// take func and return func both

func takeReturnBoth(callbackk func())func(x int ,y int){
	callbackk()
	return mul
}

func add(x int, y int){
	z:=x+y
	fmt.Println(z)
}
func sub(x int, y int){
	z:=x-y
	fmt.Println(z)
}
func mul(x int, y int){
	z:=x*y
	fmt.Println(z)
}

func print(){
	fmt.Println("I'm callback function for higher order function")
}

func main(){
	process(4,5,add)

	subtract:=returnFunc()
	subtract(7,3)
	multiplication:=takeReturnBoth(print)
	multiplication(4,5)

}


//First Order function-->normal value niye kaj kore like int,bool,float etc.
//Higher Order Function-->First order function niye kaj korbe.



/*higher order function--> any of the following 3

1. As  a parameter-->take function 
2.Return-->function 
3.Both


---------Higher Order function is also known as First class function and treated as first class citizen --------

1.higher order function ke first class function o bola hoi.
because-->

i)🔹 First-class citizen মানে কী?
যেটা করা যায়:

Variable এ assign করা যায়

Function এর argument হিসেবে পাঠানো যায়

Function থেকে return করা যায়

যে জিনিসগুলা পারে → সেগুলাই first-class citizen(example:int,float,bool,string eishob data. and also function(func ke variable e assign kora jai,argument hishebe pathanu jai))

ii)🔹 First-Class Function হলে কী কী করা যায়?

1️⃣ Variable এ assign করা যায়
2️⃣ Function এর argument হিসেবে পাঠানো যায়
3️⃣ Function থেকে return করা যায়

tai HOF is also called first class function
*/

/*Callback function:
higher order function e jei function ke argument hishebe pass kora hoi and higher order e function er vitore call kora hoi takei callback function bole
*/
