package main;

import (
	"fmt"
	"strconv"
)

func main(){

	var version float32 =1.01
	solidVersion :=int(version)
	fmt.Printf("Solid Version id %d " , solidVersion)
	fmt.Println()

	// var releaseName string ="septRelease-snapshot"
	// var bluffReleaseString string =string(releaseName)

	somerelease := 12
	// somereleaseNum := 42
	// print unicode char value for 12 and 42 
	var bluffRelease string =string(somerelease)
	fmt.Println(bluffRelease)


	// this pkg and function properly conver int to string 
	// Itoa means Int to ascii string
	var rightBluffRelease string =strconv.Itoa(somerelease)
	fmt.Println(rightBluffRelease)


}