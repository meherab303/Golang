package main

import (
	"fmt"
	"functions/Initt"
)

func main(){

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
	
	fmt.Println("main functions init")
	
}

func init(){
	
	Initt.Init()
}