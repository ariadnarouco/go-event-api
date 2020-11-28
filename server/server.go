package server

import (
	"net/http"

	"github.com/go-event-api/controllers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//Run will start the server on port 8080 and set a series of endpoints and handlers.
func Run() {
	log.SetFormatter(&log.JSONFormatter{})
	controller := controllers.NewEventController()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.HomeLink)
	router.HandleFunc("/event", controller.CreateEvent).Methods("POST")
	router.HandleFunc("/events", controller.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", controller.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", controller.UpdateEvent).Methods("PUT")
	router.HandleFunc("/events/{id}", controller.DeleteEvent).Methods("DELETE")
	log.WithFields(log.Fields{
		"server": "8080",
		"app":    "event-api",
	}).Info("The server is running")

	log.Fatal(http.ListenAndServe(":8080", router))
}
