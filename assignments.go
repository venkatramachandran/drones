package main

import (
	"log"
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
	log.Println("sorted drones")
	//sort packages by deadline
	sort.Sort(sort.Reverse(packages))
	log.Println("reverse sorted packages")
	currentDroneIndex := 0
	for _, p := range packages {
		log.Printf("got package as %v\n", p)
		//find if a drone can deliver
		if currentDroneIndex >= drones.Len() {
			//no drone available to deliver this package
			o.AddUnassignedPackage(p.PackageID)
		} else {
			drone := drones[currentDroneIndex]
			log.Printf("got drone as %v, available in %d\n", drone, drone.AvailableIn())
			log.Printf("time.now():%d\n", time.Now().Unix())
			estimatedTimeFromDepotToDestination := DepotLocation.DistanceTo(p.Destination) / DroneSpeed
			estimatedDeliveryTime := time.Now().Unix() +
				drone.AvailableIn() + int64(estimatedTimeFromDepotToDestination)
			log.Printf("estimated delivery time:%d\n", estimatedDeliveryTime)
			if p.Deadline >= estimatedDeliveryTime {
				//this drone can deliver this package
				o.AddAssignment(drone.DroneID, p.PackageID)
				currentDroneIndex++
			} else {
				//this drone cannot deliver this package
				o.AddUnassignedPackage(p.PackageID)
			}
		}
	}
	return o
}
