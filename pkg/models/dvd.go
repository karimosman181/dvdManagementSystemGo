package models

import (
	"dvdManagementSystemGo/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type DVD struct {
	gorm.Model
	Movie  string `gorm:""json:"movie"`
	Studio string `json:"studio"`
}

/**
 *
 * Init and get connection to db
 **/
func init() {
	config.Connect()
	db = config.GETDB()
	db.AutoMigrate(&DVD{})
}

/**
 *
 * create new dvd record
 **/
func (d *DVD) CreateDvd() *DVD {
	db.NewRecord(d)
	db.Create(&d)
	return d
}

/**
 *
 *  get all dvd records
 **/
func GetAllDvds() []DVD {
	var DVDs []DVD
	db.Find(&DVDs)
	return DVDs
}

/**
 *
 * get dvd by id
 **/
func GetDVDById(Id int64) (*DVD, *gorm.DB) {
	var getDvd DVD
	db := db.Where("ID=?", Id).Find(&getDvd)
	return &getDvd, db
}

/**
 *
 * update dvd
 **/
func (d *DVD) UpdateDvd(Id int64) *DVD {
	db.Update(d).Where("ID=?", Id)
	return d
}

/**
 *
 * delete dvd
 **/
func DeleteDvd(Id int64) DVD {
	var dvd DVD
	db.Where("ID=?", Id).Delete(dvd)
	return dvd
}
