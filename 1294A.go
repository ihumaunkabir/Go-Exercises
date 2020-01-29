package main

import(

	"fmt"
)

func main(){

	var q int

	fmt.Scan(&q)

	for i := 0; i<q; i++ {
		
		var a, b, c int64
		var n int64
		var sum int64 = 0

		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
		fmt.Scan(&n)

		sum = a + b + c + n

		if sum % 3 > 0 {
			fmt.Println("NO")
			continue
		} 

		sum /= 3

		if sum >= a && sum >= b && sum >= c {
			fmt.Println("YES")

		} else{
			fmt.Println("NO")

		}

	}


}