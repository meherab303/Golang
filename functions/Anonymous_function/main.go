package main

import "fmt"

func main() {
	//basic anonymous or IIFE functions
	func(){
		fmt.Println("Hello! I'm Anonymous function")
	}()

	//IIFE-->Immediate Invoked Functions Expression
	//invoke mane call kora . Programing er bashai invoke bole.we should use invoke instead of call
	//anonymous function ke immediately ba sathe sathe call kora hole take IIFE bole.

	//anonymous or IIFE with parameter
	func(a int ,b int){
		fmt.Println(a+b)
	}(4,5)

	for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)
    }()
}


	/*
	Why use IIFE in Go?
	==To execute temporary logic, limit variable scope, and avoid loop variable issues in goroutines.
	*/
	
}