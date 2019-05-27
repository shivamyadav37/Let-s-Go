//Program To show Use of Variables in Go.
package main
import "fmt"
var x string = "Hi ,There " //Here we declare the variable outside main function.
				//now the value of x can be used by other functions too.
func main(){
	fmt.Println(x)
}
