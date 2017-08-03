package main

//Package indicates a package available to assign to a drone
type Package struct {
	Destination Location `json:"destination"`
	PackageID   int      `json:"packageId"`
	Deadline    int64    `json:"deadline"`
}

//Packages is a collection of Packages
type Packages []Package

//Len provides the number of Packages
func (p Packages) Len() int {
	return len(p)
}

//Swap swaps positions of Packages
func (p Packages) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//Less returns true if i < j
func (p Packages) Less(i, j int) bool {
	return p[i].Deadline < p[j].Deadline
}
