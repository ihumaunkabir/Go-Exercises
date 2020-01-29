 package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	m := make(map[string] int)
	a := strings.Fields(s)

	for _, v := range a{
			m[v]++
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
