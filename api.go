package main

import (
	"fmt"
)

var drones Drones
var packages Packages

//GetDrones returns the available drones
func GetDrones() Drones {
	err := Request(DroneURL, &drones)
	if err != nil {
		fmt.Printf("error while fetching drones: %s\n", err)
	}
	return drones
}

//GetPackages returns the available packages
func GetPackages() Packages {
	err := Request(PackageURL, &packages)
	if err != nil {
		fmt.Printf("error while fetching packages: %s\n", err)
	}
	return packages
}
