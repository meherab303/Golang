package main
import (
	"fmt"
	
	"package_scope/packag_scope"
	
)


func main(){




	// package scope function call
	addSUm:=packag_scope.Add(10,20)//function Add--> packag_scope(package name) এর অংশ।
	fmt.Println("add from package scope=>",addSUm)




	
	
}








