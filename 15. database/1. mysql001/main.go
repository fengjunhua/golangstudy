package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"

)

type User struct {
	ID         uint     `gorm:"id"         json:"id"`
	Username   string   `gorm:"username"   json:"username"`
	Firstname  string   `gorm:"firstname"  json:"firstname"`
	Lastname   string   `gorm:"lastname"   json:"lastname"`
	Telephone  uint     `gorm:"telephone"  json:"telephone"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

func main() {

	p1 := User{ID:1,Username: "蔡安宁",Firstname: "蔡",Lastname: "安宁",Telephone: 13520859533}
	p2 := User{ID:2,Username: "冯君华",Firstname: "冯",Lastname: "君华",Telephone: 13520859533}

	fmt.Println(p1)
	fmt.Println(p2)
	db, err := gorm.Open("mysql", "root:120225fjhFJH!@tcp(182.92.236.162:3306)/kubernetes?charset=utf8")
    defer db.Close()
	if err != nil{
		panic(err)
	}
	//db.NewRecord(p2)
	s := db.Create(&p2).Error
	fmt.Println(s)
	/*
	var user User
	err = db.Model(user).Create(&p1).Error
	if err !=nil {
		panic(err)
	}

	 */
	result := &User{}
	db.Where("username = ?", "蔡安宁").First(result)
	fmt.Println(*result,result.Username)

    //result := db.Create(&p1)
    //fmt.Println(result)


}
