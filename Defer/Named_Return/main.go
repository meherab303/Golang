package main

import "fmt"

func calc(b int)(result int){
	result=b
	fmt.Println(b)
	return 
}


/*
----Named return----

1)return variable shurutei create hoy and default value 0
2)named return er time e  function er shob kichu (including defer) execute howar por return value store hoy
3)return variable ta local variable hishebei kaj kore
4) return time e variale e mention korleo na korleo return variable tai return hobe
5)ekadhik return variable create kora jai

----unnamed return type----

1)jokhon e return dekhe tokhon e return value store hoye jai defer execute howar agei.


*/

func main(){
	a:=calc(10)
	fmt.Println(a)
}