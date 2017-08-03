package main

import (
	"encoding/json"
	"net/http"
)

//DroneURL is the endpoint to fetch drones
const DroneURL = "https://codetest.kube.getswift.co/drones"

//PackageURL is the endpoint to fetch packages
const PackageURL = "https://codetest.kube.getswift.co/packages"

//Request makes a HTTP GET call, converts the body to
//object and returns it
func Request(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
