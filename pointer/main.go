package main
import "fmt"

type User struct{
	Name string
	Age int
}

func (u *User)userDetails(){

	u.Age=24  // internally eitake (*u).Age e convert kore.typing convient er jnno eita internally handle kora hoi
}

func print(number *[3]int){
	fmt.Println(number)
}

func main(){


	x:=20

	p:=&x
	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)

	*p=30
	fmt.Println(p)
	fmt.Println(x)

	// pointer variable declaration
	y:=30
	var pointer *int
	pointer=&y
	fmt.Println(pointer)

	//instance
	user:=User{
		Name:"Nahin",
		Age:20,
	}
	fmt.Println(user.Age) 
	user.userDetails()  // go internally eitake (&user).userdetails() e convert kore  .eita core convient wer jnno
	fmt.Println(user.Age)


	//func e pointer
	arr:=[3]int{1,2,3}
	print(&arr)


	 



}

//go te pointer arithmatic nei like p+1,p++ eigula nei.

//pointer use kore 1).Less memory use (large data copy kora lagena) 2).Performance faster 