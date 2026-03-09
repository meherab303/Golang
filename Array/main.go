package main

import "fmt"

func change( arr[9] int){
	fmt.Println(arr)
	arr[0]=100
	fmt.Println(arr)
}

//using pointer
func changes(arr *[9]int){
	arr[0]=100
}

func main(){

	var arr1[2]int=[2]int{3,4}
	arr2:=[2]int {6,7}
	arr3:=[...]int {1,2,3,4,5,6,7,8,9,}  // size automatically detect

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(len(arr3))  // array length

	//iterate
	for i:=0;i<len(arr2);i++{
		fmt.Println(arr2[i])
	}

	//using range 
	for index,value:=range arr3{
		fmt.Println(value,index)
	}

	//array copy 
	arr4:=arr3
	fmt.Println(arr4)

	//array pass in func
	change(arr4)
	fmt.Println(arr4)

	//using pointer
	changes(&arr4)
	fmt.Println(arr4)

	//multi dimension
	var arr5=[2][3] int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(arr5)

	//iterate using range
	for i:=0;i<len(arr5);i++{
		for j:=0;j<len(arr5[i]);j++{
			fmt.Println(i,j,arr5[i][j])
		}
	}

	//using range

	for i,row :=range arr5{
		for j,col:=range row{
			fmt.Println(i,j,col)
		}
	}
}