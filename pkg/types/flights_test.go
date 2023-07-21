package types

import (
	"encoding/json"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

type flightsList struct {
	Src         string   `json:"src"`
	Dst         string   `json:"dst"`
	Connections int      `json:"connections"`
	Airports    []string `json:"airports"`
}

func TestFlights_GetSubRoutes(t *testing.T) {
	var flights = Flights{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}
	sub := flights.GetSubRoutes()
	result := [][][]string{{{"SFO", "ATL"}}, {{"SFO", "ATL"}, {"ATL", "GSO"}}, {{"SFO", "ATL"}, {"ATL", "GSO"}, {"GSO", "IND"}}, {{"SFO", "ATL"}, {"ATL", "GSO"}, {"GSO", "IND"}, {"IND", "EWR"}}, {{"ATL", "GSO"}}, {{"ATL", "GSO"}, {"GSO", "IND"}}, {{"ATL", "GSO"}, {"GSO", "IND"}, {"IND", "EWR"}}, {{"GSO", "IND"}}, {{"GSO", "IND"}, {"IND", "EWR"}}, {{"IND", "EWR"}}}
	if !reflect.DeepEqual(sub, result) {
		t.Log("algorithm get sub routes error")
		t.FailNow()
	}
}

func TestFlightInList_Sort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// opens the flight list used in the test
	data, err := os.ReadFile("./flights_test.json")
	if err != nil {
		t.Logf("error opening file ./flights_test.json: %v", err)
		t.FailNow()
	}

	// fill in the flight list
	var FlightsList []flightsList
	err = json.Unmarshal(data, &FlightsList)
	if err != nil {
		t.Logf("json.Unmarshal().error: %v", err)
		t.FailNow()
	}

	for _, flightsData := range FlightsList {

		// fills in the list to be used in the project
		var flights Flights = make([][]string, 0)
		for i := 0; i != len(flightsData.Airports)-1; i += 1 {
			flights = append(flights, []string{flightsData.Airports[i], flightsData.Airports[i+1]})
		}

		// sort the list in random order, because the list can be received out of order
		rand.Shuffle(len(flights), func(i, j int) {
			flights[i], flights[j] = flights[j], flights[i]
		})

		// sorts the received flight list
		flights.Sort()

		// check flight list order - start
		if flights[0][kSrc] != flightsData.Src {
			t.Logf("sort algorithm error")
			t.FailNow()
		}

		if flights[len(flights)-1][kDst] != flightsData.Dst {
			t.Logf("sort algorithm error")
			t.FailNow()
		}

		if len(flights) == 1 {
			continue
		}

		for i := 0; i != len(flights)-2; i += 1 {
			if flights[i][kDst] != flights[i+1][kSrc] {
				t.Logf("sort algorithm error")
				t.FailNow()
			}
		}
		// check flight list order - end
	}

}
