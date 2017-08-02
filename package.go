package main

//Package indicates a package available to assign to a drone
type Package struct {
	Destination Location `json:"destination"`
	PackageID   int      `json:"packageId"`
	Deadline    int64    `json:"deadline"`
}
