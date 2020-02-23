package main

import "fmt"

func main(){
	
	var a int64

	fmt.Scan(&a)

	if a <= 2 || a%2 != 0{
		fmt.Println("NO")
	} else{
		fmt.Println("YES")
	}

}