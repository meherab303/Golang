package main

import "fmt"
	

/* ---(No.1)---

1)Slice হলো array-এর উপর একটি dynamic view।
2)Slice নিজে data store করে না
3)Slice array-এর data reference করে
4)slice header
 ├── pointer → underlying array address
 ├── length
 └── capacity

*/

func changeSlice(a []int) []int{
	a[0]=10
	a=append(a, 11)
	return a
}

func main(){
	//slice from array
	array1:=[5]int{1,2,3,4,5}

	slice1:=array1[1:4] //array1 er index 1 theke 4 er aag porjnto {2,3,4}-->ptr=array1 er index 1, len=3,capacity=4
	// capacity = array length - start index
	fmt.Println(slice1)

	//slice from slice
	slice2:=slice1[1:2]  //slice1 er index 1 theke 2 er aag porjnto-->ptr=array1 er index 2(karon eitao ultimately array1 kei point kore thakbe), len=1,capacity=3,
	fmt.Println(slice2)

	//slice literal
	slice3:=[]int{5,6,7,8,9,10}   // slice litaral e len and capacity same thake
	fmt.Println(slice3)

	//slice using make func with len
	slice4:=make([]int,3)  // capacity na dile capacity =length
	slice4[2]=10
	fmt.Println(slice4)

	//sluce using make func with len and capacity
	slice5:=make([]int,3,5)
	slice5[2]=30

	// slice[3]=20 eita dile runtime error ashbe .out of index range.karon Go index access er time e length check kore.capacity check korena
	

	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	//after append 
	slice5=append(slice5,3,4,5,6,7)  // append new slice return kore
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	//Nil slice
	var slice6 []int
	fmt.Println(slice6)
	slice6=append(slice6,6)
	fmt.Println(slice6)
	fmt.Println(cap(slice6)) // onek shomoi performance better er jnno capacity 4 diye shuru hoy

	//pointer address
	fmt.Println(&slice6[0])  //pointer er address-->slice er starting index er address e slice er ptr e thake

	//another way of declaration
	array2:=[]int{1,2,3,4,5,6}
	slice7:=array2[0:4:5] //s[a:b:c] ==> a→ start index(inclusive),b → end index(exclusive),c → max capacity index (exclusive),len=b-a, cap=c-a
	fmt.Println(slice7)
	fmt.Println(cap(slice7))


	//interview q-1

	var x []int
	x=append(x, 1)
	x=append(x, 2)
	x=append(x, 3)

	y:=x  //len=3,cap=4

	x=append(x, 4)
	y=append(y, 5)

	x[0]=10

	fmt.Println("x==>",x)
	fmt.Println("y==>",y)

	//interview q-2

	slice8:=[]int{1,2,3,4,5}
	slice8=append(slice8,6)
	slice8=append(slice8,7)

	a:=slice8[4:] //len=3 hbe,capacity hbe=6

	z:=changeSlice(a)

	fmt.Println("x==>",slice8)
	fmt.Println("y==>",z)
	

	// fmt.Println(slice8[7])--> runtime error ashbe karon Go index access er time e length check kore.capacity check korena
	

	fmt.Println(slice8[0:10])  //slice8[0:10] হলো slice expression। এটা slice8 থেকে index 0 থেকে 9 পর্যন্ত element নিয়ে নতুন slice বানায়। jeheto slice8 er capacity 10 tai 0 theke 10 er aag porjnto new slice banate parbe.jeigulai value nei oigulai by default 0 hbe.eikhane index 7 er por bakigula 0 hobe
	fmt.Println(cap(slice8[0:10])) //new capacity=old capacity -start index
	fmt.Println(len(slice8[0:10])) //length = end - start
	

}

/*  ---(No.2)---

 slice e 3ta part thake 
 1)pointer--> jei array theke slice hobe (underlying array bole oitake) oi array er jei index theke slice start hobe tar reference address

 2)length-->slice er length-->len(slice)
 ==>(length = end - start)

 3)capacity--->array te slice er start index theke array er baki full length porjonto
 ==>cap(slice)

 **i) when slice from array:
 ==>capacity = array length - start index
 reason:array length=array capacity

 **ii) when slice from slice
 ==>new capacity=old capacity -start index

 **iii) slice expression (slice[start:end]) rule: 0 ≤ start ≤ end ≤ cap(slice)

 **iv) Go Index Access Rule : Index Access er time e length check kore . capacity check korena.length er shoman ba boro index kokhonoi access kora jabena and value o boshano jabena 

 */


 /* ---(No.3)---

  what is capacity: slice সর্বোচ্চ কতটা element পর্যন্ত বাড়তে পারে (append করে) নতুন memory allocate না করেই। Capacity direct indexing এর জন্য না, append করার জন্য।



    -----Capacity full হলে-----
1️⃣ নতুন বড় array বানাবে
2️⃣ পুরোনো array এর data copy করবে
3️⃣ নতুন value add করবে
4️⃣ slice pointer নতুন array কে point করবে

** new underlying array heap e create hobe .karon append ekta func .eitar stackframe e create hole memory delete hoye jabe jokhon stackframe pop hobe.

** jodi existing array thake and capacity<=len  tahole shob data copy hobe heap e .then new value append hbe.existing array memory te theke jabe.

**jodi heap e create howa underlying array er poreo new value append kora hoi and capacity <=len tahole abar new underlying array create hobe .ager shob value copy hobe.then new value append hobe.Gc heap er ager underlying array delete kore dibe.

** capacity usually double hoye jai  ager capacity theke jeno barbar underlying array create korte na hoi.
** new capacity = 2*old capacity

*/