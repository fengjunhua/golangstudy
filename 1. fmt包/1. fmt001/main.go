package main

import "fmt"

func main() {
	var a string
	a = "开始"
	fmt.Print('a', "\n")
	fmt.Print('a', "\n")
	fmt.Println("a%s", a)
	fmt.Printf("aa\n")
	fmt.Printf("bb%s", a)
	s := &a
	fmt.Println(s)
	fmt.Println(*s)
	fmt.Println(*&a)
}
