package main
import "fmt"
func main(){
	fmt.Println(len("Hello world"))
//it prints the length of string..
	fmt.Println("Hello world"[1])
//strings are indexed at 0
//so [1] should give e , but it gives 101
//101 is nothing but byte, and byte is also an integer..
	

//Now let's concatenate the strings..
	fmt.Println("Hello"+" World")
}
