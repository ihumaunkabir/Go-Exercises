package main

import (

	"fmt"
	"math"

)

func main(){
	
	var i int = 5
	var j float64 = math.Pow(3,3)

	if i < 10 {

		fmt.Printf("%v is less than 10\n", i)
	} else {

		fmt.Printf("Entered in ELSE")
	}

	fmt.Printf("%v is great\n", j)

}