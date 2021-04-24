package entities

import (
	"errors"
	"github.com/hojabri/suss/pkg/validator"
)

type IpLog struct {
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Radius    uint16  `json:"radius"`
	Speed     float64 `json:"speed"`
	Ip        string  `json:"ip"`
	Timestamp uint64  `json:"timestamp"`
}

func (i *IpLog) Validate() error {
	if i.Lat == 0 {
		return errors.New("lat could not be zero")
	}
	if i.Lon == 0 {
		return errors.New("lon could not be zero")
	}
	if i.Radius == 0 {
		return errors.New("radius could not be zero")
	}
	if i.Speed == 0 {
		return errors.New("speed could not be zero")
	}
	if i.Ip == "" {
		return errors.New("ip could not be empty")
	}
	
	if err := validator.Pattern(i.Ip, `^((2[5][0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.(2[5][0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.(2[5][0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.(2[5][0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2}))$`); err != nil {
		return err
	}
	
	if i.Timestamp == 0 {
		return errors.New("timestamp could not be zero")
	}
	
	return nil
}
