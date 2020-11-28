package services

import (
	"github.com/go-event-api/domain"
	"github.com/go-event-api/repository"
)

func NewEventService() EventService {
	return EventService{
		Repository: repository.NewEventRepository(),
	}
}

type eventService interface {
	GetEvents() ([]domain.Event, error)
	CreateEvent(newEvent domain.Event) error
	GetOneEvent(id string) (domain.Event, error)
}

type EventService struct {
	Repository repository.EventRepository
}

func (e *EventService) GetEvents() ([]domain.Event, error) {
	return e.Repository.GetEvents()
}

func (e *EventService) CreateEvent(newEvent domain.Event) error {
	return e.Repository.CreateEvent(newEvent)
}

func (e *EventService) GetOneEvent(id string) (domain.Event, error) {
	return e.Repository.GetOneEvent(id)
}

func (e *EventService) UpdateEvent(updatedEvent domain.Event) (domain.Event, error) {
	return e.Repository.UpdateEvent(updatedEvent)
}

func (e *EventService) DeleteEvent(id string) error {
	return e.Repository.DeleteEvent(id)
}
