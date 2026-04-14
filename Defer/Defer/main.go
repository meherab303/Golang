package main

import "fmt"


func main(){
	fmt.Println("first defer")

	defer fmt.Println("second defer")
	defer fmt.Println("fourth defer")
	defer fmt.Println("fifth defere")

	fmt.Println("third defer")

	x := 10
    defer fmt.Println(x)  // defer call korar time ei argument capture hoy.defer execute er time e x=10 e hobe
    x = 20
	
}

/*
---------defer--------------

1) defer= outer function return howar thik agee execute hobe
2)defer only function call kei delay korte pare.Variable delay korte parena
3)Multiple defer = LIFO order e execute hobe

**important**

4) defer store hoi heap/stack(easy er jonno dhore nilam heap e ) memoty te .Linked list data structure use kore store kore and stack er behavior execute kore (but stack na)

5)stack e ekta defer pointer list thake jeitai shobar last defer er pointer thake

6) heap e (defer func er pointer +argument+next) eigula thake.next e ager defer er pointer thake

*/