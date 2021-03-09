package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

func main() {
	sysType := runtime.GOOS
	fmt.Println(sysType)
	//var home string
	if sysType == "Linux" {
		home,err := homeLinux()
		if err != nil{
           fmt.Println(err)
		}
		fmt.Println(home)
	}else if sysType == "windows" {
		home,err := homeWindows()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(home)
	}

}

func homeLinux() (string,error) {
	home := os.Getenv("HOME")

	return home,nil
}

func homeWindows()(string,error){
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == ""{
		home = os.Getenv("USERPROFILE")
	}
	if home == ""{
		return "",errors.New("HOMEDRIVE,HOMEPATH,and USERPROFILE is blank!")
	}
	return home,nil
}
