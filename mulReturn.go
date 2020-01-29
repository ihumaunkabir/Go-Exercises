package main

import(

	"fmt"

)

func mulRet(x, y, z int) (int, int, int){
	
	return x+y , x*y, 4*x*y
}

func main(){
	
	add, mul, twice := mulRet(3, 3, 4)

	fmt.Println(add, mul, twice)

}