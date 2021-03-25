package main

import (
	"fmt"
)

func main()  {

	ch := make(chan int)
	defer close(ch)
	//go func() { ch <- 3+4 }()

	ch <- 3+4
	//time.Sleep(3)
	//var v int
	v := <-ch
	fmt.Println(v)


}