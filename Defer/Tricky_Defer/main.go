package main

import "fmt"


//unnamed return type
func calc()int{
	result:=20

	defer func ()  { //closure hishebe heap e store hobe karon result variable access korse
		result=result+10
		fmt.Println("calc result",result)
	}()

	return result
}

//named return type(easy)
func calcu()(result int){
	result=20

	defer func ()  { //closure hishebe heap e store hobe karon result variable access korse
		result=result+10
		fmt.Println("calcu result",result)
	}()

	return result
}


//named return type(hard)
func calculate()(result int){

	
    result=10
	show:= func(a int){    //closure hishebe heap e store hobe karon result variable access korse
		result=a+10
		fmt.Println("calculate result",result)
	}

	defer show(result) // eikhane 10 capture hoise

    result=30

	defer show(result) //eikhane 30 capture hoise

    return result
}


func main(){
	a:=calc()
	b:=calcu()
	c:=calculate()

	fmt.Println("calc a",a)
	fmt.Println("calcu b",b)
	fmt.Println("calculate c",c)
}

/*

---defer named return variable modify korte pare.why?


1) defer er time e closure jokhon hoy tokhon inner function +outer function 2tai exist thake.tai heap e variable  store howar time e value copy hoina.rather oi variable er reference  store hoy.

normaly closure e value hubabu copy hoi.karon outer function exist korena jokhon inner func ta call hoy(inner func return hoy and onno scope e  call hoy)

2)jeheto named return type e return value defer execute howar por store hoy tai 30 return hobe

3) jeheto unnamed return type e return value defer execute howar agei store hoy tai 20 return hobe.

4) shob defer closure hoina.




*/