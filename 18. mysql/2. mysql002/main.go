package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

func main() {
	dns := ""
	db, err := gorm.Open(mysql.Open(dns),&gorm.Config{

	})
	if err != nil{
		panic(err)
	}

    fmt.Println(db)


}
