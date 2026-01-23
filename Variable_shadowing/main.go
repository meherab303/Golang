package main
import "fmt"


//Go তে := (short variable declaration) শুধুমাত্র function এর ভিতরে ব্যবহার করা যায়
//package level এ := একদমই allowed না
var x int = 10
func main(){
	fmt.Println(x) 
	age:=25
	if age>20{
		x:=37    // new variable declare hoise.not update(jodi x=37 hoto tahole bujhaito je update kora hoise )
		fmt.Println(x)
	}
	fmt.Println(x)
}


// একই নামের একটা নতুন variable ছোট scope-এ declare করা, যেটা বাইরের (outer) variable-টাকে ঢেকে (shadow) ফেলে।