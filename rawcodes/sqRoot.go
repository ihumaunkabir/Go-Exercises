package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	var j float64
	for i := 1; i < 10; i++ {

		z := float64(i)
		z -= (z*z -x) / (2*z)
		
		if z*z == x{
			j = z
		}
	}
	
	return j
}

func main() {

	fmt.Println(Sqrt(9))

}
