package main

import "fmt"

func main(){
	
	sum := 0

	for i:= 1; i<100; i++ {

		sum += i

	}

	fmt.Println(sum)

	check := 1

	for ; check < 1000 ; {

		check += check
	}

	fmt.Println(check)

}