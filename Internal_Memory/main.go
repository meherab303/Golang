package main

import "fmt"

const a=10   //code-segment

var p=100   //data segment


func call(){   //code segment
	add:=func(x int,y int){  //code segment e func add allocate hbe.abar call er stackframe e add name memory allocate hbe jeita code segment er add func er address(reference) ta thakbe.
		z:=x+y  //add er stackeframe .x,y add er stack frame e
		fmt.Println(z)
	}

	add(a,p)
}


  

func main(){  //code segment
	call()

	fmt.Println("The End")
}

func init (){  //code segment e 
	fmt.Println("I'm init func")
}



/*
 Internal memory or RAM 4ta part e devided

 1.code segment-->read only data(const,all func)
 2.data segment-->all global variable
 3.stack-->function call hole stack e memory allocate hoi oi func er jnno.jototuk jaiga allocate hoi oi area ta holo stackframe.stackframe e local variable declare hoi
 4.heap-->gc(grabage collector) eitake control kore

 */
