package main

import "fmt"

func main(){
	var a int8 =-128
	var b int8=127

	var c uint16=12324

	var d float32=10.234566

	var flag bool=true  //1byte memory allocate ,format= %v

	var f byte=255  // byte=alias name for uint8

	var r rune= 'ব'  //rune=alias for int32=unicode character er jnno use hoi=%c

	

	fmt.Println(r) //unicode value
	fmt.Printf("%c\n",r) //character ,printf diye formate ta print kora jai  
	fmt.Printf("%.2f\n",d)
	fmt.Printf("%v\n",flag)


	fmt.Printf("type of r = %T\n",r)
	fmt.Printf("type of f = %T\n",f)

		 
	fmt.Println(a,b,c,d,flag,f)
}