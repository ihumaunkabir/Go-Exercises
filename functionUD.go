package main

import (
	
	"fmt"
	"math"

)

func cir(x float32) float32{
	
	return math.Pi * x * x
}


func main(){
	
	var r , area float32
	r = 4.4

	area = cir(r)

	fmt.Println(area)
}