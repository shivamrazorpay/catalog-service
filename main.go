package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service-catalog/auth"
	"service-catalog/boot"
	"service-catalog/internal"
)

func main() {

	boot.Initialize()

	r := mux.NewRouter()

	// Routes
	servicesRouter := r.PathPrefix("/services").Subrouter()
	servicesRouter.Use(auth.BasicAuthMiddleware())
	{
		servicesRouter.HandleFunc("", internal.GetServices).Methods("GET")
		servicesRouter.HandleFunc("/{serviceId}", internal.GetServiceById).Methods("GET")
		servicesRouter.HandleFunc("/{serviceId}/versions", internal.GetServiceVersions).Methods("GET")

		servicesRouter.HandleFunc("", internal.CreateService).Methods("POST")
		servicesRouter.HandleFunc("/{serviceId}", internal.UpdateService).Methods("POST")
	}

	// Starting server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
