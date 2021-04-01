package main

import "fmt"

func add(a *int, b *int) int {
	c := *a + *b
	return c
}

func main() {
	var i,j *int
	a := 1
	b := 2
	i = &a
	j = &b
	ret := add(i,j)

	fmt.Println(ret)
	
}
