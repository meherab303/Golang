package main

import "fmt"

func main() {
	//anonymous functions--> jei func er name nei tai anonymous

	/* func(){
	 	fmt.Println("Hello! I'm Anonymous function")
	 }
	*/

	//basic  IIFE functions
	func(){
		fmt.Println("Hello! I'm Anonymous function")
	}()

	//IIFE-->Immediatly Invoked Functions Expression
	//invoke mane call kora . Programing er bashai invoke bole.we should use invoke instead of call
	//anonymous function ke immediately ba sathe sathe call kora hole take IIFE bole.

	// IIFE with parameter
	func(a int ,b int){
		fmt.Println(a+b)
	}(4,5)

	


	/*
	Why use IIFE in Go?
	==To execute temporary logic, limit variable scope, and avoid loop variable issues in goroutines.
	*/


	//functions expresseion or assign function in variable

	add:=func(a int ,b int){
		c:=a+b
		fmt.Println(c)
	}
	add(4,5)
	
}