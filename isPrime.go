package main

import "fmt"

func isPrime(n int) bool{
	
	var T bool = false

	for i := 2; i < n; i++ {
		
		if n % i == 0{
			T = true
		}

	}

	if T == true{
		return false
	} else{
		return true
	}

}


func main(){

	var n int
	var isp bool

	fmt.Scan(&n)

	isp = isPrime(n)

	if isp == true{
		fmt.Println("Prime")
	} else{
		fmt.Println("Not Prime")
	}

}