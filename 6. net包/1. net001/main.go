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



	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error:" + err.Error())
	}
	fmt.Println(interfaces)
	for _, inter := range interfaces {
		fmt.Println(inter.Name)
		fmt.Println(inter.Index)
		fmt.Println(inter.HardwareAddr)
		fmt.Println(inter.Flags)
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("Error:" + err.Error())
	}
	fmt.Println(addrs)
}