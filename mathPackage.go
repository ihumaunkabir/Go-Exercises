package main

import (

	"fmt"
	"math"
	"math/rand"

)

func main(){
	
	fmt.Printf("Square Root: %g \n", math.Sqrt(10))
	fmt.Println(rand.Seed)
	fmt.Println(rand.Intn(50))
	fmt.Println(math.Pi)

	var i int
	var j float32
	j = math.Pi
	i = 90
	fmt.Println(i * 10)
	fmt.Println(j * 100)


}