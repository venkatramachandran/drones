package main

import (
	"sort"
	"time"
)

//Assignment is a package to drone assignment
type Assignment struct {
	DroneID   int `json:"droneId"`
	PackageID int `json:"packageId"`
}

//Output is the returned output
type Output struct {
	Assignments       []Assignment `json:"assignments"`
	UnassignedPakages []int        `json:"unassignedPackageIds"`
}

//AddAssignment adds a new assignment to the output
func (o *Output) AddAssignment(droneID, packageID int) {
	o.Assignments = append(o.Assignments, Assignment{DroneID: droneID, PackageID: packageID})
}

//AddUnassignedPackage adds a new package to the unassigned list
func (o *Output) AddUnassignedPackage(packageID int) {
	o.UnassignedPakages = append(o.UnassignedPakages, packageID)
}

//RunAssignment creates the assignment of package to drone
func RunAssignment(drones Drones, packages Packages) Output {
	o := Output{}
	//sort drones by next available time
	sort.Sort(drones)
	//sort packages by deadline
	packages = sort.Reverse(packages).(Packages)
	currentDroneIndex := 0
	for _, p := range packages {
		//find if a drone can deliver
		drone := drones[currentDroneIndex]
		if p.Deadline < time.Now().Unix()+drone.AvailableIn() {
			//this drone can deliver this package
			o.AddAssignment(drone.DroneID, p.PackageID)
			currentDroneIndex++
		} else {
			//this drone cannot deliver this package
			o.AddUnassignedPackage(p.PackageID)
		}
	}
	return o
}
