package  main;

import "fmt"


var age int = 20 


// exposed to outside world if used capital var 
var I int =657

func main(){
	
	var age int = 34
	// cant redeclare if already declared in same scope 
	// age :=21
	fmt.Println(age)

}	