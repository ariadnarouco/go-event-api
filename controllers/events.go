package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-event-api/domain"
	"github.com/go-event-api/services"
	"github.com/gorilla/mux"
)

func NewEventController() EventController {
	return EventController{
		Service: services.NewEventService(),
	}
}

type EventController struct {
	Service services.EventService
}

func (c *EventController) HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func (c *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent domain.Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	err = c.Service.CreateEvent(newEvent)
	if err != nil {
		w.WriteHeader(getStatus(err.Error()))
		json.NewEncoder(w).Encode("error " + err.Error())

		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func (c *EventController) GetOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	oneEvent, err := c.Service.GetOneEvent(eventID)
	if err != nil {
		w.WriteHeader(getStatus(err.Error()))
		json.NewEncoder(w).Encode("error " + err.Error())
		return
	}

	json.NewEncoder(w).Encode(oneEvent)
	w.WriteHeader(http.StatusOK)

}

func (c *EventController) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent domain.Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	updatedEvent, err := c.Service.UpdateEvent(newEvent)
	if err != nil {
		w.WriteHeader(getStatus(err.Error()))
		json.NewEncoder(w).Encode("error " + err.Error())

		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedEvent)
}

func (c *EventController) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	err := c.Service.DeleteEvent(eventID)
	if err != nil {
		w.WriteHeader(getStatus(err.Error()))
		json.NewEncoder(w).Encode("error " + err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *EventController) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := c.Service.GetEvents()
	if err != nil {
		w.WriteHeader(getStatus(err.Error()))
		json.NewEncoder(w).Encode("error " + err.Error())
		return
	}
	json.NewEncoder(w).Encode(events)
	w.WriteHeader(http.StatusOK)
}

func getStatus(errorType string) int {

	switch sc := errorType; sc {
	case "not-found":
		return http.StatusNotFound
	case "already-exists":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
