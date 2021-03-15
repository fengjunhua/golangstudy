package main

import (
	"fmt"
	"net"
)

func main() {
    var ip string
    ip = "10.121.140.91"
	if IP := net.ParseIP(ip); IP == nil {
		fmt.Println(IP)
	}
	type Perr struct{
		Name string `"json:"`
	}


}