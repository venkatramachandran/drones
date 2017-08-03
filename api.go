package main

import (
	"fmt"
)

var drones []Drone
var packages []Package

//GetDrones returns the available drones
func GetDrones() []Drone {
	err := Request(DroneURL, &drones)
	if err != nil {
		fmt.Printf("error while fetching drones: %s\n", err)
	}
	return drones
}

//GetPackages returns the available packages
func GetPackages() []Package {
	err := Request(PackageURL, &packages)
	if err != nil {
		fmt.Printf("error while fetching packages: %s\n", err)
	}
	return packages
}
