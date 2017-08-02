package main

//Drone deliver items
type Drone struct {
	DroneID  int       `json:"droneId"`
	Location Location  `json:"location"`
	Packages []Package `json:"packages"`
}

//HasPackageAssigned returns a boolean indicating whether
// the drone is servicing a package or not
func (d Drone) HasPackageAssigned() bool {
	return len(d.Packages) > 0
}
