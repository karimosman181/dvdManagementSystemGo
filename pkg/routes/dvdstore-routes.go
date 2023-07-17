package routes

import (
	"dvdManagementSystemGo/pkg/controllers"

	"github.com/gorilla/mux"
)

/**
 *
 * create routes
 **/
var RegisterDVDStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/dvd/", controllers.CreateDvd).Methods("POST")
	router.HandleFunc("/dvd/", controllers.GetDvds).Methods("GET")
	router.HandleFunc("/dvd/{DvdId}", controllers.GetDvdById).Methods("GET")
	router.HandleFunc("/dvd/{DvdId}", controllers.UpdateDvd).Methods("PUT")
	router.HandleFunc("/dvd/{DvdId}", controllers.DeleteDvd).Methods("DELETE")
}
