package main

import "fmt"


// struct read only tai code segment e thake
type User struct{  //struct name(User) lowercase-->(unexported),Capital-->(exported)
	Name string   //here Name ,Age is  member variable or property
	Age int       // property name lowercase-->(unexported),Capital-->(exported)
}

//Nested Struct
type Client struct{
	Person User
	disease string
}

func main(){

	//Struct variable create-1
	var user1 User
	fmt.Println(user1.Name,"-->",user1.Age)  // by default Age=0,Name=" "

	//assign value-1
	user1.Name="Fahim"
	user1.Age=25
	fmt.Println(user1.Name,"-->",user1.Age)

	//assign value-2
	user1=User{    // USer type er instance or object creat jar name user1 and process is called instantiate
		Name:"Sajib",
		Age:26,
	}
	fmt.Println(user1.Name,"-->",user1.Age)

	//Struct variable create-2 and assign value-2(struct literal)

	user2:=User{   //preferable
		Name: "Al Mamun",
		Age:24,
	}
	fmt.Println(user2.Name,"-->",user2.Age)

	////Struct variable create-3 and assign value-3

	user3:=User{"Zakir",27}  // you must maintain order
	fmt.Println(user3.Name,"-->",user3.Age)

	//Nested Struct

	user4:=Client{
		Person:User{
			Name:"Nahin",
			Age:20,
		},
		disease: "fever",
	}
	fmt.Println(user4.Person.Name,"-->",user4.Person.Age,"-->",user4.disease )



	
	

}

func init(){
	fmt.Println("Learning Struct")
}