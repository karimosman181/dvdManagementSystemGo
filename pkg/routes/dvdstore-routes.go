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
	router.HandleFunc("/dvd/", controllers.GetDvd).Methods("GET")
	router.HandleFunc("/dvd/{Id}", controllers.GetDvdById).Methods("GET")
	router.HandleFunc("/dvd/{Id}", controllers.UpdateDvd).Methods("PUT")
	router.HandleFunc("/dvd/{Id}", controllers.DeleteDvd).Methods("DELETE")
}
