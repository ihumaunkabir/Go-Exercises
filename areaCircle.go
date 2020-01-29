package main

import (
	 "fmt"
	 "math"
)

func main(){
	
	var r float64
	var ar float64

	fmt.Scan(&r)
	ar = math.Pi * r * r

	fmt.Printf("%.10f", ar)

}