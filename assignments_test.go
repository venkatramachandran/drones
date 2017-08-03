package main

import (
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestRunAssignments(t *testing.T) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	//test cases
	//1. One drone at the depot with one package to deliver at 19 km with 1 hour deadline
	//expected output: drone is assigned to package
	drone1 := Drone{
		DroneID:  1,
		Location: DepotLocation,
	}
	package1 := Package{
		PackageID: 1,
		Destination: Location{
			Latitude:  -37.636269,
			Longitude: 144.941876,
		},
		Deadline: time.Now().Add(1 * time.Hour).Unix(),
	}
	drones := []Drone{drone1}
	packages := []Package{package1}
	output1 := RunAssignment(drones, packages)
	expectedOutput1 := Output{
		Assignments:       []Assignment{Assignment{PackageID: 1, DroneID: 1}},
		UnassignedPakages: []int{},
	}
	if len(output1.UnassignedPakages) != len(expectedOutput1.UnassignedPakages) {
		t.Fatalf("testcase %d failed! Expected length of unassinged packages: %d, got %d\n",
			1, len(expectedOutput1.UnassignedPakages), len(output1.UnassignedPakages))
	}
	if len(output1.Assignments) != len(expectedOutput1.Assignments) {
		t.Fatalf("testcase %d failed! Expected length of assigned packages: %d, got %d\n",
			1, len(expectedOutput1.Assignments), len(output1.Assignments))
	}
	//2. One drone 5km from the depot, with one package to deliver at 22km with 2 hour deadline
	//expected output: drone is assigned to the package
	droneCase2 := Drone{
		DroneID: 2,
		Location: Location{
			Latitude:  -37.781669,
			Longitude: 144.922477,
		},
	}
	packageCase2 := Package{
		PackageID: 2,
		Destination: Location{
			Latitude:  -37.947816,
			Longitude: 145.092937,
		},
		Deadline: time.Now().Add(2 * time.Hour).Unix(),
	}
	dronesCase2 := []Drone{droneCase2}
	packagesCase2 := []Package{packageCase2}
	outputCase2 := RunAssignment(dronesCase2, packagesCase2)
	expectedOutputCase2 := Output{
		Assignments:       []Assignment{Assignment{PackageID: 2, DroneID: 2}},
		UnassignedPakages: []int{},
	}
	if len(outputCase2.UnassignedPakages) != len(expectedOutputCase2.UnassignedPakages) {
		t.Fatalf("testcase %d failed! Expected length of unassinged packages: %d, got %d\n",
			2, len(expectedOutputCase2.UnassignedPakages), len(outputCase2.UnassignedPakages))
	}
	if len(outputCase2.Assignments) != len(expectedOutputCase2.Assignments) {
		t.Fatalf("testcase %d failed! Expected length of assigned packages: %d, got %d\n",
			2, len(expectedOutputCase2.Assignments), len(outputCase2.Assignments))
	}
	//3. One drone, 3 km from depot, with a package to deliver 5km away and 2 packages with 2 hour deadlines
	//expected output: drone is assigned to the first package and 2nd package is unassigned
	droneCase3 := Drone{
		DroneID: 3,
		Location: Location{
			Latitude:  -37.788316,
			Longitude: 144.958869,
		},
		Packages: []Package{
			{
				PackageID: 3,
				Destination: Location{
					Latitude:  -37.725886,
					Longitude: 144.965735,
				},
			},
		},
	}
	packageCase3 := Package{
		PackageID: 4,
		Destination: Location{
			Latitude:  -37.947816,
			Longitude: 145.092937,
		},
		Deadline: time.Now().Add(2 * time.Hour).Unix(),
	}
	packageCase3dot1 := Package{
		PackageID: 5,
		Destination: Location{
			Latitude:  -37.947816,
			Longitude: 145.092937,
		},
		Deadline: time.Now().Add(2 * time.Hour).Unix(),
	}
	dronesCase3 := []Drone{droneCase3}
	packagesCase3 := []Package{packageCase3, packageCase3dot1}
	outputCase3 := RunAssignment(dronesCase3, packagesCase3)
	expectedOutputCase3 := Output{
		Assignments:       []Assignment{Assignment{PackageID: 4, DroneID: 3}},
		UnassignedPakages: []int{5},
	}
	if len(outputCase3.UnassignedPakages) != len(expectedOutputCase3.UnassignedPakages) {
		t.Fatalf("testcase %d failed! Expected length of unassinged packages: %d, got %d\n",
			3, len(expectedOutputCase3.UnassignedPakages), len(outputCase3.UnassignedPakages))
	}
	if len(outputCase3.Assignments) != len(expectedOutputCase3.Assignments) {
		t.Fatalf("testcase %d failed! Expected length of assigned packages: %d, got %d\n",
			3, len(expectedOutputCase3.Assignments), len(outputCase3.Assignments))
	}
	//4. Two drones, first 2 km from depot and another 10km from depot, with 3 packages, with 1 hr, 90 mins and 2 hours deadline.
	//expected output: drone with 1 hour deadline is unassigned, other 2 are assigned to drones
	droneCase4 := Drone{
		DroneID: 7,
		Location: Location{
			Latitude:  -37.797269,
			Longitude: 144.957668,
		},
	}
	droneCase4dot1 := Drone{
		DroneID: 8,
		Location: Location{
			Latitude:  -37.718690,
			Longitude: 144.962818,
		},
	}
	packageCase4 := Package{
		PackageID: 11,
		Destination: Location{
			Latitude:  -37.947816,
			Longitude: 145.092937,
		},
		Deadline: time.Now().Add(1 * time.Hour).Unix(),
	}
	packageCase4dot1 := Package{
		PackageID: 12,
		Destination: Location{
			Latitude:  -37.947816,
			Longitude: 145.092937,
		},
		Deadline: time.Now().Add(90 * time.Minute).Unix(),
	}
	packageCase4dot2 := Package{
		PackageID: 13,
		Destination: Location{
			Latitude:  -37.947816,
			Longitude: 145.092937,
		},
		Deadline: time.Now().Add(2 * time.Hour).Unix(),
	}
	dronesCase4 := []Drone{droneCase4, droneCase4dot1}
	packagesCase4 := []Package{packageCase4, packageCase4dot1, packageCase4dot2}
	outputCase4 := RunAssignment(dronesCase4, packagesCase4)
	expectedOutputCase4 := Output{
		Assignments: []Assignment{Assignment{PackageID: 13, DroneID: 8},
			Assignment{PackageID: 12, DroneID: 7}},
		UnassignedPakages: []int{11},
	}
	if len(outputCase4.UnassignedPakages) != len(expectedOutputCase4.UnassignedPakages) {
		t.Fatalf("testcase %d failed! Expected length of unassinged packages: %d, got %d\n",
			4, len(expectedOutputCase4.UnassignedPakages), len(outputCase4.UnassignedPakages))
	}
	if len(outputCase4.Assignments) != len(expectedOutputCase4.Assignments) {
		t.Fatalf("testcase %d failed! Expected length of assigned packages: %d, got %d\n",
			4, len(expectedOutputCase4.Assignments), len(outputCase4.Assignments))
	}
}
