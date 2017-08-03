package main

import "math"

//radiusOfEarth indicates radius of earth in KM
const radiusOfEarth = 6378.137

func degreesToRadians(degree float64) float64 {
	return degree * math.Pi / 180
}

//Location is a geo-location
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//DistanceTo gives distance from current location to the other
//location in kilometers, rounded to 2 decimals
func (l Location) DistanceTo(o Location) float64 {
	dLat := degreesToRadians(o.Latitude - l.Latitude)
	dLon := degreesToRadians(o.Longitude - l.Longitude)

	lLatInRadians := degreesToRadians(l.Latitude)
	oLatInRaidans := degreesToRadians(o.Latitude)

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLon/2), 2)*
		math.Cos(lLatInRadians)*math.Cos(oLatInRaidans)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	toRound := int64(c * radiusOfEarth * 100)

	rounded := float64(toRound) / float64(100.0)

	return rounded
}

//DepotLocation is the location of the depot
var DepotLocation = Location{Latitude: -37.816664, Longitude: 144.9637108}
