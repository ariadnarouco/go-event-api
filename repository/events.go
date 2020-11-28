package repository

import (
	"fmt"

	"github.com/go-event-api/domain"
)

type events map[string]domain.Event

var allEvents events

type eventRepository interface {
	GetEvents() (events, error)
}

type EventRepository struct {
	events events
}

func NewEventRepository() EventRepository {
	return EventRepository{
		events: initEvents(),
	}
}

func initEvents() events {
	return events{
		"1": {
			ID:          "1",
			Title:       "Introduction to Golang",
			Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		},
		"2": {
			ID:          "2",
			Title:       "Introduction to Golang",
			Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		},
		"3": {
			ID:          "3",
			Title:       "Introduction to Golang",
			Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		},
	}
}

func (events *events) getValues() []domain.Event {
	var result []domain.Event
	for _, e := range *events {
		result = append(result, e)
	}
	return result
}

func (r *EventRepository) GetEvents() ([]domain.Event, error) {
	return r.events.getValues(), nil
}

func (r *EventRepository) CreateEvent(newEvent domain.Event) error {
	_, ok := r.events[newEvent.ID]
	if ok {
		return fmt.Errorf("already-exists")
	}
	r.events[newEvent.ID] = newEvent
	return nil
}

func (r *EventRepository) GetOneEvent(id string) (domain.Event, error) {
	for _, e := range r.events {
		if e.ID == id {
			return e, nil
		}
	}
	return domain.Event{}, fmt.Errorf("not-found")
}

func (r *EventRepository) UpdateEvent(updatedEvent domain.Event) (domain.Event, error) {
	_, ok := r.events[updatedEvent.ID]
	if !ok {
		return domain.Event{}, fmt.Errorf("not-found")
	}
	r.events[updatedEvent.ID] = updatedEvent
	return updatedEvent, nil
}

func (r *EventRepository) DeleteEvent(id string) error {
	_, ok := r.events[id]
	if !ok {
		return fmt.Errorf("not-found")
	}
	delete(r.events, id)
	return nil
}
