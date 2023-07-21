package types

import "reflect"

const (
	// source of flight
	kSrc = 0
	// destination of flight
	kDst = 1
)

// SubRoutes receives combinatorial analysis of all sub routes
type SubRoutes struct {
	Start string
	End   string
	Route [][]string
}

// Flights gets list of flight connections in [][src, dst] format
// Example: {{"DUB","LHR"},{"LHR","GVA"},{"GVA","MXP"},{"MXP","NCE"}}
type Flights [][]string

// removeDuplicates remove duplicate itens
func (e *Flights) removeDuplicates(arr [][][]string) [][][]string {
	var result [][][]string

	for _, subArray := range arr {
		if !e.contains(result, subArray) {
			result = append(result, subArray)
		}
	}

	return result
}

// contains in slice
func (e *Flights) contains(arr [][][]string, subArray [][]string) bool {
	for _, a := range arr {
		if reflect.DeepEqual(a, subArray) {
			return true
		}
	}
	return false
}

// generateCombinatorialAnalysis generate combinatorial analysis
func (e *Flights) generateCombinatorialAnalysis(slice [][]string) [][][]string {
	var result [][][]string

	for i := 0; i < len(slice); i++ {
		var subArray [][]string
		for j := 0; j <= i; j++ {
			subArray = append(subArray, slice[j])
		}
		result = append(result, subArray)
	}

	e.removeDuplicates(result)

	return result
}

// GetSubRoutes returns all sub routes of a route
func (e *Flights) GetSubRoutes() (routes [][][]string) {
	routes = make([][][]string, 0)

	e.Sort()
	for k := range *e {
		routes = append(routes, e.generateCombinatorialAnalysis((*e)[k:])...)
	}

	return
}

// Sort sort the flight connections list
func (e *Flights) Sort() {
	sorted := make([][]string, 0)
	sorted = append(sorted, (*e)[0])

	// delete first key in list of flight connections
	*e = append((*e)[:0], (*e)[0+1:]...)

	for {

		// prevents an infinite loop bug
		security := 10

		for flightKey, flightData := range *e {

			// put key at top of sorted array
			if sorted[0][kSrc] == flightData[kDst] {
				// em cima
				sorted = append([][]string{flightData}, sorted...)
				*e = append((*e)[:flightKey], (*e)[flightKey+1:]...)
				break
			}

			// put key at bottom of sorted array
			if flightData[kSrc] == sorted[len(sorted)-1][kDst] {
				sorted = append(sorted, flightData)
				*e = append((*e)[:flightKey], (*e)[flightKey+1:]...)
				break
			}
		}

		if len(*e) == 0 {
			*e = sorted
			security = 20
			return
		}

		// security counter to prevent infinite loop bug
		security -= 1
		if security == 0 {
			panic("houston! we have a infinite loop bug")
		}
	}
}
