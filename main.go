package main

import (
	"fmt"
	"sort"
)

func main() {
	drones := GetDrones()
	packages := GetPackages()
	sort.Sort(drones)
	sort.Sort(packages)
	fmt.Printf("%v\n", RunAssignment(drones, packages))
}
