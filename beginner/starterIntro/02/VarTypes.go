package main


import "fmt"


func main(){

	var a int8=3;
	var b int16= -456
	var c uint64= 456

	// avaialable 
	// uint- 8,16,32,64
	// int - 8,16,32,64 
	fmt.Println(a,b,c)

	// var add int= a+b
	//  cant work as diff int use type cast 
 	
	var add int = int(a) +int(b)
	fmt.Print(add)


	// float32 float64
	var someNum float32 = 34.5456
	fmt.Print(someNum)


	// complex128  float64 stored default
	var  comp complex64 = 1+2i   // float32 will be stored default
	fmt.Print(comp)
	fmt.Printf("%v , %T \n", real(comp) ,real(comp))

	fmt.Printf("%v , %T ", imag(comp) ,imag(comp))
}



