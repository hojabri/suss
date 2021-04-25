package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hojabri/suss/pkg/entities"
	"github.com/hojabri/suss/pkg/maxmind"
	"github.com/hojabri/suss/pkg/repository/crud"
	"github.com/hojabri/suss/pkg/suspiciousDetector"
	"github.com/hojabri/suss/pkg/susslogger"
	"net/http"
)

type SUSSService struct {
}

var EventsRepository crud.EventsRepository
var err error

func init() {
	EventsRepository, err = crud.NewEventsRepository()
	if err != nil {
		susslogger.Log().Fatal("error when creating new event repository: %s", err.Error())
	}
}

func (s SUSSService) NewUserSessionEvent(c *fiber.Ctx) *entities.Response {
	newEvent := entities.NewEvent{}
	if err = c.BodyParser(&newEvent); err != nil {
		susslogger.Log().Error(err)
		return &entities.Response{
			Code:     http.StatusNotAcceptable,
			Body:     err.Error(),
			Title:    "NotAcceptable",
			Message:  "error when parsing newEvent body",
			Instance: "NewUserSessionEvent",
		}
	}
	
	if err = newEvent.Validate(); err != nil {
		susslogger.Log().Error(err)
		return &entities.Response{
			Code:     http.StatusNotAcceptable,
			Body:     err.Error(),
			Title:    "NotAcceptable",
			Message:  "event body is not valid",
			Instance: "NewUserSessionEvent",
		}
	}
	
	geoInfo, err := maxmind.GetIpLocationInfo(newEvent.IpAddress)
	if err != nil {
		susslogger.Log().Error(err)
		return &entities.Response{
			Code:     http.StatusUnprocessableEntity,
			Body:     err.Error(),
			Title:    "UnprocessableEntity",
			Message:  fmt.Sprintf("geo information could not be found for IP:%s", newEvent.IpAddress),
			Instance: "NewUserSessionEvent",
		}
	}
	
	event := &entities.Event{
		Username:      newEvent.Username,
		UnixTimestamp: newEvent.UnixTimestamp,
		EventUuid:     newEvent.EventUuid,
		IpAddress:     newEvent.IpAddress,
		Lat:           geoInfo.Lat,
		Lon:           geoInfo.Lon,
		Radius:        geoInfo.Radius,
	}
	
	// Insert newEvent into db (will return the same event, if event_uuid already exists in the CityDB)
	event, err = EventsRepository.Create(event)
	if err != nil {
		return &entities.Response{
			Code:     http.StatusInternalServerError,
			Body:     err,
			Title:    "InternalServerError",
			Message:  err.Error(),
			Instance: "NewUserSessionEvent",
		}
	}
	
	// Retrieving preceding newEvent of this username
	precedingEvent, err := EventsRepository.PrecedingEvent(event)
	if err != nil {
		return &entities.Response{
			Code:     http.StatusInternalServerError,
			Body:     err,
			Title:    "InternalServerError",
			Message:  err.Error(),
			Instance: "NewUserSessionEvent",
		}
	}
	
	// Retrieving subsequent newEvent of this username
	subsequentEvent, err := EventsRepository.SubsequentEvent(event)
	if err != nil {
		return &entities.Response{
			Code:     http.StatusInternalServerError,
			Body:     err,
			Title:    "InternalServerError",
			Message:  err.Error(),
			Instance: "NewUserSessionEvent",
		}
	}
	
	travelToCurrentGeoSuspicious, speedToCurrentGeo := suspiciousDetector.IsMovementSuspicious(precedingEvent, event)
	travelFromCurrentGeoSuspicious, speedFromCurrentGeo := suspiciousDetector.IsMovementSuspicious(event, subsequentEvent)
	
	var precedingIpAccess entities.IpLog
	if precedingEvent != nil {
		precedingIpAccess = entities.IpLog{
			Lat:       precedingEvent.Lat,
			Lon:       precedingEvent.Lon,
			Radius:    precedingEvent.Radius,
			Speed:     speedToCurrentGeo,
			Ip:        precedingEvent.IpAddress,
			Timestamp: precedingEvent.UnixTimestamp,
		}
	}
	
	var subsequentIpAccess entities.IpLog
	if subsequentEvent != nil {
		subsequentIpAccess = entities.IpLog{
			Lat:       subsequentEvent.Lat,
			Lon:       subsequentEvent.Lon,
			Radius:    subsequentEvent.Radius,
			Speed:     speedFromCurrentGeo,
			Ip:        subsequentEvent.IpAddress,
			Timestamp: subsequentEvent.UnixTimestamp,
		}
	}
	
	processedResult := entities.ProcessedResult{
		CurrentGeo: entities.Geo{
			Lat:    event.Lat,
			Lon:    event.Lon,
			Radius: event.Radius,
		},
		TravelToCurrentGeoSuspicious:   travelToCurrentGeoSuspicious,
		TravelFromCurrentGeoSuspicious: travelFromCurrentGeoSuspicious,
		PrecedingIpAccess:              precedingIpAccess,
		SubsequentIpAccess:             subsequentIpAccess,
	}
	


	
	return &entities.Response{
		Code:     http.StatusOK,
		Body:     processedResult,
		Title:    "OK",
		Message:  "NewUserSessionEvent info",
		Instance: "NewUserSessionEvent",
	}
	
}
