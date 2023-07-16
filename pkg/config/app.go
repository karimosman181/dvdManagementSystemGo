package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

/**
 *
 *  connection to db
 **/
func Connect() {
	d, err := gorm.Open("mysql", "gouser:gouser123/dvdstore?charaset=uft8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

/**
 *
 * return db connection
 **/
func GETDB() *gorm.DB {
	return db
}
