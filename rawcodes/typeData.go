package main

import (

	"fmt"
	"math/cmplx"

)

var (

	ToBe bool = true
	intnum int = 45
	z complex128 = cmplx.Sqrt(-5 + 12i)
)
func main(){
	
	fmt.Printf("Type: %T , Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T , Value: %v \n", z, z)


}