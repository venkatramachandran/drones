package main

//Location is a geo-location
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//DistanceTo gives distance from current location to the other location in kilometers
func (l Location) DistanceTo(o Location) float64 {
	return 0
}
