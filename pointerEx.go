package main

import "fmt"

func main(){
	
	var p *int
	var i, j int = 21, 25
	var q *int

	q = &i
	p = &j

	fmt.Println(*p, *q)

}