package controllers

import (
	"dvdManagementSystemGo/pkg/models"
	"dvdManagementSystemGo/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewDVD models.DVD

/**
 *
 * get all dvds function controller
 **/
func GetDvds(w http.ResponseWriter, r *http.Request) {
	//get all dvd records from db
	NewDVDs := models.GetAllDvds()

	//json encode
	res, _ := json.Marshal(NewDVDs)

	//prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/**
 *
 * get dvd by id function controller
 **/
func GetDvdById(w http.ResponseWriter, r *http.Request) {
	//get params
	vars := mux.Vars(r)

	//get DvdId from params
	dvdId := vars["DvdId"]
	ID, err := strconv.ParseInt(dvdId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	//get dvd record from models
	DvdDetails, _ := models.GetDVDById(ID)

	//json encode
	res, _ := json.Marshal(DvdDetails)

	//prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/**
 *
 * create dvd function controller
 **/
func CreateDvd(w http.ResponseWriter, r *http.Request) {
	// refernce dvd struct
	CreateDvd := &models.DVD{}

	// json decode request body
	utils.ParseBody(r, CreateDvd)

	//create new record
	b := CreateDvd.CreateDvd()

	//json encode
	res, _ := json.Marshal(b)

	//prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/**
 *
 * Delete dvd function controller
 **/
func DeleteDvd(w http.ResponseWriter, r *http.Request) {
	//get params
	vars := mux.Vars(r)

	//get DvdId from params
	dvdId := vars["DvdId"]
	ID, err := strconv.ParseInt(dvdId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	//get dvd record from models
	DvdDetails := models.DeleteDvd(ID)

	//json encode
	res, _ := json.Marshal(DvdDetails)

	//prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
