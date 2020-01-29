package main

import "fmt"

func nakedRet(n int) (x, y int){
	
	x = n/2
	y = n - x

	return

}

func main(){

	fmt.Println(nakedRet(35))

}