package crud

import (
	"github.com/google/uuid"
	"github.com/hojabri/suss/pkg/entities"
	"github.com/hojabri/suss/pkg/repository/sqlite"
	"gorm.io/gorm"
)

type EventsRepository interface {
	Create(event *entities.Event) (*entities.Event, error)
	Update(event *entities.Event) error
	Delete(event *entities.Event) error
	FindAll() []*entities.Event
	FindByID(eventID uuid.UUID) (*entities.Event, error)
	DeleteByID(eventID uuid.UUID) error
	FindByUsername(username string) (*entities.Event, error)
	LastEvent(username string)  (*entities.Event, error)
	PrecedingEvent(currentEvent *entities.Event)  (*entities.Event, error)
	SubsequentEvent(currentEvent *entities.Event)  (*entities.Event, error)
}

type eventsRepository struct {
	connection *gorm.DB
}

func (e eventsRepository) PrecedingEvent(currentEvent *entities.Event) (*entities.Event, error) {
	var event entities.Event
	result := e.connection.Limit(1).
		Order("unix_timestamp desc").
		Find(&event, "username = ? AND unix_timestamp < ? AND event_uuid!= ?",
			currentEvent.Username ,
			currentEvent.UnixTimestamp,
			currentEvent.EventUuid)
	
	if result.Error!=nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &event, nil
	}
	return nil, nil
}

func (e eventsRepository) SubsequentEvent(currentEvent *entities.Event) (*entities.Event, error) {
	var event entities.Event
	result := e.connection.Limit(1).
		Order("unix_timestamp").
		Find(&event, "username = ? AND unix_timestamp > ? AND event_uuid!= ?",
			currentEvent.Username ,
			currentEvent.UnixTimestamp,
			currentEvent.EventUuid)
	
	if result.Error!=nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &event, nil
	}
	return nil, nil
}

func (e eventsRepository) LastEvent(username string) (*entities.Event, error) {
	var event entities.Event
	result := e.connection.Limit(1).Order("unix_timestamp desc").Find(&event, "username = ?", username)

	if result.Error!=nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &event, nil
	}
	return nil, nil
}

func (e eventsRepository) Create(event *entities.Event) (*entities.Event, error) {
	//we should check existence before saving to DB
	
	findEvent, err := e.FindByID(event.EventUuid)
	if err !=nil {
		return nil, err
	}
	
	// If exists, we don't save and return the existing event information
	if findEvent !=nil {
		return findEvent, nil
	}
	
	// If not, create new event
	event.EventUuid = uuid.New()
	result := e.connection.Create(&event)
	if result.Error != nil {
		return nil, result.Error
	}
	return event, nil
}

func (e eventsRepository) Update(event *entities.Event) error {
	result := e.connection.Save(&event)
	return result.Error
}

func (e eventsRepository) Delete(event *entities.Event) error {
	result := e.connection.Delete(&event)
	return result.Error
}

func (e eventsRepository) FindAll() []*entities.Event {
	var screens []*entities.Event
	e.connection.Find(&screens)
	return screens
}

func (e eventsRepository) FindByID(eventID uuid.UUID) (*entities.Event, error) {
	var event entities.Event
	result := e.connection.Find(&event, "event_uuid = ?", eventID)

	if result.Error!=nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &event, nil
	}
	return nil, nil
}

func (e eventsRepository) DeleteByID(eventID uuid.UUID) error {
	screen := entities.Event{}
	screen.EventUuid = eventID
	result := e.connection.Delete(&screen)
	return result.Error
}

func (e eventsRepository) FindByUsername(username string) (*entities.Event, error) {
	var event entities.Event
	result := e.connection.Find(&event, "username = ?", username)

	if result.Error!=nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &event, nil
	}
	return nil, nil
}

func NewEventsRepository() (EventsRepository, error) {
	if sqlite.DB == nil {
		err := sqlite.Connect()
		if err != nil {
			return nil, err
		}
	}
	return &eventsRepository{
		connection: sqlite.DB,
	}, nil
}