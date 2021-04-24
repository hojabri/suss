package entities

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hojabri/suss/pkg/validator"
)

type NewEvent struct {
	Username      string    `json:"username" gorm:"index"`
	UnixTimestamp uint64    `json:"unix_timestamp"`
	EventUuid     uuid.UUID `json:"event_uuid"`
	IpAddress     string    `json:"ip_address"`
}

type Event struct {
	EventUuid     uuid.UUID `json:"event_uuid" gorm:"type:uuid;primarykey;"`
	Username      string    `json:"username" gorm:"index"`
	UnixTimestamp uint64    `json:"unix_timestamp"`
	IpAddress     string    `json:"ip_address"`
	Lat           float64   `json:"lat"`
	Lon           float64   `json:"lon"`
	Radius        uint16    `json:"radius"`
}

func (e *NewEvent) Validate() error {
	if e.IpAddress == "" {
		return errors.New("IpAddress could not be empty")
	}
	if e.Username == "" {
		return errors.New("username could not be empty")
	}
	if e.UnixTimestamp == 0 {
		return errors.New("UnixTimestamp could not be zero")
	}
	
	if err := validator.Pattern(e.IpAddress, `^((2[5][0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.(2[5][0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.(2[5][0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.(2[5][0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2}))$`); err != nil {
		return err
	}
	
	if err := validator.Pattern(e.EventUuid.String(), "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"); err != nil {
		return err
	}
	
	return nil
}
