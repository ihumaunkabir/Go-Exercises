package main

import "fmt"

func main(){
	
	var n, p, i int64
	var mx int64 = 0

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&p)
		if p > mx{
			mx = p
		}
	}

	fmt.Println(mx)

}