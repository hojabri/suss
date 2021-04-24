package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hojabri/suss/pkg/entities"
	"github.com/hojabri/suss/pkg/repository/crud"
	"github.com/hojabri/suss/pkg/susslogger"
	"net/http"
)

type SUSSService struct {
}

var EventsRepository crud.EventsRepository
var err error

func init() {
	EventsRepository,err = crud.NewEventsRepository()
	if err!=nil {
		susslogger.Log().Fatal("error when creating new event repository: %s", err.Error())
	}
}

func (s SUSSService) NewUserSessionEvent(c *fiber.Ctx) *entities.Response {
	event := entities.Event{}
	if err := c.BodyParser(&event); err != nil {
		susslogger.Log().Error(err)
		return &entities.Response{
			Code:     http.StatusNotAcceptable,
			Body:     err.Error(),
			Title:    "NotAcceptable",
			Message:  "error when parsing event body",
			Instance: "NewUserSessionEvent",
		}
	}

	// Retrieving preceding event of this username
	precedingEvent, err :=EventsRepository.PrecedingEvent(&event)
	if err != nil {
		return &entities.Response{
			Code:     http.StatusInternalServerError,
			Body:     err,
			Title:    "InternalServerError",
			Message:  err.Error(),
			Instance: "NewUserSessionEvent",
		}
	}
	
	// Retrieving subsequent event of this username
	subsequentEvent, err :=EventsRepository.SubsequentEvent(&event)
	if err != nil {
		return &entities.Response{
			Code:     http.StatusInternalServerError,
			Body:     err,
			Title:    "InternalServerError",
			Message:  err.Error(),
			Instance: "NewUserSessionEvent",
		}
	}

	// Insert event into db
	id, err := EventsRepository.Create(&event)
	if err != nil {
		return &entities.Response{
			Code:     http.StatusInternalServerError,
			Body:     err,
			Title:    "InternalServerError",
			Message:  err.Error(),
			Instance: "NewUserSessionEvent",
		}
	}
	event.ID = id


	return &entities.Response{
		Code:     http.StatusOK,
		Body: fiber.Map{
			"event" :      event,
			"precedingEvent" : precedingEvent,
			"subsequentEvent" : subsequentEvent,
		},
		Title:    "OK",
		Message:  "NewUserSessionEvent info",
		Instance: "NewUserSessionEvent",
	}

}
