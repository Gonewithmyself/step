package main

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Player struct {
	Id       uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	var (
		user   = "root"
		psw    = "since1999"
		addr   = "127.0.0.1:3306"
		dbname = "test"
	)
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		user, psw, addr, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Player{})

	// res := db.Create(&Player{
	// 	Name:  "hello",
	// 	Email: "123",
	// })
	// fmt.Println(res.Error)

	var u Player
	// err = db.Table("users").Raw("SELECT * FROM players WHERE id=?", 1).Scan(&u).Error
	// err = Raw(db, "players", "WHERE id=?", 1).Scan(&u).Error
	err = db.Table("users").Where("id=?", 1).Find(&u).Error
	fmt.Println(err, u)

	u.Name = "user"
	err = db.Table("players").Save(&u).Error
	fmt.Println(err, u)
}

func Raw(db *gorm.DB, tb, sql string, values ...interface{}) (tx *gorm.DB) {
	var buf strings.Builder
	buf.WriteString("SELECT * FROM ")
	buf.WriteString(tb)
	buf.WriteString(" ")
	buf.WriteString(sql)
	return db.Raw(buf.String(), values...)
}
