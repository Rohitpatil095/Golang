package  main


import "fmt";


func main(){

	var resp string = "update to date record"
	var resp2 string = "record failed status"

	// all string are immutable string 
	// cant do string[2] ='g'
	
	// string concat 
	fmt.Print(resp+resp2 , "\n")

	bytesData :=[]byte(resp);

	//print ascii/utf value of each char in string
	fmt.Print(bytesData , "\n")

	// string slicing 
	fmt.Println("Sliced value at 3 is ", string(resp[3:9]))
	fmt.Println("Sliced value at 3 is ", string(resp[10:len(resp)]))

}