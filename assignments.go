package main

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
func RunAssignment(drones []Drone, packages []Package) Output {
	return Output{}
}
