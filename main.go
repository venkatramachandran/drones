package main

import (
	"fmt"
)

func main() {
	drones := GetDrones()
	packages := GetPackages()
	fmt.Printf("%v\n", RunAssignment(drones, packages))
}
