package main
import (
	"fmt"	
)

//Go তে := (short variable declaration) শুধুমাত্র function এর ভিতরে ব্যবহার করা যায়
//package level এ := একদমই allowed না

var I int=100

func main(){
	 a:=10  //short variable declaration
	//2. var a int =10
	//3. var a=10
	fmt.Println(a)
	//reassign
	a=20
	fmt.Println(a)
	const b=100  // const ,not changable this value
}








