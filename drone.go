package main

import "log"

//Drone deliver items
type Drone struct {
	DroneID  int      `json:"droneId"`
	Location Location `json:"location"`
	Packages Packages `json:"packages"`
}

//HasPackageAssigned returns a boolean indicating whether
// the drone is servicing a package or not
func (d Drone) HasPackageAssigned() bool {
	return len(d.Packages) > 0
}

func (d Drone) distanceToDestination() float64 {
	if d.HasPackageAssigned() {
		distance := d.Location.DistanceTo(d.Packages[0].Destination)
		log.Printf("distance to destination: %f\n", distance)
		return distance
	}
	return 0
}

func (d Drone) distanceToDepot() float64 {
	if d.HasPackageAssigned() {
		distance := d.distanceToDestination() +
			d.Packages[0].Destination.DistanceTo(DepotLocation)
		log.Printf("distance to depot: %f\n", distance)
		return distance
	}
	distance := d.Location.DistanceTo(DepotLocation)
	log.Printf("distance to depot: %f\n", distance)
	return distance
}

//AvailableIn provides the time (in seconds) after which
//the drone will be available to deliver a package
func (d Drone) AvailableIn() int64 {
	timeToDepot := int64(d.distanceToDepot() / DroneSpeed * 3600)
	log.Printf("Time to depot in seconds: %d\n", timeToDepot)
	return timeToDepot
}

//DroneSpeed is the drone's speed in kmph
const DroneSpeed = 20

//Drones is a collection of Drones
type Drones []Drone

//Len provides the number of drones
func (d Drones) Len() int {
	return len(d)
}

//Swap swaps positions of drones
func (d Drones) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

//Less returns true if i < j
func (d Drones) Less(i, j int) bool {
	return d[i].AvailableIn() < d[j].AvailableIn()
}
