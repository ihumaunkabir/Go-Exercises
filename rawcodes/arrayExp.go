package main

import "fmt"

func main(){
	
	var a [10]int
	a[0] = 1
	a[1] = 100

	fmt.Println(a[1])


	evens := [5]int{2, 4, 6, 8, 10}

	fmt.Println(evens)

	var new []int = evens[1:3]

	fmt.Println(new)


	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	d := names[0:2]
	e := names[1:3]
	fmt.Println(d, e)

	e[0] = "XXX"
	fmt.Println(d, e)
	fmt.Println(names)


}