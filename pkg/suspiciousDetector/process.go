package suspiciousDetector

import (
	"github.com/hojabri/suss/pkg/config"
	"github.com/hojabri/suss/pkg/entities"
	"github.com/hojabri/suss/pkg/haversine"
	"math"
)

func IsMovementSuspicious(fromEvent *entities.Event, toEvent *entities.Event) (bool, float64) {
	
	if fromEvent == nil || toEvent == nil {
		return false, 0
	}
	
	sussThreshold := config.Config.GetFloat64("SUSPICIOUS_THRESHOLD")
	
	fromGeo := entities.Geo{
		Lat:    fromEvent.Lat,
		Lon:    fromEvent.Lon,
		Radius: fromEvent.Radius,
	}
	toGeo := entities.Geo{
		Lat:    toEvent.Lat,
		Lon:    toEvent.Lon,
		Radius: toEvent.Radius,
	}
	


	
	distanceMile, _ := haversine.Distance(fromGeo, toGeo)
	
	switch config.Config.GetString("DETECTION_MODE") {
	case "Optimistic":
		distanceMile = distanceMile -
			(kilometersToMiles(float64(fromEvent.Radius)) +
				kilometersToMiles(float64(toEvent.Radius)))
	case "Normal":
	case "Pessimistic":
		distanceMile = distanceMile +
			(kilometersToMiles(float64(fromEvent.Radius)) +
				kilometersToMiles(float64(toEvent.Radius)))
	default:
	}
	
	duration := math.Abs(float64(toEvent.UnixTimestamp-fromEvent.UnixTimestamp)) / 3600 //duration in hour
	speed := math.Abs(distanceMile / duration)
	if speed > sussThreshold {
		return true, speed
	}
	return false, speed
	
}

func milesToKilometers(miles float64) float64 {
	return miles * 1.609344
}

func kilometersToMiles(km float64) float64 {
	return km / 1.609344
}
