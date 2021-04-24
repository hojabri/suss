package haversine

// Reference:
// I used the implementation of haversine from this repository:
// https://github.com/umahmood/haversine
// And confirmed the formula with: https://en.wikipedia.org/wiki/Haversine_formula

import (
	"github.com/hojabri/suss/pkg/entities"
	"math"
)

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRadiusKm = 6371 // radius of the earth in kilometers.
)

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth. This function returns two units of measure, the first is the
// distance in miles, the second is the distance in kilometers.
func Distance(p, q entities.Geo) (mi, km float64) {
	lat1 := degreesToRadians(p.Lat)
	lon1 := degreesToRadians(p.Lon)
	lat2 := degreesToRadians(q.Lat)
	lon2 := degreesToRadians(q.Lon)
	
	diffLat := lat2 - lat1
	diffLon := lon2 - lon1
	
	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)
	
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	
	mi = c * earthRadiusMi
	km = c * earthRadiusKm
	
	return mi, km
}