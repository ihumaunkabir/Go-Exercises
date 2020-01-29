package main

import "fmt"

func main(){
	
	var i int = 10

	switch i {

		case 5:
			fmt.Println("Five")
		case 10:
			fmt.Println("Ten")
		case 11:
			fmt.Println("Eleven")
		default:
			fmt.Println("Default")
			

	}

}