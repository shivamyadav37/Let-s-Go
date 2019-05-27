//Today we are going to write a program to understand boolean expressions in Go Language.
//Author--> Shivam Yadav

package main
import "fmt"
func main(){

	fmt.Println(true&&true)
	//this prints true as according to truth table 1&1=1
	

	fmt.Println(true&&false)
	//this prints false since 1&0=0..
	
	fmt.Println(true&&true)
	//it prints true since 1&1=1

	fmt.Println(true||true)
	//it prints true since 1+1=1 , as according to the truth table..

	fmt.Println(true||false)
	//it prints true since 1+0=1..
	fmt.Println(!true)
	//it prints false sionce not true=false..

}
