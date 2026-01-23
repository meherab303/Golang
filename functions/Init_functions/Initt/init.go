package Initt

import "fmt"

func Init() {
	fmt.Println("I'm the first function that is executed first in the program")
}

func init() {
	fmt.Println("I'm  different package's init function that is executed before main functions init ")
}