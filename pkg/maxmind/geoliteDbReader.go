package maxmind

import (
	"fmt"
	"github.com/hojabri/suss/pkg/config"
	"github.com/hojabri/suss/pkg/entities"
	"github.com/hojabri/suss/pkg/susslogger"
	"github.com/oschwald/geoip2-golang"
	"net"
)
var DB *geoip2.Reader
var err error

func OpenDB() error {
	DB, err = geoip2.Open(fmt.Sprintf("geolitedb/%s",config.Config.GetString("GEODB")))
	if err != nil {
		return err
	}
	return nil
}
func CloseDB() error {
	err = DB.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetGeoInfo(ipStr string) (*entities.Geo,error) {
	ip := net.ParseIP(ipStr)
	var record *geoip2.City
	record, err = DB.City(ip)
	if err != nil {
		susslogger.Log().Error(err)
		return nil, err
	}

	
	return &entities.Geo{
		Lat:    record.Location.Latitude,
		Lon:    record.Location.Longitude,
		Radius: record.Location.AccuracyRadius,
	},nil
	
}