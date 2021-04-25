package maxmind

import (
	"fmt"
	"github.com/hojabri/suss/pkg/config"
	"github.com/hojabri/suss/pkg/entities"
	"github.com/hojabri/suss/pkg/susslogger"
	"github.com/oschwald/geoip2-golang"
	"net"
)
var CityDB *geoip2.Reader
var err error

func OpenCityDB() error {
	CityDB, err = geoip2.Open(fmt.Sprintf("geodb/%s",config.Config.GetString("GEO_CITY_DB")))
	if err != nil {
		return err
	}
	return nil
}
func CloseCityDB() error {
	err = CityDB.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetIpLocationInfo(ipStr string) (*entities.Geo,error) {
	ip := net.ParseIP(ipStr)
	var record *geoip2.City
	record, err = CityDB.City(ip)
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