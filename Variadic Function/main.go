package main 

import "fmt"

/*

1)variadic function হলো এমন function যেটা একই type এর একাধিক argument variable সংখ্যায় নিতে পারে
2)কতগুলো argument আসবে আগে থেকে নির্দিষ্ট না থাকলে variadic function ব্যবহার করা হয়
3)synta : func funcName(param ...Type) ; here ... মানে variable number of arguments
4)param (Variadic parameter) ta slice e convert hoye jai
5)Variadic parameter সবসময় last parameter hobe
6)Println নিজেই variadic function

*/

func print(numbers...int){

	//eikhane numbers(variadic parameter) ta internally numbers:= []int{1,2,3,4,4,5,6} slice e convert hoye jai 

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func printslice(nums... int){

	for _,val:=range nums{
		fmt.Println(val)
	}

}
func mixedParameters(a string ,numbers...int){  //Variadic parameter সবসময় last parameter hobe
	fmt.Println(a,"got=>")
	total:=0
	for _,val := range numbers{
		total+=val
	}
	fmt.Println(total)
}

func nilSlice(numbers...int){
	fmt.Println(numbers)
}

func main(){


	//Variadic argument pass
	print(1,2,3,4,4,5,6)

	//Slice variadic এ pass করা
	nums := []int{1,2,3}
	printslice(nums...)

	//Mixed parameters pass 
	mixedParameters("Nahin",10,20,30,40,50)

	//nil slice variadic parameter
	nilSlice()


}