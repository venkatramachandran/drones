package main

import (
	"testing"
)

type testCase struct {
	From     Location
	To       Location
	Distance float64
}

var testData = []testCase{
	{Location{Latitude: 0, Longitude: 0}, Location{Latitude: 0, Longitude: 0}, 0},
	{Location{Latitude: -37.8274812, Longitude: 144.9352465},
		Location{Latitude: -33.8679049, Longitude: 151.1924823},
		715.98},
}

func TestDistance(t *testing.T) {
	for i, data := range testData {
		distance := data.From.DistanceTo(data.To)
		if distance != data.Distance {
			t.Fatalf("testcase %d failed! Expected: %f, got %f\n",
				i, data.Distance, distance)
		}
	}

}
