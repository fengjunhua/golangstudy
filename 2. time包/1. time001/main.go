package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().Local().Year())
	fmt.Println(time.Now().Local().Month())
	fmt.Println(time.Now().Local().Day())
	fmt.Println(time.Now().Local().Date())
	fmt.Println(time.Now().Local().Hour())

}
