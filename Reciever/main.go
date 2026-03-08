package main
import "fmt"


//Go-তে receiver function মানে হলো struct এর সাথে attach করা function। 
// receiver + function = method
// Go compiler receiver func ke internally Normal func ei convert kore

type User struct {
	Name string
	Age int
}

type PersonInfo struct{
	Person User
	Address string	
}

func(u User) printUserDetails( a int){   //here (u User)--> reciever,printUserdetails-->method,u-->instance,a-->parameter
	fmt.Println(u.Age+a)
	fmt.Println(u.Name)
}


//Value receiver
func (p User) personDetails(){
  

	//perosn object ta only copy hoi.modify hobena
	pj.Age=40
	fmt.Println(p.Name,"-->",p.Age,"reciever")

}

// pointer reciver
func (usr *User) updateUser(){
	 usr.Age=33  
	 fmt.Println(usr.Name,"-->",usr.Age,"receiever") //Go compiler automatically dereference kore dei(*usr.Name,*usr.Age)
}


//Nested struct receiver
func (ui PersonInfo) PrintUserInfo(){
	fmt.Println(ui.Person.Name,"-->",ui.Person.Age,"-->",ui.Address)
}

func main(){

	var user1 User
	user1=User{
		Name:"Nahin",
		Age:20,
	}
	user1.printUserDetails(4)
	

	person:=User{
		Name:"Fahim",
		Age:25,
	}
	person.personDetails()
	fmt.Println(person.Name,"-->",person.Age,"instance")

	user:=User{
		Name:"Sajib",
		Age:26,
	}
	user.updateUser()
	fmt.Println(user.Name," -->",user.Age,"instance")

	userInfo:=PersonInfo{
		Person:User{
			Name:"Mamun",
			Age:20,
		},
		Address: "shamoly",
	}
	userInfo.PrintUserInfo()

}