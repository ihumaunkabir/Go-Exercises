package main

import "fmt"

func main(){
	
	s := []int {2, 4, 5, 6, 7, 8}
	printSlice(s)

	s = s[1:]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	var new []int

	new = append(new, 2)
	new = append(new, 5)

	new = append(new, 4, 5, 6)

	printSlice(new)


}

func printSlice(s []int){
	fmt.Printf("Capacity: %d, Length: %d\n %v\n", cap(s), len(s), s )
}