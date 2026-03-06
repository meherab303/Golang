package main

import "fmt"

const a=10
var p =100

func outer() func(){

	money :=100
	age:=30
	fmt.Println("age ",age)
 
	show:=func(){    
		money=money+a+p  //closure hoise eikhane karon outer func er variable ke mone rakhtese(money variable) .a and p heap e store hobena.
		// money variable heap e jabe escape analysis er pore karon eikhane show return hoye geche.show pore call kora use hoite pare tokhon   money undefined ashte pare karon money outer func e declare hoise.outer func stack theke pop holei money pop hoye jabe.tai pore call kore use  korar jonno heap e store kora hoi money ke.
		fmt.Println(money)
	}
	return show  // heap memory te store hbe and  anonymous func er  address ta ref hishebe thakbe.heap e store hbe karon show pore call kora hoite pare.show jeheto outer func er tai outer func stack theke pop hole show o pop hbe.tai pore call kore kore use korar jnno heap e store hoi jeno pore  call kore pawa jai.
	//show যদি money use না করত, তাহলে show heap এ store হতো না।শুধু code segment এ logic থাকতো, reference stack/register থেকে manage হতো।
}

func call(){
	incr1:=outer()   // incr1 e heap er show er  address  ta ref hishbe store hbe.mane eikhane heap er jei address e show store hoise oi address ta incre1 e ref hishebe boshbe
	incr1()  // jotobar call hobe totobar money update hobe.
	incr1()

	incr2:=outer() // incr2 e heap er show er  address  ta ref hishbe store hbe.mane eikhane heap er jei address e show store hoise oi address ta incre2 e ref hishebe boshbe
	incr2()
	incr2()
}

func main(){

	call()
}

func init(){

	fmt.Println("==Bank==")
}

// heap er shorto-->যদি function পরে call হবে এবং কোনো outer variable use করবে, তাহলে সেই outer variable + function reference heap  e closure hishebe rakho”

/*
--- closure er shorto 2ta--
1.inner function uses outer function variable
2.inner funcstion is returned or assgined in a variable

Note: 
1)jodi inner func ke return na kora hoi taholeo inner func ke closure bola hobe jdi she outer func er variable use kore.
2) return na kora hole Heap e closure hishebe allocate hobena
3) jdi inner func ke outer func er moddei call kora hoi tokhon o heap e allocate hobena.karon tokhon ar onno kothao call kora jabena

*/

/*
---Escape Analysis---
1.compiler er ekta process ja decide kore kon variable stack e rakha jabe ar kon variable heap e rakha jabe
2.Escape Analysis--> compile time e hoy
3.Heap Allocation--> Runtime  e hoy

*/
