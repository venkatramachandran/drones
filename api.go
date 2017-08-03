package main

import (
	"fmt"
)

var drones []Drone
var packages []Package

func init() {
	err := Request(DroneURL, &drones)
	if err != nil {
		fmt.Printf("error while fetching drones: %s\n", err)
	}
	err = Request(PackageURL, &packages)
	if err != nil {
		fmt.Printf("error while fetching packages: %s\n", err)
	}
}

//GetDrones returns the available drones
func GetDrones() []Drone {
	return drones
}

//GetPackages returns the available packages
func GetPackages() []Package {
	return packages
}
